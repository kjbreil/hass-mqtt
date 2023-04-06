package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SensorOptions struct {
	states                    SensorState // External state update location
	availabilityMode          string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate      string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc          func() string
	deviceClass               string // "The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend."
	enabledByDefault          bool   // "Flag which defines if the entity should be enabled when first added."
	encoding                  string // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory            string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	expireAfter               int    // "If set, it defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`. Default the sensors state never expires."
	forceUpdate               bool   // "Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history."
	icon                      string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	jsonAttributesFunc        func() string
	lastResetValueTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the last_reset. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes."
	name                      string // "The name of the MQTT sensor."
	objectId                  string // "Used instead of `name` for automatic generation of `entity_id`"
	payloadAvailable          string // "The payload that represents the available state."
	payloadNotAvailable       string // "The payload that represents the unavailable state."
	qos                       int    // "The maximum QoS level of the state topic."
	stateClass                string // "The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor."
	stateFunc                 func() string
	suggestedDisplayPrecision int    // "The number of decimals which should be used in the sensor's state after rounding."
	uniqueId                  string // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	unitOfMeasurement         string // "Defines the units of measurement of the sensor, if any."
	valueTemplate             string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value. If the template throws an error, the current state will be used instead."
}

func NewSensorOptions() *SensorOptions {
	return &SensorOptions{}
}
func (o *SensorOptions) States() *SensorState {
	return &o.states
}
func (o *SensorOptions) AvailabilityMode(mode string) *SensorOptions {
	o.availabilityMode = mode
	return o
}
func (o *SensorOptions) AvailabilityTemplate(template string) *SensorOptions {
	o.availabilityTemplate = template
	return o
}
func (o *SensorOptions) AvailabilityFunc(f func() string) *SensorOptions {
	o.availabilityFunc = f
	return o
}
func (o *SensorOptions) DeviceClass(class string) *SensorOptions {
	o.deviceClass = class
	return o
}
func (o *SensorOptions) EnabledByDefault(d bool) *SensorOptions {
	o.enabledByDefault = d
	return o
}
func (o *SensorOptions) Encoding(encoding string) *SensorOptions {
	o.encoding = encoding
	return o
}
func (o *SensorOptions) EntityCategory(category string) *SensorOptions {
	o.entityCategory = category
	return o
}
func (o *SensorOptions) ExpireAfter(after int) *SensorOptions {
	o.expireAfter = after
	return o
}
func (o *SensorOptions) ForceUpdate(update bool) *SensorOptions {
	o.forceUpdate = update
	return o
}
func (o *SensorOptions) Icon(icon string) *SensorOptions {
	o.icon = icon
	return o
}
func (o *SensorOptions) JsonAttributesTemplate(template string) *SensorOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *SensorOptions) JsonAttributesFunc(f func() string) *SensorOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *SensorOptions) LastResetValueTemplate(template string) *SensorOptions {
	o.lastResetValueTemplate = template
	return o
}
func (o *SensorOptions) Name(name string) *SensorOptions {
	o.name = name
	return o
}
func (o *SensorOptions) ObjectId(id string) *SensorOptions {
	o.objectId = id
	return o
}
func (o *SensorOptions) PayloadAvailable(available string) *SensorOptions {
	o.payloadAvailable = available
	return o
}
func (o *SensorOptions) PayloadNotAvailable(available string) *SensorOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *SensorOptions) Qos(qos int) *SensorOptions {
	o.qos = qos
	return o
}
func (o *SensorOptions) StateClass(class string) *SensorOptions {
	o.stateClass = class
	return o
}
func (o *SensorOptions) StateFunc(f func() string) *SensorOptions {
	o.stateFunc = f
	return o
}
func (o *SensorOptions) SuggestedDisplayPrecision(precision int) *SensorOptions {
	o.suggestedDisplayPrecision = precision
	return o
}
func (o *SensorOptions) UniqueId(id string) *SensorOptions {
	o.uniqueId = id
	return o
}
func (o *SensorOptions) UnitOfMeasurement(measurement string) *SensorOptions {
	o.unitOfMeasurement = measurement
	return o
}
func (o *SensorOptions) ValueTemplate(template string) *SensorOptions {
	o.valueTemplate = template
	return o
}
