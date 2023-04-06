package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type TextOptions struct {
	states                 TextState // External state update location
	availabilityMode       string    // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	commandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	commandFunc            func(mqtt.Message, mqtt.Client)
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	jsonAttributesFunc     func() string
	max                    int    // "The maximum size of a text being set or received (maximum is 255)."
	min                    int    // "The minimum size of a text being set or received."
	mode                   string // "The mode off the text entity. Must be either `text` or `password`."
	name                   string // "The name of the text entity."
	objectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	pattern                string // "A valid regular expression the text being set or received must match with."
	qos                    int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	retain                 bool   // "If the published message should have the retain flag on or not."
	stateFunc              func() string
	uniqueId               string // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception."
	valueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the text state value from the payload received on `state_topic`."
}

func NewTextOptions() *TextOptions {
	return &TextOptions{}
}
func (o *TextOptions) States() *TextState {
	return &o.states
}
func (o *TextOptions) AvailabilityMode(mode string) *TextOptions {
	o.availabilityMode = mode
	return o
}
func (o *TextOptions) AvailabilityTemplate(template string) *TextOptions {
	o.availabilityTemplate = template
	return o
}
func (o *TextOptions) AvailabilityFunc(f func() string) *TextOptions {
	o.availabilityFunc = f
	return o
}
func (o *TextOptions) CommandTemplate(template string) *TextOptions {
	o.commandTemplate = template
	return o
}
func (o *TextOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *TextOptions {
	o.commandFunc = f
	return o
}
func (o *TextOptions) EnabledByDefault(d bool) *TextOptions {
	o.enabledByDefault = d
	return o
}
func (o *TextOptions) Encoding(encoding string) *TextOptions {
	o.encoding = encoding
	return o
}
func (o *TextOptions) EntityCategory(category string) *TextOptions {
	o.entityCategory = category
	return o
}
func (o *TextOptions) JsonAttributesTemplate(template string) *TextOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *TextOptions) JsonAttributesFunc(f func() string) *TextOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *TextOptions) Max(max int) *TextOptions {
	o.max = max
	return o
}
func (o *TextOptions) Min(min int) *TextOptions {
	o.min = min
	return o
}
func (o *TextOptions) Mode(mode string) *TextOptions {
	o.mode = mode
	return o
}
func (o *TextOptions) Name(name string) *TextOptions {
	o.name = name
	return o
}
func (o *TextOptions) ObjectId(id string) *TextOptions {
	o.objectId = id
	return o
}
func (o *TextOptions) Pattern(pattern string) *TextOptions {
	o.pattern = pattern
	return o
}
func (o *TextOptions) Qos(qos int) *TextOptions {
	o.qos = qos
	return o
}
func (o *TextOptions) Retain(retain bool) *TextOptions {
	o.retain = retain
	return o
}
func (o *TextOptions) StateFunc(f func() string) *TextOptions {
	o.stateFunc = f
	return o
}
func (o *TextOptions) UniqueId(id string) *TextOptions {
	o.uniqueId = id
	return o
}
func (o *TextOptions) ValueTemplate(template string) *TextOptions {
	o.valueTemplate = template
	return o
}
