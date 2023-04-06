package entities

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	strcase "github.com/iancoleman/strcase"
	common "github.com/kjbreil/hass-mqtt/common"
	"log"
	"reflect"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Scene struct {
	AvailabilityMode     *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc     func() string
	CommandTopic         *string `json:"command_topic,omitempty"` // "The MQTT topic to publish `payload_on` to activate the scene."
	commandFunc          func(mqtt.Message, mqtt.Client)
	EnabledByDefault     *bool       `json:"enabled_by_default,omitempty"`    // "Flag which defines if the entity should be enabled when first added."
	EntityCategory       *string     `json:"entity_category,omitempty"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                 *string     `json:"icon,omitempty"`                  // "Icon for the scene."
	Name                 *string     `json:"name,omitempty"`                  // "The name to use when displaying this scene."
	ObjectId             *string     `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable     *string     `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable  *string     `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOn            *string     `json:"payload_on,omitempty"`            // "The payload that will be sent to `command_topic` when activating the MQTT scene."
	Qos                  *int        `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain               *bool       `json:"retain,omitempty"`                // "If the published message should have the retain flag on or not."
	UniqueId             *string     `json:"unique_id,omitempty"`             // "An ID that uniquely identifies this scene entity. If two scenes have the same unique ID, Home Assistant will raise an exception."
	MQTT                 *MQTTFields `json:"-"`                               // MQTT configuration parameters
	states               sceneState  // Internal Holder of States
	States               *SceneState `json:"-"` // External state update location
}

func NewScene(o *SceneOptions) (*Scene, error) {
	var s Scene

	s.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		s.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		s.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		s.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		s.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		s.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		s.Icon = &o.icon
	}
	if !reflect.ValueOf(o.name).IsZero() {
		s.Name = &o.name
	} else {
		return nil, fmt.Errorf("name not defined")
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		s.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		s.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		s.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadOn).IsZero() {
		s.PayloadOn = &o.payloadOn
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		s.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		s.Retain = &o.retain
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		s.UniqueId = &o.uniqueId
	} else {
		uniqueId := strcase.ToDelimited(o.name, uint8(0x2d))
		s.UniqueId = &uniqueId
	}
	return &s, nil
}

type sceneState struct {
	Availability *string
}
type SceneState struct{}

func (d *Scene) GetRawId() string {
	return "scene"
}
func (d *Scene) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Scene) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Scene) GetName() string {
	return *d.Name
}
func (d *Scene) PopulateDevice(_ string, _ string, _ string, _ string, _ string) {}
func (d *Scene) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.availabilityFunc()
		if d.states.Availability == nil || state != *d.states.Availability || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
		}
	}
}
func (d *Scene) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.CommandTopic != nil {
		t := c.Subscribe(*d.CommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	token := c.Publish(GetDiscoveryTopic(d), 2, true, message)
	token.WaitTimeout(common.WaitTimeout)
	d.availabilityFunc()
	d.UpdateState()
}
func (d *Scene) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, false, "offline")
	token.WaitTimeout(common.WaitTimeout)
	if d.CommandTopic != nil {
		t := c.Unsubscribe(*d.CommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Scene) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Scene) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(0)
	}
	if d.Retain == nil {
		d.Retain = new(bool)
		*d.Retain = false
	}
	if d.UniqueId == nil {
		d.UniqueId = new(string)
		*d.UniqueId = strcase.ToDelimited(*d.Name, uint8(0x2d))
	}
	if d.availabilityFunc == nil {
		d.availabilityFunc = AvailabilityFunc
	}
	if d.MQTT == nil {
		d.MQTT = new(MQTTFields)
	}
	d.AddMessageHandler()
	d.populateTopics()
}
func (d *Scene) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Scene) populateTopics() {
	if d.availabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.commandFunc != nil {
		d.CommandTopic = new(string)
		*d.CommandTopic = GetTopic(d, "command_topic")
		common.TopicStore[*d.CommandTopic] = &d.commandFunc
	}
}
func (d *Scene) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Scene) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
