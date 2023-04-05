package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SensorOptions struct {
	States                    SensorState // External state update location
	AvailabilityMode          string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate      string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc          func() string
	DeviceClass               string // "The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend."
	EnabledByDefault          bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding                  string // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory            string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	ExpireAfter               int    // "If set, it defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`. Default the sensors state never expires."
	ForceUpdate               bool   // "Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history."
	Icon                      string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesFunc        func() string
	LastResetValueTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the last_reset. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes."
	Name                      string // "The name of the MQTT sensor."
	ObjectId                  string // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable          string // "The payload that represents the available state."
	PayloadNotAvailable       string // "The payload that represents the unavailable state."
	Qos                       int    // "The maximum QoS level of the state topic."
	StateClass                string // "The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor."
	StateFunc                 func() string
	SuggestedDisplayPrecision int    // "The number of decimals which should be used in the sensor's state after rounding."
	UniqueId                  string // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	UnitOfMeasurement         string // "Defines the units of measurement of the sensor, if any."
	ValueTemplate             string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value. If the template throws an error, the current state will be used instead."
}

func NewSensorOptions() *SensorOptions {
	return &SensorOptions{}
}
func (o *SensorOptions) GetStates() *SensorState {
	return &o.States
}
func (o *SensorOptions) SetAvailabilityMode(mode string) *SensorOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *SensorOptions) SetAvailabilityTemplate(template string) *SensorOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *SensorOptions) SetAvailabilityFunc(f func() string) *SensorOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *SensorOptions) SetDeviceClass(class string) *SensorOptions {
	o.DeviceClass = class
	return o
}
func (o *SensorOptions) SetEnabledByDefault(d bool) *SensorOptions {
	o.EnabledByDefault = d
	return o
}
func (o *SensorOptions) SetEncoding(encoding string) *SensorOptions {
	o.Encoding = encoding
	return o
}
func (o *SensorOptions) SetEntityCategory(category string) *SensorOptions {
	o.EntityCategory = category
	return o
}
func (o *SensorOptions) SetExpireAfter(after int) *SensorOptions {
	o.ExpireAfter = after
	return o
}
func (o *SensorOptions) SetForceUpdate(update bool) *SensorOptions {
	o.ForceUpdate = update
	return o
}
func (o *SensorOptions) SetIcon(icon string) *SensorOptions {
	o.Icon = icon
	return o
}
func (o *SensorOptions) SetJsonAttributesTemplate(template string) *SensorOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *SensorOptions) SetJsonAttributesFunc(f func() string) *SensorOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *SensorOptions) SetLastResetValueTemplate(template string) *SensorOptions {
	o.LastResetValueTemplate = template
	return o
}
func (o *SensorOptions) SetName(name string) *SensorOptions {
	o.Name = name
	return o
}
func (o *SensorOptions) SetObjectId(id string) *SensorOptions {
	o.ObjectId = id
	return o
}
func (o *SensorOptions) SetPayloadAvailable(available string) *SensorOptions {
	o.PayloadAvailable = available
	return o
}
func (o *SensorOptions) SetPayloadNotAvailable(available string) *SensorOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *SensorOptions) SetQos(qos int) *SensorOptions {
	o.Qos = qos
	return o
}
func (o *SensorOptions) SetStateClass(class string) *SensorOptions {
	o.StateClass = class
	return o
}
func (o *SensorOptions) SetStateFunc(f func() string) *SensorOptions {
	o.StateFunc = f
	return o
}
func (o *SensorOptions) SetSuggestedDisplayPrecision(precision int) *SensorOptions {
	o.SuggestedDisplayPrecision = precision
	return o
}
func (o *SensorOptions) SetUniqueId(id string) *SensorOptions {
	o.UniqueId = id
	return o
}
func (o *SensorOptions) SetUnitOfMeasurement(measurement string) *SensorOptions {
	o.UnitOfMeasurement = measurement
	return o
}
func (o *SensorOptions) SetValueTemplate(template string) *SensorOptions {
	o.ValueTemplate = template
	return o
}
