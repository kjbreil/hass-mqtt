package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type BinarySensorOptions struct {
	States                 BinarySensorState // External state update location
	AvailabilityMode       string            // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string            // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	DeviceClass            string // "Sets the [class of the device](/integrations/binary_sensor/#device-class), changing the device state and icon that is displayed on the frontend."
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	ExpireAfter            int    // "If set, it defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`. Default the sensors state never expires."
	ForceUpdate            bool   // "Sends update events (which results in update of [state object](/docs/configuration/state_object/)'s `last_changed`) even if the sensor's state hasn't changed. Useful if you want to have meaningful value graphs in history or want to create an automation that triggers on *every* incoming state message (not only when the sensor's new state is different to the current one)."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc     func() string
	Name                   string // "The name of the binary sensor."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	OffDelay               int    // "For sensors that only send `on` state updates (like PIRs), this variable sets a delay in seconds after which the sensor's state will be updated back to `off`."
	PayloadAvailable       string // "The string that represents the `online` state."
	PayloadNotAvailable    string // "The string that represents the `offline` state."
	PayloadOff             string // "The string that represents the `off` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	PayloadOn              string // "The string that represents the `on` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	Qos                    int    // "The maximum QoS level to be used when receiving messages."
	StateFunc              func() string
	UniqueId               string // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that returns a string to be compared to `payload_on`/`payload_off` or an empty string, in which case the MQTT message will be removed. Remove this option when `payload_on` and `payload_off` are sufficient to match your payloads (i.e no pre-processing of original message is required)."
}

func NewBinarySensorOptions() *BinarySensorOptions {
	return &BinarySensorOptions{}
}
func (o *BinarySensorOptions) GetStates() *BinarySensorState {
	return &o.States
}
func (o *BinarySensorOptions) SetAvailabilityMode(mode string) *BinarySensorOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *BinarySensorOptions) SetAvailabilityTemplate(template string) *BinarySensorOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *BinarySensorOptions) SetAvailabilityFunc(f func() string) *BinarySensorOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *BinarySensorOptions) SetDeviceClass(class string) *BinarySensorOptions {
	o.DeviceClass = class
	return o
}
func (o *BinarySensorOptions) SetEnabledByDefault(d bool) *BinarySensorOptions {
	o.EnabledByDefault = d
	return o
}
func (o *BinarySensorOptions) SetEncoding(encoding string) *BinarySensorOptions {
	o.Encoding = encoding
	return o
}
func (o *BinarySensorOptions) SetEntityCategory(category string) *BinarySensorOptions {
	o.EntityCategory = category
	return o
}
func (o *BinarySensorOptions) SetExpireAfter(after int) *BinarySensorOptions {
	o.ExpireAfter = after
	return o
}
func (o *BinarySensorOptions) SetForceUpdate(update bool) *BinarySensorOptions {
	o.ForceUpdate = update
	return o
}
func (o *BinarySensorOptions) SetIcon(icon string) *BinarySensorOptions {
	o.Icon = icon
	return o
}
func (o *BinarySensorOptions) SetJsonAttributesTemplate(template string) *BinarySensorOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *BinarySensorOptions) SetJsonAttributesFunc(f func() string) *BinarySensorOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *BinarySensorOptions) SetName(name string) *BinarySensorOptions {
	o.Name = name
	return o
}
func (o *BinarySensorOptions) SetObjectId(id string) *BinarySensorOptions {
	o.ObjectId = id
	return o
}
func (o *BinarySensorOptions) SetOffDelay(delay int) *BinarySensorOptions {
	o.OffDelay = delay
	return o
}
func (o *BinarySensorOptions) SetPayloadAvailable(available string) *BinarySensorOptions {
	o.PayloadAvailable = available
	return o
}
func (o *BinarySensorOptions) SetPayloadNotAvailable(available string) *BinarySensorOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *BinarySensorOptions) SetPayloadOff(off string) *BinarySensorOptions {
	o.PayloadOff = off
	return o
}
func (o *BinarySensorOptions) SetPayloadOn(on string) *BinarySensorOptions {
	o.PayloadOn = on
	return o
}
func (o *BinarySensorOptions) SetQos(qos int) *BinarySensorOptions {
	o.Qos = qos
	return o
}
func (o *BinarySensorOptions) SetStateFunc(f func() string) *BinarySensorOptions {
	o.StateFunc = f
	return o
}
func (o *BinarySensorOptions) SetUniqueId(id string) *BinarySensorOptions {
	o.UniqueId = id
	return o
}
func (o *BinarySensorOptions) SetValueTemplate(template string) *BinarySensorOptions {
	o.ValueTemplate = template
	return o
}
