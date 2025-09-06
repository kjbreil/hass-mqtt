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
type DeviceTracker struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary message containing device tracker attributes. This topic can be used to set the location of the device tracker under the following conditions:\n- If the attributes in the JSON message include `longitude`, `latitude`, and `gps_accuracy` (optional).\n - If the device tracker is within a configured [zone](/integrations/zone/).\n\nIf these conditions are met, it is not required to configure `state_topic`.\n\n Be aware that any location message received at `state_topic`  overrides the location received via `json_attributes_topic` until a message configured with `payload_reset` is received at `state_topic`. For a more generic usage example of the `json_attributes_topic`, refer to the [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                  // "The name of the MQTT device_tracker."
	ObjectId               *string `json:"object_id,omitempty"`             // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	PayloadAvailable       *string `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadHome            *string `json:"payload_home,omitempty"`          // "The payload value that represents the 'home' state for the device."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadNotHome         *string `json:"payload_not_home,omitempty"`      // "The payload value that represents the 'not_home' state for the device."
	PayloadReset           *string `json:"payload_reset,omitempty"`         // "The payload value that will have the device's location automatically derived from Home Assistant's zones."
	Platform               *string `json:"platform,omitempty"`              // "Must be `device_tracker`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	Qos                    *int    `json:"qos,omitempty"`                   // "The maximum QoS level to be used when receiving and publishing messages."
	SourceType             *string `json:"source_type,omitempty"`           // "Attribute of a device tracker that affects state when being used to track a [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`, or `bluetooth_le`."
	StateTopic             *string `json:"state_topic,omitempty"`           // "The MQTT topic subscribed to receive device tracker state changes. The states defined in `state_topic` override the location states defined by the `json_attributes_topic`. This state override is turned inactive if the `state_topic` receives a message containing `payload_reset`. The `state_topic` can only be omitted if `json_attributes_topic` is used. An empty payload is ignored. Valid payloads are `not_home`, `home` or any other custom location or zone name. Payloads for `not_home`, `home` can be overridden with the `payload_not_home`and `payload_home` config options."
	stateFunc              func() string
	UniqueId               *string             `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this device_tracker. If two device_trackers have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	ValueTemplate          *string             `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) that returns a device tracker state."
	MQTT                   *MQTTFields         `json:"-"`                        // MQTT configuration parameters
	states                 deviceTrackerState  // Internal Holder of States
	States                 *DeviceTrackerState `json:"-"` // External state update location
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
func NewDeviceTracker(o *DeviceTrackerOptions) (*DeviceTracker, error) {
	var d DeviceTracker

	d.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		d.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		d.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		d.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		d.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		d.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		d.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.name).IsZero() {
		d.Name = &o.name
	} else {
		return nil, fmt.Errorf("name not defined")
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		d.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		d.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadHome).IsZero() {
		d.PayloadHome = &o.payloadHome
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		d.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadNotHome).IsZero() {
		d.PayloadNotHome = &o.payloadNotHome
	}
	if !reflect.ValueOf(o.payloadReset).IsZero() {
		d.PayloadReset = &o.payloadReset
	}
	if !reflect.ValueOf(o.platform).IsZero() {
		d.Platform = &o.platform
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		d.Qos = &o.qos
	}
	if !reflect.ValueOf(o.sourceType).IsZero() {
		d.SourceType = &o.sourceType
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		d.stateFunc = o.stateFunc
	} else {
		d.stateFunc = func() string {
			return d.States.State
		}
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		d.UniqueId = &o.uniqueId
	} else {
		uniqueId := strcase.ToDelimited(o.name, uint8(0x2d))
		uniqueId = strings.ReplaceAll(uniqueId, "'", "_")
		d.UniqueId = &uniqueId
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		d.ValueTemplate = &o.valueTemplate
	}
	return &d, nil
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

func (d *DeviceTracker) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *DeviceTracker) State(s string) {
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
func (d *DeviceTracker) GetDomainEntity() string {
	return fmt.Sprintf("device_tracker.%s", strings.ReplaceAll(*d.UniqueId, "-", "_"))
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
