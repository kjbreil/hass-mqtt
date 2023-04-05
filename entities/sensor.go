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
type Sensor struct {
	AvailabilityMode          *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate      *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic         *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates."
	availabilityFunc          func() string
	Device                    Device  `json:"device,omitempty"`                   // Device configuration parameters
	DeviceClass               *string `json:"device_class,omitempty"`             // "The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend."
	EnabledByDefault          *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding                  *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory            *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	ExpireAfter               *int    `json:"expire_after,omitempty"`             // "If set, it defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`. Default the sensors state never expires."
	ForceUpdate               *bool   `json:"force_update,omitempty"`             // "Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history."
	Icon                      *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate    *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic       *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies `force_update` of the current sensor state when a message is received on this topic."
	jsonAttributesFunc        func() string
	LastResetValueTemplate    *string `json:"last_reset_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the last_reset. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes."
	Name                      *string `json:"name,omitempty"`                      // "The name of the MQTT sensor."
	ObjectId                  *string `json:"object_id,omitempty"`                 // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable          *string `json:"payload_available,omitempty"`         // "The payload that represents the available state."
	PayloadNotAvailable       *string `json:"payload_not_available,omitempty"`     // "The payload that represents the unavailable state."
	Qos                       *int    `json:"qos,omitempty"`                       // "The maximum QoS level of the state topic."
	StateClass                *string `json:"state_class,omitempty"`               // "The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor."
	StateTopic                *string `json:"state_topic,omitempty"`               // "The MQTT topic subscribed to receive sensor values. If `device_class`, `state_class`, `unit_of_measurement` or `suggested_display_precision` is set, and a numeric value is expected, an empty value `''` will be ignored and will not update the state, a `'None'` value will set the sensor to an `unknown` state."
	stateFunc                 func() string
	SuggestedDisplayPrecision *int         `json:"suggested_display_precision,omitempty"` // "The number of decimals which should be used in the sensor's state after rounding."
	UniqueId                  *string      `json:"unique_id,omitempty"`                   // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	UnitOfMeasurement         *string      `json:"unit_of_measurement,omitempty"`         // "Defines the units of measurement of the sensor, if any."
	ValueTemplate             *string      `json:"value_template,omitempty"`              // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value. If the template throws an error, the current state will be used instead."
	MQTT                      *MQTTFields  `json:"-"`                                     // MQTT configuration parameters
	states                    sensorState  // Internal Holder of States
	States                    *SensorState `json:"-"` // External state update location
}

func NewSensor(o *SensorOptions) *Sensor {
	var s Sensor

	s.States = &o.States
	if !reflect.ValueOf(o.AvailabilityMode).IsZero() {
		s.AvailabilityMode = &o.AvailabilityMode
	}
	if !reflect.ValueOf(o.AvailabilityTemplate).IsZero() {
		s.AvailabilityTemplate = &o.AvailabilityTemplate
	}
	if !reflect.ValueOf(o.AvailabilityFunc).IsZero() {
		s.availabilityFunc = o.AvailabilityFunc
	}
	if !reflect.ValueOf(o.DeviceClass).IsZero() {
		s.DeviceClass = &o.DeviceClass
	}
	if !reflect.ValueOf(o.EnabledByDefault).IsZero() {
		s.EnabledByDefault = &o.EnabledByDefault
	}
	if !reflect.ValueOf(o.Encoding).IsZero() {
		s.Encoding = &o.Encoding
	}
	if !reflect.ValueOf(o.EntityCategory).IsZero() {
		s.EntityCategory = &o.EntityCategory
	}
	if !reflect.ValueOf(o.ExpireAfter).IsZero() {
		s.ExpireAfter = &o.ExpireAfter
	}
	if !reflect.ValueOf(o.ForceUpdate).IsZero() {
		s.ForceUpdate = &o.ForceUpdate
	}
	if !reflect.ValueOf(o.Icon).IsZero() {
		s.Icon = &o.Icon
	}
	if !reflect.ValueOf(o.JsonAttributesTemplate).IsZero() {
		s.JsonAttributesTemplate = &o.JsonAttributesTemplate
	}
	if !reflect.ValueOf(o.JsonAttributesFunc).IsZero() {
		s.jsonAttributesFunc = o.JsonAttributesFunc
	}
	if !reflect.ValueOf(o.LastResetValueTemplate).IsZero() {
		s.LastResetValueTemplate = &o.LastResetValueTemplate
	}
	if !reflect.ValueOf(o.Name).IsZero() {
		s.Name = &o.Name
	}
	if !reflect.ValueOf(o.ObjectId).IsZero() {
		s.ObjectId = &o.ObjectId
	}
	if !reflect.ValueOf(o.PayloadAvailable).IsZero() {
		s.PayloadAvailable = &o.PayloadAvailable
	}
	if !reflect.ValueOf(o.PayloadNotAvailable).IsZero() {
		s.PayloadNotAvailable = &o.PayloadNotAvailable
	}
	if !reflect.ValueOf(o.Qos).IsZero() {
		s.Qos = &o.Qos
	}
	if !reflect.ValueOf(o.StateClass).IsZero() {
		s.StateClass = &o.StateClass
	}
	if !reflect.ValueOf(o.StateFunc).IsZero() {
		s.stateFunc = o.StateFunc
	} else {
		s.stateFunc = func() string {
			return s.States.State
		}
	}
	if !reflect.ValueOf(o.SuggestedDisplayPrecision).IsZero() {
		s.SuggestedDisplayPrecision = &o.SuggestedDisplayPrecision
	}
	if !reflect.ValueOf(o.UniqueId).IsZero() {
		s.UniqueId = &o.UniqueId
	}
	if !reflect.ValueOf(o.UnitOfMeasurement).IsZero() {
		s.UnitOfMeasurement = &o.UnitOfMeasurement
	}
	if !reflect.ValueOf(o.ValueTemplate).IsZero() {
		s.ValueTemplate = &o.ValueTemplate
	}
	return &s
}

type sensorState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type SensorState struct {
	JsonAttributes string
	State          string
}

func (d *Sensor) SetJsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Sensor) SetState(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Sensor) GetRawId() string {
	return "sensor"
}
func (d *Sensor) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Sensor) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Sensor) GetName() string {
	return *d.Name
}
func (d *Sensor) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Sensor) UpdateState() {
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
func (d *Sensor) Subscribe() {
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
func (d *Sensor) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, false, "offline")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Sensor) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Sensor) Initialize() {
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
func (d *Sensor) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Sensor) populateTopics() {
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
func (d *Sensor) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Sensor) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
