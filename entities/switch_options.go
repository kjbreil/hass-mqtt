package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SwitchOptions struct {
	states                 SwitchState // External state update location
	availabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	commandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to generate the payload to send to `command_topic`. The switch command template accepts the parameters `value`. The `value` parameter will contain the configured value for either `payload_on` or `payload_off`."
	commandFunc            func(mqtt.Message, mqtt.Client)
	deviceClass            string // "The [type/class](/integrations/switch/#device-class) of the switch to set the icon in the frontend. The `device_class` can be `null`."
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	entityPicture          string // "Picture URL for the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name to use when displaying this switch. Can be set to `null` if only the device name is relevant."
	objectId               string // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	optimistic             bool   // "Flag that defines if switch works in optimistic mode."
	payloadAvailable       string // "The payload that represents the available state."
	payloadNotAvailable    string // "The payload that represents the unavailable state."
	payloadOff             string // "The payload that represents `off` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_off` for details) and sending as `off` command to the `command_topic`."
	payloadOn              string // "The payload that represents `on` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_on`  for details) and sending as `on` command to the `command_topic`."
	platform               string // "Must be `switch`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                    int    // "The maximum QoS level to be used when receiving and publishing messages."
	retain                 bool   // "If the published message should have the retain flag on or not."
	stateOff               string // "The payload that represents the `off` state. Used when value that represents `off` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `off`."
	stateOn                string // "The payload that represents the `on` state. Used when value that represents `on` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `on`."
	stateFunc              func() string
	uniqueId               string // "An ID that uniquely identifies this switch device. If two switches have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	valueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's state from the `state_topic`. To determine the switches's state result of this template will be compared to `state_on` and `state_off`."
}

func NewSwitchOptions() *SwitchOptions {
	return &SwitchOptions{}
}
func (o *SwitchOptions) States() *SwitchState {
	return &o.states
}
func (o *SwitchOptions) AvailabilityMode(mode string) *SwitchOptions {
	o.availabilityMode = mode
	return o
}
func (o *SwitchOptions) AvailabilityTemplate(template string) *SwitchOptions {
	o.availabilityTemplate = template
	return o
}
func (o *SwitchOptions) AvailabilityFunc(f func() string) *SwitchOptions {
	o.availabilityFunc = f
	return o
}
func (o *SwitchOptions) CommandTemplate(template string) *SwitchOptions {
	o.commandTemplate = template
	return o
}
func (o *SwitchOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *SwitchOptions {
	o.commandFunc = f
	return o
}
func (o *SwitchOptions) DeviceClass(class string) *SwitchOptions {
	o.deviceClass = class
	return o
}
func (o *SwitchOptions) EnabledByDefault(d bool) *SwitchOptions {
	o.enabledByDefault = d
	return o
}
func (o *SwitchOptions) Encoding(encoding string) *SwitchOptions {
	o.encoding = encoding
	return o
}
func (o *SwitchOptions) EntityCategory(category string) *SwitchOptions {
	o.entityCategory = category
	return o
}
func (o *SwitchOptions) EntityPicture(picture string) *SwitchOptions {
	o.entityPicture = picture
	return o
}
func (o *SwitchOptions) Icon(icon string) *SwitchOptions {
	o.icon = icon
	return o
}
func (o *SwitchOptions) JsonAttributesTemplate(template string) *SwitchOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *SwitchOptions) JsonAttributesFunc(f func() string) *SwitchOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *SwitchOptions) Name(name string) *SwitchOptions {
	o.name = name
	return o
}
func (o *SwitchOptions) ObjectId(id string) *SwitchOptions {
	o.objectId = id
	return o
}
func (o *SwitchOptions) Optimistic(optimistic bool) *SwitchOptions {
	o.optimistic = optimistic
	return o
}
func (o *SwitchOptions) PayloadAvailable(available string) *SwitchOptions {
	o.payloadAvailable = available
	return o
}
func (o *SwitchOptions) PayloadNotAvailable(available string) *SwitchOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *SwitchOptions) PayloadOff(off string) *SwitchOptions {
	o.payloadOff = off
	return o
}
func (o *SwitchOptions) PayloadOn(on string) *SwitchOptions {
	o.payloadOn = on
	return o
}
func (o *SwitchOptions) Platform(platform string) *SwitchOptions {
	o.platform = platform
	return o
}
func (o *SwitchOptions) Qos(qos int) *SwitchOptions {
	o.qos = qos
	return o
}
func (o *SwitchOptions) Retain(retain bool) *SwitchOptions {
	o.retain = retain
	return o
}
func (o *SwitchOptions) StateOff(off string) *SwitchOptions {
	o.stateOff = off
	return o
}
func (o *SwitchOptions) StateOn(on string) *SwitchOptions {
	o.stateOn = on
	return o
}
func (o *SwitchOptions) StateFunc(f func() string) *SwitchOptions {
	o.stateFunc = f
	return o
}
func (o *SwitchOptions) UniqueId(id string) *SwitchOptions {
	o.uniqueId = id
	return o
}
func (o *SwitchOptions) ValueTemplate(template string) *SwitchOptions {
	o.valueTemplate = template
	return o
}
