package entities

import (
	"encoding/json"
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
type Switch struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	CommandTopic           *string `json:"command_topic,omitempty"` // "The MQTT topic to publish commands to change the switch state."
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	DeviceClass            *string `json:"device_class,omitempty"`             // "The [type/class](/integrations/switch/#device-class) of the switch to set the icon in the frontend."
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                  // "The name to use when displaying this switch."
	ObjectId               *string `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool   `json:"optimistic,omitempty"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable       *string `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOff             *string `json:"payload_off,omitempty"`           // "The payload that represents `off` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_off` for details) and sending as `off` command to the `command_topic`."
	PayloadOn              *string `json:"payload_on,omitempty"`            // "The payload that represents `on` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_on`  for details) and sending as `on` command to the `command_topic`."
	Qos                    *int    `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 *bool   `json:"retain,omitempty"`                // "If the published message should have the retain flag on or not."
	StateOff               *string `json:"state_off,omitempty"`             // "The payload that represents the `off` state. Used when value that represents `off` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `off`."
	StateOn                *string `json:"state_on,omitempty"`              // "The payload that represents the `on` state. Used when value that represents `on` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `on`."
	StateTopic             *string `json:"state_topic,omitempty"`           // "The MQTT topic subscribed to receive state updates."
	stateFunc              func() string
	UniqueId               *string      `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this switch device. If two switches have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string      `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's state from the `state_topic`. To determine the switches's state result of this template will be compared to `state_on` and `state_off`."
	MQTT                   *MQTTFields  `json:"-"`                        // MQTT configuration parameters
	states                 switchState  // Internal Holder of States
	States                 *SwitchState `json:"-"` // External state update location
}

func NewSwitch(o *SwitchOptions) *Switch {
	var s Switch

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
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		s.commandFunc = o.commandFunc
	} else {
		s.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.deviceClass).IsZero() {
		s.DeviceClass = &o.deviceClass
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		s.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		s.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		s.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		s.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		s.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		s.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.name).IsZero() {
		s.Name = &o.name
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		s.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.optimistic).IsZero() {
		s.Optimistic = &o.optimistic
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		s.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		s.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadOff).IsZero() {
		s.PayloadOff = &o.payloadOff
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
	if !reflect.ValueOf(o.stateOff).IsZero() {
		s.StateOff = &o.stateOff
	}
	if !reflect.ValueOf(o.stateOn).IsZero() {
		s.StateOn = &o.stateOn
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		s.stateFunc = o.stateFunc
	} else {
		s.stateFunc = func() string {
			return s.States.State
		}
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		s.UniqueId = &o.uniqueId
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		s.ValueTemplate = &o.valueTemplate
	}
	return &s
}

type switchState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type SwitchState struct {
	JsonAttributes string
	State          string
}

func (d *Switch) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Switch) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Switch) GetRawId() string {
	return "switch"
}
func (d *Switch) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Switch) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Switch) GetName() string {
	return *d.Name
}
func (d *Switch) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Switch) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.availabilityFunc()
		if d.states.Availability == nil || state != *d.states.Availability || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
		}
	}
	if d.JsonAttributesTopic != nil {
		state := d.jsonAttributesFunc()
		if d.states.JsonAttributes == nil || state != *d.states.JsonAttributes || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.JsonAttributesTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.JsonAttributes = &state
		}
	}
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
}
func (d *Switch) Subscribe() {
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
func (d *Switch) UnSubscribe() {
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
func (d *Switch) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Switch) Initialize() {
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
func (d *Switch) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Switch) populateTopics() {
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
func (d *Switch) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Switch) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
