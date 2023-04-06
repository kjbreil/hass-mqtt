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
type Number struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`  // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"` // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	CommandTemplate        *string `json:"command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandTopic           *string `json:"command_topic,omitempty"`    // "The MQTT topic to publish commands to change the number."
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	DeviceClass            *string `json:"device_class,omitempty"`             // "The [type/class](/integrations/number/#device-class) of the number."
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as number attributes. Implies `force_update` of the current number state when a message is received on this topic."
	jsonAttributesFunc     func() string
	Max                    *float64 `json:"max,omitempty"`           // "Maximum value."
	Min                    *float64 `json:"min,omitempty"`           // "Minimum value."
	Mode                   *string  `json:"mode,omitempty"`          // "Control how the number should be displayed in the UI. Can be set to `box` or `slider` to force a display mode."
	Name                   *string  `json:"name,omitempty"`          // "The name of the Number."
	ObjectId               *string  `json:"object_id,omitempty"`     // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool    `json:"optimistic,omitempty"`    // "Flag that defines if number works in optimistic mode."
	PayloadReset           *string  `json:"payload_reset,omitempty"` // "A special payload that resets the state to `None` when received on the `state_topic`."
	Qos                    *int     `json:"qos,omitempty"`           // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 *bool    `json:"retain,omitempty"`        // "If the published message should have the retain flag on or not."
	StateTopic             *string  `json:"state_topic,omitempty"`   // "The MQTT topic subscribed to receive number values."
	stateFunc              func() string
	Step                   *float64     `json:"step,omitempty"`                // "Step value. Smallest value `0.001`."
	UniqueId               *string      `json:"unique_id,omitempty"`           // "An ID that uniquely identifies this Number. If two Numbers have the same unique ID Home Assistant will raise an exception."
	UnitOfMeasurement      *string      `json:"unit_of_measurement,omitempty"` // "Defines the unit of measurement of the sensor, if any."
	ValueTemplate          *string      `json:"value_template,omitempty"`      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value."
	MQTT                   *MQTTFields  `json:"-"`                             // MQTT configuration parameters
	states                 numberState  // Internal Holder of States
	States                 *NumberState `json:"-"` // External state update location
}

func NewNumber(o *NumberOptions) *Number {
	var n Number

	n.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		n.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		n.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.commandTemplate).IsZero() {
		n.CommandTemplate = &o.commandTemplate
	}
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		n.commandFunc = o.commandFunc
	} else {
		n.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.deviceClass).IsZero() {
		n.DeviceClass = &o.deviceClass
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		n.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		n.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		n.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		n.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		n.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		n.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.max).IsZero() {
		n.Max = &o.max
	}
	if !reflect.ValueOf(o.min).IsZero() {
		n.Min = &o.min
	}
	if !reflect.ValueOf(o.mode).IsZero() {
		n.Mode = &o.mode
	}
	if !reflect.ValueOf(o.name).IsZero() {
		n.Name = &o.name
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		n.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.optimistic).IsZero() {
		n.Optimistic = &o.optimistic
	}
	if !reflect.ValueOf(o.payloadReset).IsZero() {
		n.PayloadReset = &o.payloadReset
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		n.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		n.Retain = &o.retain
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		n.stateFunc = o.stateFunc
	} else {
		n.stateFunc = func() string {
			return n.States.State
		}
	}
	if !reflect.ValueOf(o.step).IsZero() {
		n.Step = &o.step
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		n.UniqueId = &o.uniqueId
	}
	if !reflect.ValueOf(o.unitOfMeasurement).IsZero() {
		n.UnitOfMeasurement = &o.unitOfMeasurement
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		n.ValueTemplate = &o.valueTemplate
	}
	return &n
}

type numberState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type NumberState struct {
	JsonAttributes string
	State          string
}

func (d *Number) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Number) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Number) GetRawId() string {
	return "number"
}
func (d *Number) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Number) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Number) GetName() string {
	return *d.Name
}
func (d *Number) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Number) UpdateState() {
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
func (d *Number) Subscribe() {
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
func (d *Number) UnSubscribe() {
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
func (d *Number) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Number) Initialize() {
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
func (d *Number) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Number) populateTopics() {
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
func (d *Number) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Number) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
