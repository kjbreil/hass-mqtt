package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type NumberOptions struct {
	states                 NumberState // External state update location
	availabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityFunc       func() string
	commandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	commandFunc            func(mqtt.Message, mqtt.Client)
	deviceClass            string // "The [type/class](/integrations/number/#device-class) of the number."
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	jsonAttributesFunc     func() string
	max                    float64 // "Maximum value."
	min                    float64 // "Minimum value."
	mode                   string  // "Control how the number should be displayed in the UI. Can be set to `box` or `slider` to force a display mode."
	name                   string  // "The name of the Number."
	objectId               string  // "Used instead of `name` for automatic generation of `entity_id`"
	optimistic             bool    // "Flag that defines if number works in optimistic mode."
	payloadReset           string  // "A special payload that resets the state to `None` when received on the `state_topic`."
	qos                    int     // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	retain                 bool    // "If the published message should have the retain flag on or not."
	stateFunc              func() string
	step                   float64 // "Step value. Smallest value `0.001`."
	uniqueId               string  // "An ID that uniquely identifies this Number. If two Numbers have the same unique ID Home Assistant will raise an exception."
	unitOfMeasurement      string  // "Defines the unit of measurement of the sensor, if any."
	valueTemplate          string  // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value."
}

func NewNumberOptions() *NumberOptions {
	return &NumberOptions{}
}
func (o *NumberOptions) States() *NumberState {
	return &o.states
}
func (o *NumberOptions) AvailabilityMode(mode string) *NumberOptions {
	o.availabilityMode = mode
	return o
}
func (o *NumberOptions) AvailabilityFunc(f func() string) *NumberOptions {
	o.availabilityFunc = f
	return o
}
func (o *NumberOptions) CommandTemplate(template string) *NumberOptions {
	o.commandTemplate = template
	return o
}
func (o *NumberOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *NumberOptions {
	o.commandFunc = f
	return o
}
func (o *NumberOptions) DeviceClass(class string) *NumberOptions {
	o.deviceClass = class
	return o
}
func (o *NumberOptions) EnabledByDefault(d bool) *NumberOptions {
	o.enabledByDefault = d
	return o
}
func (o *NumberOptions) Encoding(encoding string) *NumberOptions {
	o.encoding = encoding
	return o
}
func (o *NumberOptions) EntityCategory(category string) *NumberOptions {
	o.entityCategory = category
	return o
}
func (o *NumberOptions) Icon(icon string) *NumberOptions {
	o.icon = icon
	return o
}
func (o *NumberOptions) JsonAttributesTemplate(template string) *NumberOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *NumberOptions) JsonAttributesFunc(f func() string) *NumberOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *NumberOptions) Max(max float64) *NumberOptions {
	o.max = max
	return o
}
func (o *NumberOptions) Min(min float64) *NumberOptions {
	o.min = min
	return o
}
func (o *NumberOptions) Mode(mode string) *NumberOptions {
	o.mode = mode
	return o
}
func (o *NumberOptions) Name(name string) *NumberOptions {
	o.name = name
	return o
}
func (o *NumberOptions) ObjectId(id string) *NumberOptions {
	o.objectId = id
	return o
}
func (o *NumberOptions) Optimistic(optimistic bool) *NumberOptions {
	o.optimistic = optimistic
	return o
}
func (o *NumberOptions) PayloadReset(reset string) *NumberOptions {
	o.payloadReset = reset
	return o
}
func (o *NumberOptions) Qos(qos int) *NumberOptions {
	o.qos = qos
	return o
}
func (o *NumberOptions) Retain(retain bool) *NumberOptions {
	o.retain = retain
	return o
}
func (o *NumberOptions) StateFunc(f func() string) *NumberOptions {
	o.stateFunc = f
	return o
}
func (o *NumberOptions) Step(step float64) *NumberOptions {
	o.step = step
	return o
}
func (o *NumberOptions) UniqueId(id string) *NumberOptions {
	o.uniqueId = id
	return o
}
func (o *NumberOptions) UnitOfMeasurement(measurement string) *NumberOptions {
	o.unitOfMeasurement = measurement
	return o
}
func (o *NumberOptions) ValueTemplate(template string) *NumberOptions {
	o.valueTemplate = template
	return o
}
