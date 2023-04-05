package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SwitchOptions struct {
	States                 SwitchState // External state update location
	AvailabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	CommandFunc            func(mqtt.Message, mqtt.Client)
	DeviceClass            string // "The [type/class](/integrations/switch/#device-class) of the switch to set the icon in the frontend."
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc     func() string
	Name                   string // "The name to use when displaying this switch."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             bool   // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable       string // "The payload that represents the available state."
	PayloadNotAvailable    string // "The payload that represents the unavailable state."
	PayloadOff             string // "The payload that represents `off` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_off` for details) and sending as `off` command to the `command_topic`."
	PayloadOn              string // "The payload that represents `on` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_on`  for details) and sending as `on` command to the `command_topic`."
	Qos                    int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 bool   // "If the published message should have the retain flag on or not."
	StateOff               string // "The payload that represents the `off` state. Used when value that represents `off` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `off`."
	StateOn                string // "The payload that represents the `on` state. Used when value that represents `on` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `on`."
	StateFunc              func() string
	UniqueId               string // "An ID that uniquely identifies this switch device. If two switches have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's state from the `state_topic`. To determine the switches's state result of this template will be compared to `state_on` and `state_off`."
}

func NewSwitchOptions() *SwitchOptions {
	return &SwitchOptions{}
}
func (o *SwitchOptions) GetStates() *SwitchState {
	return &o.States
}
func (o *SwitchOptions) SetAvailabilityMode(mode string) *SwitchOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *SwitchOptions) SetAvailabilityTemplate(template string) *SwitchOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *SwitchOptions) SetAvailabilityFunc(f func() string) *SwitchOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *SwitchOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *SwitchOptions {
	o.CommandFunc = f
	return o
}
func (o *SwitchOptions) SetDeviceClass(class string) *SwitchOptions {
	o.DeviceClass = class
	return o
}
func (o *SwitchOptions) SetEnabledByDefault(d bool) *SwitchOptions {
	o.EnabledByDefault = d
	return o
}
func (o *SwitchOptions) SetEncoding(encoding string) *SwitchOptions {
	o.Encoding = encoding
	return o
}
func (o *SwitchOptions) SetEntityCategory(category string) *SwitchOptions {
	o.EntityCategory = category
	return o
}
func (o *SwitchOptions) SetIcon(icon string) *SwitchOptions {
	o.Icon = icon
	return o
}
func (o *SwitchOptions) SetJsonAttributesTemplate(template string) *SwitchOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *SwitchOptions) SetJsonAttributesFunc(f func() string) *SwitchOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *SwitchOptions) SetName(name string) *SwitchOptions {
	o.Name = name
	return o
}
func (o *SwitchOptions) SetObjectId(id string) *SwitchOptions {
	o.ObjectId = id
	return o
}
func (o *SwitchOptions) SetOptimistic(optimistic bool) *SwitchOptions {
	o.Optimistic = optimistic
	return o
}
func (o *SwitchOptions) SetPayloadAvailable(available string) *SwitchOptions {
	o.PayloadAvailable = available
	return o
}
func (o *SwitchOptions) SetPayloadNotAvailable(available string) *SwitchOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *SwitchOptions) SetPayloadOff(off string) *SwitchOptions {
	o.PayloadOff = off
	return o
}
func (o *SwitchOptions) SetPayloadOn(on string) *SwitchOptions {
	o.PayloadOn = on
	return o
}
func (o *SwitchOptions) SetQos(qos int) *SwitchOptions {
	o.Qos = qos
	return o
}
func (o *SwitchOptions) SetRetain(retain bool) *SwitchOptions {
	o.Retain = retain
	return o
}
func (o *SwitchOptions) SetStateOff(off string) *SwitchOptions {
	o.StateOff = off
	return o
}
func (o *SwitchOptions) SetStateOn(on string) *SwitchOptions {
	o.StateOn = on
	return o
}
func (o *SwitchOptions) SetStateFunc(f func() string) *SwitchOptions {
	o.StateFunc = f
	return o
}
func (o *SwitchOptions) SetUniqueId(id string) *SwitchOptions {
	o.UniqueId = id
	return o
}
func (o *SwitchOptions) SetValueTemplate(template string) *SwitchOptions {
	o.ValueTemplate = template
	return o
}
