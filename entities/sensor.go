package entities

import (
	"encoding/json"
	"fmt"
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
type Sensor struct {
	AvailabilityMode          *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate      *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic         *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates."
	availabilityFunc          func() string
	Device                    Device  `json:"device,omitempty"`                   // Device configuration parameters
	DeviceClass               *string `json:"device_class,omitempty"`             // "The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend. The `device_class` can be `null`."
	EnabledByDefault          *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding                  *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory            *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity. When set, the entity category must be `diagnostic` for sensors."
	EntityPicture             *string `json:"entity_picture,omitempty"`           // "Picture URL for the entity."
	ExpireAfter               *int    `json:"expire_after,omitempty"`             // "If set, it defines the number of seconds after the sensor's state expires if it's not updated. After expiry, the sensor's state becomes `unavailable`. Default the sensors state never expires. By default, the sensor's state never expires. Note that when a sensor's value was sent retained to the MQTT broker, the last value sent will be replayed by the MQTT broker when Home Assistant restarts or is reloaded. As this could cause the sensor to become available with an expired state, it is not recommended to retain the sensor's state payload at the MQTT broker. Home Assistant will store and restore the sensor's state for you and calculate the remaining time to retain the sensor's state before it becomes unavailable."
	ForceUpdate               *bool   `json:"force_update,omitempty"`             // "Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history."
	Icon                      *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate    *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic       *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies `force_update` of the current sensor state when a message is received on this topic."
	jsonAttributesFunc        func() string
	LastResetValueTemplate    *string     `json:"last_reset_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the last_reset. When `last_reset_value_template` is set, the `state_class` option must be `total`. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes."
	Name                      *string     `json:"name,omitempty"`                      // "The name of the MQTT sensor. Can be set to `null` if only the device name is relevant."
	ObjectId                  *string     `json:"object_id,omitempty"`                 // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	Options                   *([]string) `json:"options,omitempty"`                   // "List of allowed sensor state value. An empty list is not allowed. The sensor's `device_class` must be set to `enum`. The `options` option cannot be used together with `state_class` or `unit_of_measurement`."
	PayloadAvailable          *string     `json:"payload_available,omitempty"`         // "The payload that represents the available state."
	PayloadNotAvailable       *string     `json:"payload_not_available,omitempty"`     // "The payload that represents the unavailable state."
	Platform                  *string     `json:"platform,omitempty"`                  // "Must be `sensor`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	Qos                       *int        `json:"qos,omitempty"`                       // "The maximum QoS level to be used when receiving and publishing messages."
	StateClass                *string     `json:"state_class,omitempty"`               // "The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor."
	StateTopic                *string     `json:"state_topic,omitempty"`               // "The MQTT topic subscribed to receive sensor values. If `device_class`, `state_class`, `unit_of_measurement` or `suggested_display_precision` is set, and a numeric value is expected, an empty value `''` will be ignored and will not update the state, a `'None'` value will set the sensor to an `unknown` state. If a `value_template` is used to parse a JSON payload, a `null` value in the JSON [will be rendered as](/docs/configuration/templating/#using-value-templates-with-mqtt) `'None'`. Note that the `device_class` can be `null`."
	stateFunc                 func() string
	SuggestedDisplayPrecision *int         `json:"suggested_display_precision,omitempty"` // "The number of decimals which should be used in the sensor's state after rounding."
	UniqueId                  *string      `json:"unique_id,omitempty"`                   // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	UnitOfMeasurement         *string      `json:"unit_of_measurement,omitempty"`         // "Defines the units of measurement of the sensor, if any. The `unit_of_measurement` can be `null`."
	ValueTemplate             *string      `json:"value_template,omitempty"`              // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the value. If the template throws an error, the current state will be used instead."
	MQTT                      *MQTTFields  `json:"-"`                                     // MQTT configuration parameters
	states                    sensorState  // Internal Holder of States
	States                    *SensorState `json:"-"` // External state update location
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
func NewSensor(o *SensorOptions) (*Sensor, error) {
	var s Sensor

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
	if !reflect.ValueOf(o.entityPicture).IsZero() {
		s.EntityPicture = &o.entityPicture
	}
	if !reflect.ValueOf(o.expireAfter).IsZero() {
		s.ExpireAfter = &o.expireAfter
	}
	if !reflect.ValueOf(o.forceUpdate).IsZero() {
		s.ForceUpdate = &o.forceUpdate
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
	if !reflect.ValueOf(o.lastResetValueTemplate).IsZero() {
		s.LastResetValueTemplate = &o.lastResetValueTemplate
	}
	if !reflect.ValueOf(o.name).IsZero() {
		s.Name = &o.name
	} else {
		return nil, fmt.Errorf("name not defined")
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		s.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.options).IsZero() {
		s.Options = &o.options
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		s.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		s.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.platform).IsZero() {
		s.Platform = &o.platform
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		s.Qos = &o.qos
	}
	if !reflect.ValueOf(o.stateClass).IsZero() {
		s.StateClass = &o.stateClass
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		s.stateFunc = o.stateFunc
	} else {
		s.stateFunc = func() string {
			return s.States.State
		}
	}
	if !reflect.ValueOf(o.suggestedDisplayPrecision).IsZero() {
		s.SuggestedDisplayPrecision = &o.suggestedDisplayPrecision
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		s.UniqueId = &o.uniqueId
	} else {
		uniqueId := strcase.ToDelimited(o.name, uint8(0x2d))
		uniqueId = strings.ReplaceAll(uniqueId, "'", "_")
		s.UniqueId = &uniqueId
	}
	if !reflect.ValueOf(o.unitOfMeasurement).IsZero() {
		s.UnitOfMeasurement = &o.unitOfMeasurement
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		s.ValueTemplate = &o.valueTemplate
	}
	return &s, nil
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

func (d *Sensor) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Sensor) State(s string) {
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
func (d *Sensor) GetDomainEntity() string {
	return fmt.Sprintf("sensor.%s", strings.ReplaceAll(*d.UniqueId, "-", "_"))
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
