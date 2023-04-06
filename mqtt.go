//go:generate go run ./helpers/

package hass_mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iancoleman/strcase"
	"github.com/kjbreil/hass-mqtt/device"
	"github.com/kjbreil/hass-mqtt/entities"
	"log"
	"path/filepath"
	"time"
)

type Client struct {
	NodeID    string
	config    Config
	OnConnect func(client mqtt.Client)
	devices   map[string]*device.Device // Device Identifiers map to devices
	mqtt      mqtt.Client
}

type Config struct {
	NodeID string `json:"node_id" yaml:"node_id"`
	MQTT   struct {
		Host     string  `json:"host" yaml:"host"`
		Port     int     `json:"port" yaml:"port"`
		SSL      bool    `json:"ssl" yaml:"ssl"`
		UserName *string `json:"user_name,omitempty" yaml:"user_name"`
		Password *string `json:"password,omitempty" yaml:"password"`
	} `json:"MQTT" yaml:"mqtt"`
}

func NewClient(c Config) (*Client, error) {

	client := &Client{
		NodeID:  strcase.ToDelimited(c.NodeID, uint8(0x2d)),
		config:  c,
		devices: make(map[string]*device.Device),
	}
	return client, nil
}

var (
	ErrNoDeviceFound = fmt.Errorf("no device found")
)

func (c *Client) Connect() error {
	if len(c.devices) == 0 {
		return ErrNoDeviceFound
	}

	entities.NodeID = c.NodeID

	var err error
	for _, d := range c.devices {
		err = d.Initialize()
		if err != nil {
			return err
		}
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", c.config.MQTT.Host, c.config.MQTT.Port))
	opts.SetClientID(c.NodeID)
	opts.SetOrderMatters(false)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message on Default Handler: %s from topic: %s\n", msg.Payload(), filepath.Base(msg.Topic()))
	})

	opts.SetOnConnectHandler(
		func(client mqtt.Client) {
			if c.OnConnect != nil {
				c.OnConnect(client)
			}
			for _, d := range c.devices {
				d.Subscribe()
			}
		},
	)
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Println("MQTT connection lost")
	}

	c.mqtt = mqtt.NewClient(opts)

	for _, d := range c.devices {
		d.SetMQTTFields(c.mqtt)
	}

	if token := c.mqtt.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Client) Add(dev *device.Device) error {
	if _, ok := c.devices[dev.GetUniqueId()]; ok {
		return fmt.Errorf("%s entity already exists", dev.GetUniqueId())
	}
	c.devices[dev.GetUniqueId()] = dev
	return nil
}

func (c *Client) Get(identifier string) *device.Device {
	if dev, ok := c.devices[identifier]; ok {
		return dev
	}
	return nil
}

func (c *Client) Disconnect() {
	for _, d := range c.devices {
		d.UnSubscribe()
	}
}
