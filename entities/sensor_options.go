package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SensorOptions struct {
	states                    SensorState // External state update location
	availabilityMode          string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate      string      // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc          func() string
	deviceClass               string // "The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend. The `device_class` can be `null`."
	enabledByDefault          bool   // "Flag which defines if the entity should be enabled when first added."
	encoding                  string // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory            string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity. When set, the entity category must be `diagnostic` for sensors."
	entityPicture             string // "Picture URL for the entity."
	expireAfter               int    // "If set, it defines the number of seconds after the sensor's state expires if it's not updated. After expiry, the sensor's state becomes `unavailable`. Default the sensors state never expires. By default, the sensor's state never expires. Note that when a sensor's value was sent retained to the MQTT broker, the last value sent will be replayed by the MQTT broker when Home Assistant restarts or is reloaded. As this could cause the sensor to become available with an expired state, it is not recommended to retain the sensor's state payload at the MQTT broker. Home Assistant will store and restore the sensor's state for you and calculate the remaining time to retain the sensor's state before it becomes unavailable."
	forceUpdate               bool   // "Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history."
	icon                      string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate    string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	jsonAttributesFunc        func() string
	lastResetValueTemplate    string     // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the last_reset. When `last_reset_value_template` is set, the `state_class` option must be `total`. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes."
	name                      string     // "The name of the MQTT sensor. Can be set to `null` if only the device name is relevant."
	objectId                  string     // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	options                   ([]string) // "List of allowed sensor state value. An empty list is not allowed. The sensor's `device_class` must be set to `enum`. The `options` option cannot be used together with `state_class` or `unit_of_measurement`."
	payloadAvailable          string     // "The payload that represents the available state."
	payloadNotAvailable       string     // "The payload that represents the unavailable state."
	platform                  string     // "Must be `sensor`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                       int        // "The maximum QoS level to be used when receiving and publishing messages."
	stateClass                string     // "The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor."
	stateFunc                 func() string
	suggestedDisplayPrecision int    // "The number of decimals which should be used in the sensor's state after rounding."
	uniqueId                  string // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	unitOfMeasurement         string // "Defines the units of measurement of the sensor, if any. The `unit_of_measurement` can be `null`."
	valueTemplate             string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the value. If the template throws an error, the current state will be used instead."
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
func (o *SensorOptions) EntityPicture(picture string) *SensorOptions {
	o.entityPicture = picture
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
func (o *SensorOptions) Options(options []string) *SensorOptions {
	o.options = options
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
func (o *SensorOptions) Platform(platform string) *SensorOptions {
	o.platform = platform
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
