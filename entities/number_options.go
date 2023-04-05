package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type NumberOptions struct {
	States                 NumberState // External state update location
	AvailabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityFunc       func() string
	CommandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandFunc            func(mqtt.Message, mqtt.Client)
	DeviceClass            string // "The [type/class](/integrations/number/#device-class) of the number."
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesFunc     func() string
	Max                    float64 // "Maximum value."
	Min                    float64 // "Minimum value."
	Mode                   string  // "Control how the number should be displayed in the UI. Can be set to `box` or `slider` to force a display mode."
	Name                   string  // "The name of the Number."
	ObjectId               string  // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             bool    // "Flag that defines if number works in optimistic mode."
	PayloadReset           string  // "A special payload that resets the state to `None` when received on the `state_topic`."
	Qos                    int     // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 bool    // "If the published message should have the retain flag on or not."
	StateFunc              func() string
	Step                   float64 // "Step value. Smallest value `0.001`."
	UniqueId               string  // "An ID that uniquely identifies this Number. If two Numbers have the same unique ID Home Assistant will raise an exception."
	UnitOfMeasurement      string  // "Defines the unit of measurement of the sensor, if any."
	ValueTemplate          string  // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value."
}

func NewNumberOptions() *NumberOptions {
	return &NumberOptions{}
}
func (o *NumberOptions) GetStates() *NumberState {
	return &o.States
}
func (o *NumberOptions) SetAvailabilityMode(mode string) *NumberOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *NumberOptions) SetAvailabilityFunc(f func() string) *NumberOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *NumberOptions) SetCommandTemplate(template string) *NumberOptions {
	o.CommandTemplate = template
	return o
}
func (o *NumberOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *NumberOptions {
	o.CommandFunc = f
	return o
}
func (o *NumberOptions) SetDeviceClass(class string) *NumberOptions {
	o.DeviceClass = class
	return o
}
func (o *NumberOptions) SetEnabledByDefault(d bool) *NumberOptions {
	o.EnabledByDefault = d
	return o
}
func (o *NumberOptions) SetEncoding(encoding string) *NumberOptions {
	o.Encoding = encoding
	return o
}
func (o *NumberOptions) SetEntityCategory(category string) *NumberOptions {
	o.EntityCategory = category
	return o
}
func (o *NumberOptions) SetIcon(icon string) *NumberOptions {
	o.Icon = icon
	return o
}
func (o *NumberOptions) SetJsonAttributesTemplate(template string) *NumberOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *NumberOptions) SetJsonAttributesFunc(f func() string) *NumberOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *NumberOptions) SetMax(max float64) *NumberOptions {
	o.Max = max
	return o
}
func (o *NumberOptions) SetMin(min float64) *NumberOptions {
	o.Min = min
	return o
}
func (o *NumberOptions) SetMode(mode string) *NumberOptions {
	o.Mode = mode
	return o
}
func (o *NumberOptions) SetName(name string) *NumberOptions {
	o.Name = name
	return o
}
func (o *NumberOptions) SetObjectId(id string) *NumberOptions {
	o.ObjectId = id
	return o
}
func (o *NumberOptions) SetOptimistic(optimistic bool) *NumberOptions {
	o.Optimistic = optimistic
	return o
}
func (o *NumberOptions) SetPayloadReset(reset string) *NumberOptions {
	o.PayloadReset = reset
	return o
}
func (o *NumberOptions) SetQos(qos int) *NumberOptions {
	o.Qos = qos
	return o
}
func (o *NumberOptions) SetRetain(retain bool) *NumberOptions {
	o.Retain = retain
	return o
}
func (o *NumberOptions) SetStateFunc(f func() string) *NumberOptions {
	o.StateFunc = f
	return o
}
func (o *NumberOptions) SetStep(step float64) *NumberOptions {
	o.Step = step
	return o
}
func (o *NumberOptions) SetUniqueId(id string) *NumberOptions {
	o.UniqueId = id
	return o
}
func (o *NumberOptions) SetUnitOfMeasurement(measurement string) *NumberOptions {
	o.UnitOfMeasurement = measurement
	return o
}
func (o *NumberOptions) SetValueTemplate(template string) *NumberOptions {
	o.ValueTemplate = template
	return o
}
