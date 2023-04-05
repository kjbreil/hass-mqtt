package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SelectOptions struct {
	States                 SelectState // External state update location
	AvailabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	CommandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandFunc            func(mqtt.Message, mqtt.Client)
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesFunc     func() string
	Name                   string     // "The name of the Select."
	ObjectId               string     // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             bool       // "Flag that defines if the select works in optimistic mode."
	Options                ([]string) // "List of options that can be selected. An empty list or a list with a single item is allowed."
	Qos                    int        // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 bool       // "If the published message should have the retain flag on or not."
	StateFunc              func() string
	UniqueId               string // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception."
	ValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value."
}

func NewSelectOptions() *SelectOptions {
	return &SelectOptions{}
}
func (o *SelectOptions) GetStates() *SelectState {
	return &o.States
}
func (o *SelectOptions) SetAvailabilityMode(mode string) *SelectOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *SelectOptions) SetAvailabilityTemplate(template string) *SelectOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *SelectOptions) SetAvailabilityFunc(f func() string) *SelectOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *SelectOptions) SetCommandTemplate(template string) *SelectOptions {
	o.CommandTemplate = template
	return o
}
func (o *SelectOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *SelectOptions {
	o.CommandFunc = f
	return o
}
func (o *SelectOptions) SetEnabledByDefault(d bool) *SelectOptions {
	o.EnabledByDefault = d
	return o
}
func (o *SelectOptions) SetEncoding(encoding string) *SelectOptions {
	o.Encoding = encoding
	return o
}
func (o *SelectOptions) SetEntityCategory(category string) *SelectOptions {
	o.EntityCategory = category
	return o
}
func (o *SelectOptions) SetIcon(icon string) *SelectOptions {
	o.Icon = icon
	return o
}
func (o *SelectOptions) SetJsonAttributesTemplate(template string) *SelectOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *SelectOptions) SetJsonAttributesFunc(f func() string) *SelectOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *SelectOptions) SetName(name string) *SelectOptions {
	o.Name = name
	return o
}
func (o *SelectOptions) SetObjectId(id string) *SelectOptions {
	o.ObjectId = id
	return o
}
func (o *SelectOptions) SetOptimistic(optimistic bool) *SelectOptions {
	o.Optimistic = optimistic
	return o
}
func (o *SelectOptions) SetOptions(options []string) *SelectOptions {
	o.Options = options
	return o
}
func (o *SelectOptions) SetQos(qos int) *SelectOptions {
	o.Qos = qos
	return o
}
func (o *SelectOptions) SetRetain(retain bool) *SelectOptions {
	o.Retain = retain
	return o
}
func (o *SelectOptions) SetStateFunc(f func() string) *SelectOptions {
	o.StateFunc = f
	return o
}
func (o *SelectOptions) SetUniqueId(id string) *SelectOptions {
	o.UniqueId = id
	return o
}
func (o *SelectOptions) SetValueTemplate(template string) *SelectOptions {
	o.ValueTemplate = template
	return o
}
