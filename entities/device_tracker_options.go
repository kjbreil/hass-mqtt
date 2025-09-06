package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type DeviceTrackerOptions struct {
	states                 DeviceTrackerState // External state update location
	availabilityMode       string             // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string             // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name of the MQTT device_tracker."
	objectId               string // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	payloadAvailable       string // "The payload that represents the available state."
	payloadHome            string // "The payload value that represents the 'home' state for the device."
	payloadNotAvailable    string // "The payload that represents the unavailable state."
	payloadNotHome         string // "The payload value that represents the 'not_home' state for the device."
	payloadReset           string // "The payload value that will have the device's location automatically derived from Home Assistant's zones."
	platform               string // "Must be `device_tracker`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                    int    // "The maximum QoS level to be used when receiving and publishing messages."
	sourceType             string // "Attribute of a device tracker that affects state when being used to track a [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`, or `bluetooth_le`."
	stateFunc              func() string
	uniqueId               string // "An ID that uniquely identifies this device_tracker. If two device_trackers have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	valueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) that returns a device tracker state."
}

func NewDeviceTrackerOptions() *DeviceTrackerOptions {
	return &DeviceTrackerOptions{}
}
func (o *DeviceTrackerOptions) States() *DeviceTrackerState {
	return &o.states
}
func (o *DeviceTrackerOptions) AvailabilityMode(mode string) *DeviceTrackerOptions {
	o.availabilityMode = mode
	return o
}
func (o *DeviceTrackerOptions) AvailabilityTemplate(template string) *DeviceTrackerOptions {
	o.availabilityTemplate = template
	return o
}
func (o *DeviceTrackerOptions) AvailabilityFunc(f func() string) *DeviceTrackerOptions {
	o.availabilityFunc = f
	return o
}
func (o *DeviceTrackerOptions) Icon(icon string) *DeviceTrackerOptions {
	o.icon = icon
	return o
}
func (o *DeviceTrackerOptions) JsonAttributesTemplate(template string) *DeviceTrackerOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *DeviceTrackerOptions) JsonAttributesFunc(f func() string) *DeviceTrackerOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *DeviceTrackerOptions) Name(name string) *DeviceTrackerOptions {
	o.name = name
	return o
}
func (o *DeviceTrackerOptions) ObjectId(id string) *DeviceTrackerOptions {
	o.objectId = id
	return o
}
func (o *DeviceTrackerOptions) PayloadAvailable(available string) *DeviceTrackerOptions {
	o.payloadAvailable = available
	return o
}
func (o *DeviceTrackerOptions) PayloadHome(home string) *DeviceTrackerOptions {
	o.payloadHome = home
	return o
}
func (o *DeviceTrackerOptions) PayloadNotAvailable(available string) *DeviceTrackerOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *DeviceTrackerOptions) PayloadNotHome(home string) *DeviceTrackerOptions {
	o.payloadNotHome = home
	return o
}
func (o *DeviceTrackerOptions) PayloadReset(reset string) *DeviceTrackerOptions {
	o.payloadReset = reset
	return o
}
func (o *DeviceTrackerOptions) Platform(platform string) *DeviceTrackerOptions {
	o.platform = platform
	return o
}
func (o *DeviceTrackerOptions) Qos(qos int) *DeviceTrackerOptions {
	o.qos = qos
	return o
}
func (o *DeviceTrackerOptions) SourceType(t string) *DeviceTrackerOptions {
	o.sourceType = t
	return o
}
func (o *DeviceTrackerOptions) StateFunc(f func() string) *DeviceTrackerOptions {
	o.stateFunc = f
	return o
}
func (o *DeviceTrackerOptions) UniqueId(id string) *DeviceTrackerOptions {
	o.uniqueId = id
	return o
}
func (o *DeviceTrackerOptions) ValueTemplate(template string) *DeviceTrackerOptions {
	o.valueTemplate = template
	return o
}
