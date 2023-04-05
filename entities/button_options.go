package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type ButtonOptions struct {
	States                 ButtonState // External state update location
	AvailabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	CommandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandFunc            func(mqtt.Message, mqtt.Client)
	DeviceClass            string // "The [type/class](/integrations/button/#device-class) of the button to set the icon in the frontend."
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the published messages."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc     func() string
	Name                   string // "The name to use when displaying this button."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       string // "The payload that represents the available state."
	PayloadNotAvailable    string // "The payload that represents the unavailable state."
	PayloadPress           string // "The payload To send to trigger the button."
	Qos                    int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 bool   // "If the published message should have the retain flag on or not."
	UniqueId               string // "An ID that uniquely identifies this button entity. If two buttons have the same unique ID, Home Assistant will raise an exception."
}

func NewButtonOptions() *ButtonOptions {
	return &ButtonOptions{}
}
func (o *ButtonOptions) GetStates() *ButtonState {
	return &o.States
}
func (o *ButtonOptions) SetAvailabilityMode(mode string) *ButtonOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *ButtonOptions) SetAvailabilityTemplate(template string) *ButtonOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *ButtonOptions) SetAvailabilityFunc(f func() string) *ButtonOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *ButtonOptions) SetCommandTemplate(template string) *ButtonOptions {
	o.CommandTemplate = template
	return o
}
func (o *ButtonOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *ButtonOptions {
	o.CommandFunc = f
	return o
}
func (o *ButtonOptions) SetDeviceClass(class string) *ButtonOptions {
	o.DeviceClass = class
	return o
}
func (o *ButtonOptions) SetEnabledByDefault(d bool) *ButtonOptions {
	o.EnabledByDefault = d
	return o
}
func (o *ButtonOptions) SetEncoding(encoding string) *ButtonOptions {
	o.Encoding = encoding
	return o
}
func (o *ButtonOptions) SetEntityCategory(category string) *ButtonOptions {
	o.EntityCategory = category
	return o
}
func (o *ButtonOptions) SetIcon(icon string) *ButtonOptions {
	o.Icon = icon
	return o
}
func (o *ButtonOptions) SetJsonAttributesTemplate(template string) *ButtonOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *ButtonOptions) SetJsonAttributesFunc(f func() string) *ButtonOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *ButtonOptions) SetName(name string) *ButtonOptions {
	o.Name = name
	return o
}
func (o *ButtonOptions) SetObjectId(id string) *ButtonOptions {
	o.ObjectId = id
	return o
}
func (o *ButtonOptions) SetPayloadAvailable(available string) *ButtonOptions {
	o.PayloadAvailable = available
	return o
}
func (o *ButtonOptions) SetPayloadNotAvailable(available string) *ButtonOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *ButtonOptions) SetPayloadPress(press string) *ButtonOptions {
	o.PayloadPress = press
	return o
}
func (o *ButtonOptions) SetQos(qos int) *ButtonOptions {
	o.Qos = qos
	return o
}
func (o *ButtonOptions) SetRetain(retain bool) *ButtonOptions {
	o.Retain = retain
	return o
}
func (o *ButtonOptions) SetUniqueId(id string) *ButtonOptions {
	o.UniqueId = id
	return o
}
