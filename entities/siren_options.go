package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SirenOptions struct {
	states                 SirenState // External state update location
	availabilityMode       string     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	availableTones         ([]string) // "A list of available tones the siren supports. When configured, this enables the support for setting a `tone` and enables the `tone` state attribute."
	commandOffTemplate     string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate a custom payload to send to `command_topic` when the siren turn off service is called. By default `command_template` will be used as template for service turn off. The variable `value` will be assigned with the configured `payload_off` setting."
	commandTemplate        string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate a custom payload to send to `command_topic`. The variable `value` will be assigned with the configured `payload_on` or `payload_off` setting. The siren turn on service parameters `tone`, `volume_level` or `duration` can be used as variables in the template. When operation in optimistic mode the corresponding state attributes will be set. Turn on parameters will be filtered if a device misses the support."
	commandFunc            func(mqtt.Message, mqtt.Client)
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name to use when displaying this siren."
	objectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	optimistic             bool   // "Flag that defines if siren works in optimistic mode."
	payloadAvailable       string // "The payload that represents the available state."
	payloadNotAvailable    string // "The payload that represents the unavailable state."
	payloadOff             string // "The payload that represents `off` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_off` for details) and sending as `off` command to the `command_topic`."
	payloadOn              string // "The payload that represents `on` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_on`  for details) and sending as `on` command to the `command_topic`."
	qos                    int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	retain                 bool   // "If the published message should have the retain flag on or not."
	stateOff               string // "The payload that represents the `off` state. Used when value that represents `off` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `off`."
	stateOn                string // "The payload that represents the `on` state. Used when value that represents `on` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `on`."
	stateFunc              func() string
	stateValueTemplate     string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's state from the `state_topic`. To determine the siren's state result of this template will be compared to `state_on` and `state_off`. Alternatively `value_template` can be used to render to a valid JSON payload."
	supportDuration        bool   // "Set to `true` if the MQTT siren supports the `duration` service turn on parameter and enables the `duration` state attribute."
	supportVolumeSet       bool   // "Set to `true` if the MQTT siren supports the `volume_set` service turn on parameter and enables the `volume_level` state attribute."
	uniqueId               string // "An ID that uniquely identifies this siren device. If two sirens have the same unique ID, Home Assistant will raise an exception."
}

func NewSirenOptions() *SirenOptions {
	return &SirenOptions{}
}
func (o *SirenOptions) States() *SirenState {
	return &o.states
}
func (o *SirenOptions) AvailabilityMode(mode string) *SirenOptions {
	o.availabilityMode = mode
	return o
}
func (o *SirenOptions) AvailabilityTemplate(template string) *SirenOptions {
	o.availabilityTemplate = template
	return o
}
func (o *SirenOptions) AvailabilityFunc(f func() string) *SirenOptions {
	o.availabilityFunc = f
	return o
}
func (o *SirenOptions) AvailableTones(tones []string) *SirenOptions {
	o.availableTones = tones
	return o
}
func (o *SirenOptions) CommandOffTemplate(template string) *SirenOptions {
	o.commandOffTemplate = template
	return o
}
func (o *SirenOptions) CommandTemplate(template string) *SirenOptions {
	o.commandTemplate = template
	return o
}
func (o *SirenOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *SirenOptions {
	o.commandFunc = f
	return o
}
func (o *SirenOptions) EnabledByDefault(d bool) *SirenOptions {
	o.enabledByDefault = d
	return o
}
func (o *SirenOptions) Encoding(encoding string) *SirenOptions {
	o.encoding = encoding
	return o
}
func (o *SirenOptions) EntityCategory(category string) *SirenOptions {
	o.entityCategory = category
	return o
}
func (o *SirenOptions) Icon(icon string) *SirenOptions {
	o.icon = icon
	return o
}
func (o *SirenOptions) JsonAttributesTemplate(template string) *SirenOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *SirenOptions) JsonAttributesFunc(f func() string) *SirenOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *SirenOptions) Name(name string) *SirenOptions {
	o.name = name
	return o
}
func (o *SirenOptions) ObjectId(id string) *SirenOptions {
	o.objectId = id
	return o
}
func (o *SirenOptions) Optimistic(optimistic bool) *SirenOptions {
	o.optimistic = optimistic
	return o
}
func (o *SirenOptions) PayloadAvailable(available string) *SirenOptions {
	o.payloadAvailable = available
	return o
}
func (o *SirenOptions) PayloadNotAvailable(available string) *SirenOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *SirenOptions) PayloadOff(off string) *SirenOptions {
	o.payloadOff = off
	return o
}
func (o *SirenOptions) PayloadOn(on string) *SirenOptions {
	o.payloadOn = on
	return o
}
func (o *SirenOptions) Qos(qos int) *SirenOptions {
	o.qos = qos
	return o
}
func (o *SirenOptions) Retain(retain bool) *SirenOptions {
	o.retain = retain
	return o
}
func (o *SirenOptions) StateOff(off string) *SirenOptions {
	o.stateOff = off
	return o
}
func (o *SirenOptions) StateOn(on string) *SirenOptions {
	o.stateOn = on
	return o
}
func (o *SirenOptions) StateFunc(f func() string) *SirenOptions {
	o.stateFunc = f
	return o
}
func (o *SirenOptions) StateValueTemplate(template string) *SirenOptions {
	o.stateValueTemplate = template
	return o
}
func (o *SirenOptions) SupportDuration(duration bool) *SirenOptions {
	o.supportDuration = duration
	return o
}
func (o *SirenOptions) SupportVolumeSet(set bool) *SirenOptions {
	o.supportVolumeSet = set
	return o
}
func (o *SirenOptions) UniqueId(id string) *SirenOptions {
	o.uniqueId = id
	return o
}
