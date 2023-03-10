package entities

import (
	"encoding/json"
	strcase "github.com/iancoleman/strcase"
	common "github.com/kjbreil/hass-mqtt/common"
	"log"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Camera struct {
	AvailabilityMode     *string       `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string       `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string       `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc     func() string `json:"-"`
	Device               struct {
		ConfigurationUrl *string `json:"configuration_url,omitempty"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      *string `json:"connections,omitempty"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [\"mac\", \"02:5b:26:a8:dc:12\"]`."
		Identifiers      *string `json:"identifiers,omitempty"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     *string `json:"manufacturer,omitempty"`      // "The manufacturer of the device."
		Model            *string `json:"model,omitempty"`             // "The model of the device."
		Name             *string `json:"name,omitempty"`              // "The name of the device."
		SuggestedArea    *string `json:"suggested_area,omitempty"`    // "Suggest an area if the device isn???t in one yet."
		SwVersion        *string `json:"sw_version,omitempty"`        // "The firmware version of the device."
		ViaDevice        *string `json:"via_device,omitempty"`        // "Identifier of a device that routes messages between this device and Home Assistant. Examples of such devices are hubs, or parent devices of a sub-device. This is used to show device topology in Home Assistant."
	} `json:"device,omitempty"` // Device configuration parameters
	EnabledByDefault       *bool         `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string       `json:"encoding,omitempty"`                 // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload. Use `image_encoding` to enable `Base64` decoding on `topic`."
	EntityCategory         *string       `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string       `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	ImageEncoding          *string       `json:"image_encoding,omitempty"`           // "The encoding of the image payloads received. Set to `\"b64\"` to enable base64 decoding of image payload. If not set, the image payload must be raw binary data."
	JsonAttributesTemplate *string       `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic    *string       `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies `force_update` of the current sensor state when a message is received on this topic."
	JsonAttributesFunc     func() string `json:"-"`
	Name                   *string       `json:"name,omitempty"`      // "The name of the camera."
	ObjectId               *string       `json:"object_id,omitempty"` // "Used instead of `name` for automatic generation of `entity_id`"
	StateTopic             *string       `json:"topic,omitempty"`     // "The MQTT topic to subscribe to."
	StateFunc              func() string `json:"-"`
	UniqueId               *string       `json:"unique_id,omitempty"` // "An ID that uniquely identifies this camera. If two cameras have the same unique ID Home Assistant will raise an exception."
	MQTT                   *MQTTFields   `json:"-"`                   // MQTT configuration parameters
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
func (d *Camera) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Camera) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Camera.Availability[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, common.QoS, common.Retain, state)
			stateStore.Camera.Availability[d.GetUniqueId()] = state
			token.WaitTimeout(common.WaitTimeout)
		}
	}
	if d.JsonAttributesTopic != nil {
		state := d.JsonAttributesFunc()
		if state != stateStore.Camera.JsonAttributes[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.JsonAttributesTopic, common.QoS, common.Retain, state)
			stateStore.Camera.JsonAttributes[d.GetUniqueId()] = state
			token.WaitTimeout(common.WaitTimeout)
		}
	}
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.Camera.State[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Camera.State[d.GetUniqueId()] = state
			token.WaitTimeout(common.WaitTimeout)
		}
	}
}
func (d *Camera) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.WaitTimeout(common.WaitTimeout)
	d.AvailabilityFunc()
	d.UpdateState()
}
func (d *Camera) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Camera) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Camera) Initialize() {
	if d.UniqueId == nil {
		d.UniqueId = new(string)
		*d.UniqueId = strcase.ToDelimited(*d.Name, uint8(0x2d))
	}
	if d.AvailabilityFunc == nil {
		d.AvailabilityFunc = AvailabilityFunc
	}
	if d.MQTT == nil {
		d.MQTT = new(MQTTFields)
	}
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Camera) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Camera) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.JsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
	}
	if d.StateFunc != nil {
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
