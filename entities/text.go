package entities

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	strcase "github.com/iancoleman/strcase"
	common "github.com/kjbreil/hass-mqtt/common"
	"log"
	"reflect"
	"strings"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Text struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	CommandTemplate        *string `json:"command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandTopic           *string `json:"command_topic,omitempty"`    // "The MQTT topic to publish the text value that is set."
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as entity attributes. Implies `force_update` of the current select state when a message is received on this topic."
	jsonAttributesFunc     func() string
	Max                    *int    `json:"max,omitempty"`         // "The maximum size of a text being set or received (maximum is 255)."
	Min                    *int    `json:"min,omitempty"`         // "The minimum size of a text being set or received."
	Mode                   *string `json:"mode,omitempty"`        // "The mode off the text entity. Must be either `text` or `password`."
	Name                   *string `json:"name,omitempty"`        // "The name of the text entity."
	ObjectId               *string `json:"object_id,omitempty"`   // "Used instead of `name` for automatic generation of `entity_id`"
	Pattern                *string `json:"pattern,omitempty"`     // "A valid regular expression the text being set or received must match with."
	Qos                    *int    `json:"qos,omitempty"`         // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 *bool   `json:"retain,omitempty"`      // "If the published message should have the retain flag on or not."
	StateTopic             *string `json:"state_topic,omitempty"` // "The MQTT topic subscribed to receive text state updates. Text state updates should match the `pattern` (if set) and meet the size constraints `min` and `max`. Can be used with `value_template` to render the incoming payload to a text update."
	stateFunc              func() string
	UniqueId               *string     `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the text state value from the payload received on `state_topic`."
	MQTT                   *MQTTFields `json:"-"`                        // MQTT configuration parameters
	states                 textState   // Internal Holder of States
	States                 *TextState  `json:"-"` // External state update location
}

func NewText(o *TextOptions) (*Text, error) {
	var t Text

	t.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		t.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		t.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		t.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.commandTemplate).IsZero() {
		t.CommandTemplate = &o.commandTemplate
	}
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		t.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.State == string(message.Payload()) {
				return
			}
			o.states.State = string(message.Payload())
			t.UpdateState()
			o.commandFunc(message, client)
		}
	} else {
		t.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		t.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		t.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		t.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		t.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		t.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.max).IsZero() {
		t.Max = &o.max
	}
	if !reflect.ValueOf(o.min).IsZero() {
		t.Min = &o.min
	}
	if !reflect.ValueOf(o.mode).IsZero() {
		t.Mode = &o.mode
	}
	if !reflect.ValueOf(o.name).IsZero() {
		t.Name = &o.name
	} else {
		return nil, fmt.Errorf("name not defined")
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		t.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.pattern).IsZero() {
		t.Pattern = &o.pattern
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		t.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		t.Retain = &o.retain
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		t.stateFunc = o.stateFunc
	} else {
		t.stateFunc = func() string {
			return t.States.State
		}
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		t.UniqueId = &o.uniqueId
	} else {
		uniqueId := strcase.ToDelimited(o.name, uint8(0x2d))
		uniqueId = strings.ReplaceAll(uniqueId, "'", "_")
		t.UniqueId = &uniqueId
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		t.ValueTemplate = &o.valueTemplate
	}
	return &t, nil
}

type textState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type TextState struct {
	JsonAttributes string
	State          string
}

func (d *Text) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Text) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Text) GetRawId() string {
	return "text"
}
func (d *Text) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Text) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Text) GetDomainEntity() string {
	return fmt.Sprintf("text.%s", strings.ReplaceAll(*d.UniqueId, "-", "_"))
}
func (d *Text) GetName() string {
	return *d.Name
}
func (d *Text) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Text) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.availabilityFunc()
		if d.states.Availability == nil || state != *d.states.Availability || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
		}
	}
	if d.JsonAttributesTopic != nil {
		state := d.jsonAttributesFunc()
		if d.states.JsonAttributes == nil || state != *d.states.JsonAttributes || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.JsonAttributesTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.JsonAttributes = &state
		}
	}
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
}
func (d *Text) Subscribe() {
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
func (d *Text) UnSubscribe() {
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
func (d *Text) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Text) Initialize() {
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
func (d *Text) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Text) populateTopics() {
	if d.availabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.commandFunc != nil {
		d.CommandTopic = new(string)
		*d.CommandTopic = GetTopic(d, "command_topic")
		common.TopicStore[*d.CommandTopic] = &d.commandFunc
	}
	if d.jsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
	}
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Text) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Text) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
