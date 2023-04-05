package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SirenOptions struct {
	States                 SirenState // External state update location
	AvailabilityMode       string     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	AvailableTones         ([]string) // "A list of available tones the siren supports. When configured, this enables the support for setting a `tone` and enables the `tone` state attribute."
	CommandOffTemplate     string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate a custom payload to send to `command_topic` when the siren turn off service is called. By default `command_template` will be used as template for service turn off. The variable `value` will be assigned with the configured `payload_off` setting."
	CommandTemplate        string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate a custom payload to send to `command_topic`. The variable `value` will be assigned with the configured `payload_on` or `payload_off` setting. The siren turn on service parameters `tone`, `volume_level` or `duration` can be used as variables in the template. When operation in optimistic mode the corresponding state attributes will be set. Turn on parameters will be filtered if a device misses the support."
	CommandFunc            func(mqtt.Message, mqtt.Client)
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc     func() string
	Name                   string // "The name to use when displaying this siren."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             bool   // "Flag that defines if siren works in optimistic mode."
	PayloadAvailable       string // "The payload that represents the available state."
	PayloadNotAvailable    string // "The payload that represents the unavailable state."
	PayloadOff             string // "The payload that represents `off` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_off` for details) and sending as `off` command to the `command_topic`."
	PayloadOn              string // "The payload that represents `on` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_on`  for details) and sending as `on` command to the `command_topic`."
	Qos                    int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 bool   // "If the published message should have the retain flag on or not."
	StateOff               string // "The payload that represents the `off` state. Used when value that represents `off` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `off`."
	StateOn                string // "The payload that represents the `on` state. Used when value that represents `on` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `on`."
	StateFunc              func() string
	StateValueTemplate     string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's state from the `state_topic`. To determine the siren's state result of this template will be compared to `state_on` and `state_off`. Alternatively `value_template` can be used to render to a valid JSON payload."
	SupportDuration        bool   // "Set to `true` if the MQTT siren supports the `duration` service turn on parameter and enables the `duration` state attribute."
	SupportVolumeSet       bool   // "Set to `true` if the MQTT siren supports the `volume_set` service turn on parameter and enables the `volume_level` state attribute."
	UniqueId               string // "An ID that uniquely identifies this siren device. If two sirens have the same unique ID, Home Assistant will raise an exception."
}

func NewSirenOptions() *SirenOptions {
	return &SirenOptions{}
}
func (o *SirenOptions) GetStates() *SirenState {
	return &o.States
}
func (o *SirenOptions) SetAvailabilityMode(mode string) *SirenOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *SirenOptions) SetAvailabilityTemplate(template string) *SirenOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *SirenOptions) SetAvailabilityFunc(f func() string) *SirenOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *SirenOptions) SetAvailableTones(tones []string) *SirenOptions {
	o.AvailableTones = tones
	return o
}
func (o *SirenOptions) SetCommandOffTemplate(template string) *SirenOptions {
	o.CommandOffTemplate = template
	return o
}
func (o *SirenOptions) SetCommandTemplate(template string) *SirenOptions {
	o.CommandTemplate = template
	return o
}
func (o *SirenOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *SirenOptions {
	o.CommandFunc = f
	return o
}
func (o *SirenOptions) SetEnabledByDefault(d bool) *SirenOptions {
	o.EnabledByDefault = d
	return o
}
func (o *SirenOptions) SetEncoding(encoding string) *SirenOptions {
	o.Encoding = encoding
	return o
}
func (o *SirenOptions) SetEntityCategory(category string) *SirenOptions {
	o.EntityCategory = category
	return o
}
func (o *SirenOptions) SetIcon(icon string) *SirenOptions {
	o.Icon = icon
	return o
}
func (o *SirenOptions) SetJsonAttributesTemplate(template string) *SirenOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *SirenOptions) SetJsonAttributesFunc(f func() string) *SirenOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *SirenOptions) SetName(name string) *SirenOptions {
	o.Name = name
	return o
}
func (o *SirenOptions) SetObjectId(id string) *SirenOptions {
	o.ObjectId = id
	return o
}
func (o *SirenOptions) SetOptimistic(optimistic bool) *SirenOptions {
	o.Optimistic = optimistic
	return o
}
func (o *SirenOptions) SetPayloadAvailable(available string) *SirenOptions {
	o.PayloadAvailable = available
	return o
}
func (o *SirenOptions) SetPayloadNotAvailable(available string) *SirenOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *SirenOptions) SetPayloadOff(off string) *SirenOptions {
	o.PayloadOff = off
	return o
}
func (o *SirenOptions) SetPayloadOn(on string) *SirenOptions {
	o.PayloadOn = on
	return o
}
func (o *SirenOptions) SetQos(qos int) *SirenOptions {
	o.Qos = qos
	return o
}
func (o *SirenOptions) SetRetain(retain bool) *SirenOptions {
	o.Retain = retain
	return o
}
func (o *SirenOptions) SetStateOff(off string) *SirenOptions {
	o.StateOff = off
	return o
}
func (o *SirenOptions) SetStateOn(on string) *SirenOptions {
	o.StateOn = on
	return o
}
func (o *SirenOptions) SetStateFunc(f func() string) *SirenOptions {
	o.StateFunc = f
	return o
}
func (o *SirenOptions) SetStateValueTemplate(template string) *SirenOptions {
	o.StateValueTemplate = template
	return o
}
func (o *SirenOptions) SetSupportDuration(duration bool) *SirenOptions {
	o.SupportDuration = duration
	return o
}
func (o *SirenOptions) SetSupportVolumeSet(set bool) *SirenOptions {
	o.SupportVolumeSet = set
	return o
}
func (o *SirenOptions) SetUniqueId(id string) *SirenOptions {
	o.UniqueId = id
	return o
}
