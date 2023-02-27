package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iancoleman/strcase"
	hass_mqtt "github.com/kjbreil/hass-mqtt"
	"github.com/kjbreil/hass-mqtt/device"
	"github.com/kjbreil/hass-mqtt/entities"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	c := hass_mqtt.Config{
		NodeID: "hass_mqtt_example",
		MQTT: struct {
			Host     string  `json:"host"`
			Port     int     `json:"port"`
			SSL      bool    `json:"ssl"`
			UserName *string `json:"user_name,omitempty"`
			Password *string `json:"password,omitempty"`
		}{
			Host: "192.168.1.12",
			Port: 1883,
		},
	}

	client, _ := hass_mqtt.NewClient(c)

	_ = client.Add(device.New("Purifier", "some_awesome_purifier", "Air Pure 1000", "Kaygel", "0.0.1"))

	fanName := "Some Awesome Purifier"
	uniqueId := strcase.ToDelimited(fmt.Sprintf("%s", fanName), uint8(0x2d))

	err := client.Get("some_awesome_purifier").Add(
		&entities.Fan{
			Name:     &fanName,
			UniqueId: &uniqueId,
			CommandFunc: func(message mqtt.Message, client mqtt.Client) {

			},
		})
	if err != nil {
		log.Panicln(err)
	}

	err = client.Connect()
	if err != nil {
		log.Panicln(err)

	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Everything is set up")
	<-done

	client.Disconnect()

}
