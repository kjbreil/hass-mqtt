package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type ButtonOptions struct {
	states                 ButtonState // External state update location
	availabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	commandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	commandFunc            func(mqtt.Message, mqtt.Client)
	deviceClass            string // "The [type/class](/integrations/button/#device-class) of the button to set the icon in the frontend."
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the published messages."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name to use when displaying this button."
	objectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	payloadAvailable       string // "The payload that represents the available state."
	payloadNotAvailable    string // "The payload that represents the unavailable state."
	payloadPress           string // "The payload To send to trigger the button."
	qos                    int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	retain                 bool   // "If the published message should have the retain flag on or not."
	uniqueId               string // "An ID that uniquely identifies this button entity. If two buttons have the same unique ID, Home Assistant will raise an exception."
}

func NewButtonOptions() *ButtonOptions {
	return &ButtonOptions{}
}
func (o *ButtonOptions) States() *ButtonState {
	return &o.states
}
func (o *ButtonOptions) AvailabilityMode(mode string) *ButtonOptions {
	o.availabilityMode = mode
	return o
}
func (o *ButtonOptions) AvailabilityTemplate(template string) *ButtonOptions {
	o.availabilityTemplate = template
	return o
}
func (o *ButtonOptions) AvailabilityFunc(f func() string) *ButtonOptions {
	o.availabilityFunc = f
	return o
}
func (o *ButtonOptions) CommandTemplate(template string) *ButtonOptions {
	o.commandTemplate = template
	return o
}
func (o *ButtonOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *ButtonOptions {
	o.commandFunc = f
	return o
}
func (o *ButtonOptions) DeviceClass(class string) *ButtonOptions {
	o.deviceClass = class
	return o
}
func (o *ButtonOptions) EnabledByDefault(d bool) *ButtonOptions {
	o.enabledByDefault = d
	return o
}
func (o *ButtonOptions) Encoding(encoding string) *ButtonOptions {
	o.encoding = encoding
	return o
}
func (o *ButtonOptions) EntityCategory(category string) *ButtonOptions {
	o.entityCategory = category
	return o
}
func (o *ButtonOptions) Icon(icon string) *ButtonOptions {
	o.icon = icon
	return o
}
func (o *ButtonOptions) JsonAttributesTemplate(template string) *ButtonOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *ButtonOptions) JsonAttributesFunc(f func() string) *ButtonOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *ButtonOptions) Name(name string) *ButtonOptions {
	o.name = name
	return o
}
func (o *ButtonOptions) ObjectId(id string) *ButtonOptions {
	o.objectId = id
	return o
}
func (o *ButtonOptions) PayloadAvailable(available string) *ButtonOptions {
	o.payloadAvailable = available
	return o
}
func (o *ButtonOptions) PayloadNotAvailable(available string) *ButtonOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *ButtonOptions) PayloadPress(press string) *ButtonOptions {
	o.payloadPress = press
	return o
}
func (o *ButtonOptions) Qos(qos int) *ButtonOptions {
	o.qos = qos
	return o
}
func (o *ButtonOptions) Retain(retain bool) *ButtonOptions {
	o.retain = retain
	return o
}
func (o *ButtonOptions) UniqueId(id string) *ButtonOptions {
	o.uniqueId = id
	return o
}
