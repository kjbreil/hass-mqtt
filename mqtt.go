//go:generate go run ./helpers/

package hass_mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	"github.com/iancoleman/strcase"
	"github.com/kjbreil/hass-mqtt/device"
	"github.com/kjbreil/hass-mqtt/entities"
	"path/filepath"
	"time"
)

type Client struct {
	NodeID    string
	config    Config
	OnConnect func(client mqtt.Client)
	devices   map[string]*device.Device // Device Identifiers map to devices
	mqtt      mqtt.Client
	logger    logr.Logger
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
	return NewClientWithLogger(c, defaultLogger())
}
func NewClientWithLogger(c Config, logger logr.Logger) (*Client, error) {
	// TODO: Validate Config
	client := &Client{
		NodeID:  strcase.ToDelimited(c.NodeID, uint8(0x2d)),
		config:  c,
		devices: make(map[string]*device.Device),
		logger:  logger,
	}
	return client, nil
}

func (c *Client) Logger() *logr.Logger {
	return &c.logger
}

var (
	ErrNoDeviceFound = fmt.Errorf("no device found")
)

func (c *Client) Connect() error {
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
		c.logger.Info(fmt.Sprintf("Received message on Default Handler: %s from topic: %s", msg.Payload(), filepath.Base(msg.Topic())))
	})

	mqtt.ERROR = newLogError(c.logger, 1)
	mqtt.CRITICAL = newLogError(c.logger, 5)
	mqtt.WARN = newLogInfo(c.logger, 8)
	mqtt.DEBUG = newLogInfo(c.logger, 10)

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
		c.logger.V(0).Info("MQTT connection lost")
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

// Subscribe is a pass-through of MQTT Subscribe
func (c *Client) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	return c.mqtt.Subscribe(topic, qos, callback)
}

// Unsubscribe is a pass-through of MQTT Unsubscribe
func (c *Client) Unsubscribe(topics ...string) mqtt.Token {
	return c.mqtt.Unsubscribe(topics...)
}

func (c *Client) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return c.mqtt.Publish(topic, qos, retained, payload)
}

func defaultLogger() logr.Logger {
	log := funcr.New(
		func(pfx, args string) { fmt.Println(pfx, args) },
		funcr.Options{
			LogCaller:    funcr.None,
			LogTimestamp: true,
			Verbosity:    1,
		})
	return log.WithName("hass-ws")
}
