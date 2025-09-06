package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SelectOptions struct {
	states                 SelectState // External state update location
	availabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	commandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to generate the payload to send to `command_topic`."
	commandFunc            func(mqtt.Message, mqtt.Client)
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	entityPicture          string // "Picture URL for the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	jsonAttributesFunc     func() string
	name                   string     // "The name of the Select. Can be set to `null` if only the device name is relevant."
	objectId               string     // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	optimistic             bool       // "Flag that defines if the select works in optimistic mode."
	options                ([]string) // "List of options that can be selected. An empty list or a list with a single item is allowed."
	platform               string     // "Must be `select`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                    int        // "The maximum QoS level to be used when receiving and publishing messages."
	retain                 bool       // "If the published message should have the retain flag on or not."
	stateFunc              func() string
	uniqueId               string // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception. Required when used with device-based discovery."
	valueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the value."
}

func NewSelectOptions() *SelectOptions {
	return &SelectOptions{}
}
func (o *SelectOptions) States() *SelectState {
	return &o.states
}
func (o *SelectOptions) AvailabilityMode(mode string) *SelectOptions {
	o.availabilityMode = mode
	return o
}
func (o *SelectOptions) AvailabilityTemplate(template string) *SelectOptions {
	o.availabilityTemplate = template
	return o
}
func (o *SelectOptions) AvailabilityFunc(f func() string) *SelectOptions {
	o.availabilityFunc = f
	return o
}
func (o *SelectOptions) CommandTemplate(template string) *SelectOptions {
	o.commandTemplate = template
	return o
}
func (o *SelectOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *SelectOptions {
	o.commandFunc = f
	return o
}
func (o *SelectOptions) EnabledByDefault(d bool) *SelectOptions {
	o.enabledByDefault = d
	return o
}
func (o *SelectOptions) Encoding(encoding string) *SelectOptions {
	o.encoding = encoding
	return o
}
func (o *SelectOptions) EntityCategory(category string) *SelectOptions {
	o.entityCategory = category
	return o
}
func (o *SelectOptions) EntityPicture(picture string) *SelectOptions {
	o.entityPicture = picture
	return o
}
func (o *SelectOptions) Icon(icon string) *SelectOptions {
	o.icon = icon
	return o
}
func (o *SelectOptions) JsonAttributesTemplate(template string) *SelectOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *SelectOptions) JsonAttributesFunc(f func() string) *SelectOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *SelectOptions) Name(name string) *SelectOptions {
	o.name = name
	return o
}
func (o *SelectOptions) ObjectId(id string) *SelectOptions {
	o.objectId = id
	return o
}
func (o *SelectOptions) Optimistic(optimistic bool) *SelectOptions {
	o.optimistic = optimistic
	return o
}
func (o *SelectOptions) Options(options []string) *SelectOptions {
	o.options = options
	return o
}
func (o *SelectOptions) Platform(platform string) *SelectOptions {
	o.platform = platform
	return o
}
func (o *SelectOptions) Qos(qos int) *SelectOptions {
	o.qos = qos
	return o
}
func (o *SelectOptions) Retain(retain bool) *SelectOptions {
	o.retain = retain
	return o
}
func (o *SelectOptions) StateFunc(f func() string) *SelectOptions {
	o.stateFunc = f
	return o
}
func (o *SelectOptions) UniqueId(id string) *SelectOptions {
	o.uniqueId = id
	return o
}
func (o *SelectOptions) ValueTemplate(template string) *SelectOptions {
	o.valueTemplate = template
	return o
}
