package entities

import (
	"encoding/json"
	strcase "github.com/iancoleman/strcase"
	common "github.com/kjbreil/hass-mqtt/common"
	"log"
	"reflect"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type BinarySensor struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive birth and LWT messages from the MQTT device. If `availability` is not defined, the binary sensor will always be considered `available` and its state will be `on`, `off` or `unknown`. If `availability` is defined, the binary sensor will be considered as `unavailable` by default and the sensor's initial state will be `unavailable`. Must not be used together with `availability`."
	availabilityFunc       func() string
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	DeviceClass            *string `json:"device_class,omitempty"`             // "Sets the [class of the device](/integrations/binary_sensor/#device-class), changing the device state and icon that is displayed on the frontend."
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	ExpireAfter            *int    `json:"expire_after,omitempty"`             // "If set, it defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`. Default the sensors state never expires."
	ForceUpdate            *bool   `json:"force_update,omitempty"`             // "Sends update events (which results in update of [state object](/docs/configuration/state_object/)'s `last_changed`) even if the sensor's state hasn't changed. Useful if you want to have meaningful value graphs in history or want to create an automation that triggers on *every* incoming state message (not only when the sensor's new state is different to the current one)."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                  // "The name of the binary sensor."
	ObjectId               *string `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	OffDelay               *int    `json:"off_delay,omitempty"`             // "For sensors that only send `on` state updates (like PIRs), this variable sets a delay in seconds after which the sensor's state will be updated back to `off`."
	PayloadAvailable       *string `json:"payload_available,omitempty"`     // "The string that represents the `online` state."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"` // "The string that represents the `offline` state."
	PayloadOff             *string `json:"payload_off,omitempty"`           // "The string that represents the `off` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	PayloadOn              *string `json:"payload_on,omitempty"`            // "The string that represents the `on` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	Qos                    *int    `json:"qos,omitempty"`                   // "The maximum QoS level to be used when receiving messages."
	StateTopic             *string `json:"state_topic,omitempty"`           // "The MQTT topic subscribed to receive sensor's state."
	stateFunc              func() string
	UniqueId               *string            `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string            `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that returns a string to be compared to `payload_on`/`payload_off` or an empty string, in which case the MQTT message will be removed. Remove this option when `payload_on` and `payload_off` are sufficient to match your payloads (i.e no pre-processing of original message is required)."
	MQTT                   *MQTTFields        `json:"-"`                        // MQTT configuration parameters
	states                 binarySensorState  // Internal Holder of States
	States                 *BinarySensorState `json:"-"` // External state update location
}

func NewBinarySensor(o *BinarySensorOptions) *BinarySensor {
	var b BinarySensor

	b.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		b.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		b.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		b.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.deviceClass).IsZero() {
		b.DeviceClass = &o.deviceClass
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		b.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		b.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		b.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.expireAfter).IsZero() {
		b.ExpireAfter = &o.expireAfter
	}
	if !reflect.ValueOf(o.forceUpdate).IsZero() {
		b.ForceUpdate = &o.forceUpdate
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		b.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		b.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		b.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.name).IsZero() {
		b.Name = &o.name
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		b.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.offDelay).IsZero() {
		b.OffDelay = &o.offDelay
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		b.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		b.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadOff).IsZero() {
		b.PayloadOff = &o.payloadOff
	}
	if !reflect.ValueOf(o.payloadOn).IsZero() {
		b.PayloadOn = &o.payloadOn
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		b.Qos = &o.qos
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		b.stateFunc = o.stateFunc
	} else {
		b.stateFunc = func() string {
			return b.States.State
		}
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		b.UniqueId = &o.uniqueId
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		b.ValueTemplate = &o.valueTemplate
	}
	return &b
}

type binarySensorState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type BinarySensorState struct {
	JsonAttributes string
	State          string
}

func (d *BinarySensor) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *BinarySensor) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *BinarySensor) GetRawId() string {
	return "binary_sensor"
}
func (d *BinarySensor) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *BinarySensor) GetUniqueId() string {
	return *d.UniqueId
}
func (d *BinarySensor) GetName() string {
	return *d.Name
}
func (d *BinarySensor) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *BinarySensor) UpdateState() {
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
func (d *BinarySensor) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 2, true, message)
	token.WaitTimeout(common.WaitTimeout)
	d.availabilityFunc()
	d.UpdateState()
}
func (d *BinarySensor) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, false, "offline")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *BinarySensor) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *BinarySensor) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(0)
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
func (d *BinarySensor) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *BinarySensor) populateTopics() {
	if d.availabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
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
func (d *BinarySensor) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *BinarySensor) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
