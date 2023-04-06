package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	hass_mqtt "github.com/kjbreil/hass-mqtt"
	"github.com/kjbreil/hass-mqtt/device"
	"github.com/kjbreil/hass-mqtt/entities"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	c := hass_mqtt.Config{
		NodeID: "hass_mqtt_example",
		MQTT: struct {
			Host     string  `json:"host" yaml:"host"`
			Port     int     `json:"port" yaml:"port"`
			SSL      bool    `json:"ssl" yaml:"ssl"`
			UserName *string `json:"user_name,omitempty" yaml:"user_name"`
			Password *string `json:"password,omitempty" yaml:"password"`
		}{
			Host:     "192.168.1.12",
			Port:     1883,
			SSL:      false,
			UserName: nil,
			Password: nil,
		},
	}

	client, _ := hass_mqtt.NewClient(c)

	//_ = client.Add(device.New("Purifier", "some_awesome_purifier", "Air Pure 1000", "Kaygel", "0.0.1"))
	//
	//fanName := "Some Awesome Purifier"
	//uniqueId := strcase.ToDelimited(fmt.Sprintf("%s", fanName), uint8(0x2d))
	//
	//err := client.Get("some_awesome_purifier").Add(
	//	&entities.Fan{
	//		Name:     &fanName,
	//		UniqueId: &uniqueId,
	//		CommandFunc: func(message mqtt.Message, client mqtt.Client) {
	//		},
	//	})
	//if err != nil {
	//	log.Panicln(err)
	//}

	d := device.New("TEST LIGHT", "test_light_mqtt", "Ligherifric", "Kaygel", "0.0.1")

	_ = client.Add(d)

	lightName := "TEST LIGHT"
	uniqueId := strcase.ToDelimited(fmt.Sprintf("%s", lightName), uint8(0x2d))

	type lightState struct {
		State    string    `json:"state"`
		LastSeen time.Time `json:"last_seen"`
	}

	lo := entities.NewLightOptions()

	lo.States().Brightness = "100"

	lo.Name("Test Light").
		UniqueId(uniqueId).
		EnableBrightness()

	l := entities.NewLight(lo)

	err := client.Get("test_light_mqtt").Add(l)

	if err != nil {
		log.Panicln(err)
	}

	err = client.Connect()
	if err != nil {
		log.Panicln(err)

	}
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			l.UpdateState()
			//i, _ := strconv.Atoi(l.States.Brightness)
			//i++
			//l.SetBrightness(fmt.Sprintf("%d", i))
		}
	}()
	l.UpdateState()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Everything is set up")
	<-done

	client.Disconnect()

}
