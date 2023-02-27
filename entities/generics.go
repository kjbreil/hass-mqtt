package entities

import (
	"log"
	"strings"

	"github.com/kjbreil/hass-mqtt/common"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func GetTopicPrefix(e Entity) string {
	uID := e.GetUniqueId()
	if uID != "" {
		uID = uID + "/"
	}
	return NodeID + "/" + e.GetRawId() + "/" + uID
}

func GetDiscoveryTopic(e Entity) string {
	uID := e.GetUniqueId()
	if uID != "" {
		uID = uID + "/"
	}
	return DiscoveryPrefix + "/" + e.GetRawId() + "/" + NodeID + "/" + uID + "config"
}

func GetTopic(e Entity, rawTopicString string) string {
	return GetTopicPrefix(e) + strings.TrimSuffix(rawTopicString, "_topic")
}

func MakeMessageHandler(e Entity) func(client mqtt.Client, msg mqtt.Message) {

	return func(client mqtt.Client, msg mqtt.Message) {

		topicFound := false

		for topic, f := range common.TopicStore {
			if msg.Topic() == topic {
				topicFound = true
				(*f)(msg, client)
				e.UpdateState()
			}
		}

		if !topicFound {
			log.Println("Unknown Message on topic " + msg.Topic())
			log.Println(msg.Payload())
		}
	}

}

type MQTTFields struct {
	Client         *mqtt.Client
	ForceUpdate    *bool
	MessageHandler mqtt.MessageHandler
	UpdateInterval *float64
}
