package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type TextOptions struct {
	States                 TextState // External state update location
	AvailabilityMode       string    // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	CommandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandFunc            func(mqtt.Message, mqtt.Client)
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesFunc     func() string
	Max                    int    // "The maximum size of a text being set or received (maximum is 255)."
	Min                    int    // "The minimum size of a text being set or received."
	Mode                   string // "The mode off the text entity. Must be either `text` or `password`."
	Name                   string // "The name of the text entity."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	Pattern                string // "A valid regular expression the text being set or received must match with."
	Qos                    int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 bool   // "If the published message should have the retain flag on or not."
	StateFunc              func() string
	UniqueId               string // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception."
	ValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the text state value from the payload received on `state_topic`."
}

func NewTextOptions() *TextOptions {
	return &TextOptions{}
}
func (o *TextOptions) GetStates() *TextState {
	return &o.States
}
func (o *TextOptions) SetAvailabilityMode(mode string) *TextOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *TextOptions) SetAvailabilityTemplate(template string) *TextOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *TextOptions) SetAvailabilityFunc(f func() string) *TextOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *TextOptions) SetCommandTemplate(template string) *TextOptions {
	o.CommandTemplate = template
	return o
}
func (o *TextOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *TextOptions {
	o.CommandFunc = f
	return o
}
func (o *TextOptions) SetEnabledByDefault(d bool) *TextOptions {
	o.EnabledByDefault = d
	return o
}
func (o *TextOptions) SetEncoding(encoding string) *TextOptions {
	o.Encoding = encoding
	return o
}
func (o *TextOptions) SetEntityCategory(category string) *TextOptions {
	o.EntityCategory = category
	return o
}
func (o *TextOptions) SetJsonAttributesTemplate(template string) *TextOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *TextOptions) SetJsonAttributesFunc(f func() string) *TextOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *TextOptions) SetMax(max int) *TextOptions {
	o.Max = max
	return o
}
func (o *TextOptions) SetMin(min int) *TextOptions {
	o.Min = min
	return o
}
func (o *TextOptions) SetMode(mode string) *TextOptions {
	o.Mode = mode
	return o
}
func (o *TextOptions) SetName(name string) *TextOptions {
	o.Name = name
	return o
}
func (o *TextOptions) SetObjectId(id string) *TextOptions {
	o.ObjectId = id
	return o
}
func (o *TextOptions) SetPattern(pattern string) *TextOptions {
	o.Pattern = pattern
	return o
}
func (o *TextOptions) SetQos(qos int) *TextOptions {
	o.Qos = qos
	return o
}
func (o *TextOptions) SetRetain(retain bool) *TextOptions {
	o.Retain = retain
	return o
}
func (o *TextOptions) SetStateFunc(f func() string) *TextOptions {
	o.StateFunc = f
	return o
}
func (o *TextOptions) SetUniqueId(id string) *TextOptions {
	o.UniqueId = id
	return o
}
func (o *TextOptions) SetValueTemplate(template string) *TextOptions {
	o.ValueTemplate = template
	return o
}
