package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type BinarySensorOptions struct {
	states                 BinarySensorState // External state update location
	availabilityMode       string            // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string            // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	deviceClass            string // "Sets the [class of the device](/integrations/binary_sensor/#device-class), changing the device state and icon that is displayed on the frontend."
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	expireAfter            int    // "If set, it defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`. Default the sensors state never expires."
	forceUpdate            bool   // "Sends update events (which results in update of [state object](/docs/configuration/state_object/)'s `last_changed`) even if the sensor's state hasn't changed. Useful if you want to have meaningful value graphs in history or want to create an automation that triggers on *every* incoming state message (not only when the sensor's new state is different to the current one)."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name of the binary sensor."
	objectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	offDelay               int    // "For sensors that only send `on` state updates (like PIRs), this variable sets a delay in seconds after which the sensor's state will be updated back to `off`."
	payloadAvailable       string // "The string that represents the `online` state."
	payloadNotAvailable    string // "The string that represents the `offline` state."
	payloadOff             string // "The string that represents the `off` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	payloadOn              string // "The string that represents the `on` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	qos                    int    // "The maximum QoS level to be used when receiving messages."
	stateFunc              func() string
	uniqueId               string // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	valueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that returns a string to be compared to `payload_on`/`payload_off` or an empty string, in which case the MQTT message will be removed. Remove this option when `payload_on` and `payload_off` are sufficient to match your payloads (i.e no pre-processing of original message is required)."
}

func NewBinarySensorOptions() *BinarySensorOptions {
	return &BinarySensorOptions{}
}
func (o *BinarySensorOptions) States() *BinarySensorState {
	return &o.states
}
func (o *BinarySensorOptions) AvailabilityMode(mode string) *BinarySensorOptions {
	o.availabilityMode = mode
	return o
}
func (o *BinarySensorOptions) AvailabilityTemplate(template string) *BinarySensorOptions {
	o.availabilityTemplate = template
	return o
}
func (o *BinarySensorOptions) AvailabilityFunc(f func() string) *BinarySensorOptions {
	o.availabilityFunc = f
	return o
}
func (o *BinarySensorOptions) DeviceClass(class string) *BinarySensorOptions {
	o.deviceClass = class
	return o
}
func (o *BinarySensorOptions) EnabledByDefault(d bool) *BinarySensorOptions {
	o.enabledByDefault = d
	return o
}
func (o *BinarySensorOptions) Encoding(encoding string) *BinarySensorOptions {
	o.encoding = encoding
	return o
}
func (o *BinarySensorOptions) EntityCategory(category string) *BinarySensorOptions {
	o.entityCategory = category
	return o
}
func (o *BinarySensorOptions) ExpireAfter(after int) *BinarySensorOptions {
	o.expireAfter = after
	return o
}
func (o *BinarySensorOptions) ForceUpdate(update bool) *BinarySensorOptions {
	o.forceUpdate = update
	return o
}
func (o *BinarySensorOptions) Icon(icon string) *BinarySensorOptions {
	o.icon = icon
	return o
}
func (o *BinarySensorOptions) JsonAttributesTemplate(template string) *BinarySensorOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *BinarySensorOptions) JsonAttributesFunc(f func() string) *BinarySensorOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *BinarySensorOptions) Name(name string) *BinarySensorOptions {
	o.name = name
	return o
}
func (o *BinarySensorOptions) ObjectId(id string) *BinarySensorOptions {
	o.objectId = id
	return o
}
func (o *BinarySensorOptions) OffDelay(delay int) *BinarySensorOptions {
	o.offDelay = delay
	return o
}
func (o *BinarySensorOptions) PayloadAvailable(available string) *BinarySensorOptions {
	o.payloadAvailable = available
	return o
}
func (o *BinarySensorOptions) PayloadNotAvailable(available string) *BinarySensorOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *BinarySensorOptions) PayloadOff(off string) *BinarySensorOptions {
	o.payloadOff = off
	return o
}
func (o *BinarySensorOptions) PayloadOn(on string) *BinarySensorOptions {
	o.payloadOn = on
	return o
}
func (o *BinarySensorOptions) Qos(qos int) *BinarySensorOptions {
	o.qos = qos
	return o
}
func (o *BinarySensorOptions) StateFunc(f func() string) *BinarySensorOptions {
	o.stateFunc = f
	return o
}
func (o *BinarySensorOptions) UniqueId(id string) *BinarySensorOptions {
	o.uniqueId = id
	return o
}
func (o *BinarySensorOptions) ValueTemplate(template string) *BinarySensorOptions {
	o.valueTemplate = template
	return o
}
