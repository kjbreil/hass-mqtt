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
type Camera struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload. Use `image_encoding` to enable `Base64` decoding on `topic`."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	ImageEncoding          *string `json:"image_encoding,omitempty"`           // "The encoding of the image payloads received. Set to `\"b64\"` to enable base64 decoding of image payload. If not set, the image payload must be raw binary data."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies `force_update` of the current sensor state when a message is received on this topic."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`      // "The name of the camera."
	ObjectId               *string `json:"object_id,omitempty"` // "Used instead of `name` for automatic generation of `entity_id`"
	StateTopic             *string `json:"topic,omitempty"`     // "The MQTT topic to subscribe to."
	stateFunc              func() string
	UniqueId               *string      `json:"unique_id,omitempty"` // "An ID that uniquely identifies this camera. If two cameras have the same unique ID Home Assistant will raise an exception."
	MQTT                   *MQTTFields  `json:"-"`                   // MQTT configuration parameters
	states                 cameraState  // Internal Holder of States
	States                 *CameraState `json:"-"` // External state update location
}

func NewCamera(o *CameraOptions) *Camera {
	var c Camera

	c.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		c.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		c.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		c.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		c.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		c.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		c.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		c.Icon = &o.icon
	}
	if !reflect.ValueOf(o.imageEncoding).IsZero() {
		c.ImageEncoding = &o.imageEncoding
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		c.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		c.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.name).IsZero() {
		c.Name = &o.name
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		c.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		c.stateFunc = o.stateFunc
	} else {
		c.stateFunc = func() string {
			return c.States.State
		}
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		c.UniqueId = &o.uniqueId
	}
	return &c
}

type cameraState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type CameraState struct {
	JsonAttributes string
	State          string
}

func (d *Camera) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Camera) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Camera) GetRawId() string {
	return "camera"
}
func (d *Camera) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Camera) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Camera) GetName() string {
	return *d.Name
}
func (d *Camera) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Camera) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.availabilityFunc()
		if d.states.Availability == nil || state != *d.states.Availability || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, 2, true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
		}
	}
	if d.JsonAttributesTopic != nil {
		state := d.jsonAttributesFunc()
		if d.states.JsonAttributes == nil || state != *d.states.JsonAttributes || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.JsonAttributesTopic, 2, true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.JsonAttributes = &state
		}
	}
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, 2, true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
}
func (d *Camera) Subscribe() {
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
func (d *Camera) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, false, "offline")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Camera) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Camera) Initialize() {
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
func (d *Camera) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Camera) populateTopics() {
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
func (d *Camera) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Camera) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
