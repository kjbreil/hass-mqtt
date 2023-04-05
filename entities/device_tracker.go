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
type DeviceTracker struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as device_tracker attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                  // "The name of the MQTT device_tracker."
	ObjectId               *string `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       *string `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadHome            *string `json:"payload_home,omitempty"`          // "The payload value that represents the 'home' state for the device."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadNotHome         *string `json:"payload_not_home,omitempty"`      // "The payload value that represents the 'not_home' state for the device."
	PayloadReset           *string `json:"payload_reset,omitempty"`         // "The payload value that will have the device's location automatically derived from Home Assistant's zones."
	Qos                    *int    `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic."
	SourceType             *string `json:"source_type,omitempty"`           // "Attribute of a device tracker that affects state when being used to track a [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`, or `bluetooth_le`."
	StateTopic             *string `json:"state_topic,omitempty"`           // "The MQTT topic subscribed to receive device tracker state changes."
	stateFunc              func() string
	UniqueId               *string             `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this device_tracker. If two device_trackers have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string             `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that returns a device tracker state."
	MQTT                   *MQTTFields         `json:"-"`                        // MQTT configuration parameters
	states                 deviceTrackerState  // Internal Holder of States
	States                 *DeviceTrackerState `json:"-"` // External state update location
}

func NewDeviceTracker(o *DeviceTrackerOptions) *DeviceTracker {
	var d DeviceTracker

	d.States = &o.States
	if !reflect.ValueOf(o.AvailabilityMode).IsZero() {
		d.AvailabilityMode = &o.AvailabilityMode
	}
	if !reflect.ValueOf(o.AvailabilityTemplate).IsZero() {
		d.AvailabilityTemplate = &o.AvailabilityTemplate
	}
	if !reflect.ValueOf(o.AvailabilityFunc).IsZero() {
		d.availabilityFunc = o.AvailabilityFunc
	}
	if !reflect.ValueOf(o.Icon).IsZero() {
		d.Icon = &o.Icon
	}
	if !reflect.ValueOf(o.JsonAttributesTemplate).IsZero() {
		d.JsonAttributesTemplate = &o.JsonAttributesTemplate
	}
	if !reflect.ValueOf(o.JsonAttributesFunc).IsZero() {
		d.jsonAttributesFunc = o.JsonAttributesFunc
	}
	if !reflect.ValueOf(o.Name).IsZero() {
		d.Name = &o.Name
	}
	if !reflect.ValueOf(o.ObjectId).IsZero() {
		d.ObjectId = &o.ObjectId
	}
	if !reflect.ValueOf(o.PayloadAvailable).IsZero() {
		d.PayloadAvailable = &o.PayloadAvailable
	}
	if !reflect.ValueOf(o.PayloadHome).IsZero() {
		d.PayloadHome = &o.PayloadHome
	}
	if !reflect.ValueOf(o.PayloadNotAvailable).IsZero() {
		d.PayloadNotAvailable = &o.PayloadNotAvailable
	}
	if !reflect.ValueOf(o.PayloadNotHome).IsZero() {
		d.PayloadNotHome = &o.PayloadNotHome
	}
	if !reflect.ValueOf(o.PayloadReset).IsZero() {
		d.PayloadReset = &o.PayloadReset
	}
	if !reflect.ValueOf(o.Qos).IsZero() {
		d.Qos = &o.Qos
	}
	if !reflect.ValueOf(o.SourceType).IsZero() {
		d.SourceType = &o.SourceType
	}
	if !reflect.ValueOf(o.StateFunc).IsZero() {
		d.stateFunc = o.StateFunc
	} else {
		d.stateFunc = func() string {
			return d.States.State
		}
	}
	if !reflect.ValueOf(o.UniqueId).IsZero() {
		d.UniqueId = &o.UniqueId
	}
	if !reflect.ValueOf(o.ValueTemplate).IsZero() {
		d.ValueTemplate = &o.ValueTemplate
	}
	return &d
}

type deviceTrackerState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type DeviceTrackerState struct {
	JsonAttributes string
	State          string
}

func (d *DeviceTracker) SetJsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *DeviceTracker) SetState(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *DeviceTracker) GetRawId() string {
	return "device_tracker"
}
func (d *DeviceTracker) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *DeviceTracker) GetUniqueId() string {
	return *d.UniqueId
}
func (d *DeviceTracker) GetName() string {
	return *d.Name
}
func (d *DeviceTracker) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *DeviceTracker) UpdateState() {
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
func (d *DeviceTracker) Subscribe() {
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
func (d *DeviceTracker) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, false, "offline")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *DeviceTracker) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *DeviceTracker) Initialize() {
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
func (d *DeviceTracker) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *DeviceTracker) populateTopics() {
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
func (d *DeviceTracker) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *DeviceTracker) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
