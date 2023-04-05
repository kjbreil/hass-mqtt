package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type DeviceTrackerOptions struct {
	States                 DeviceTrackerState // External state update location
	AvailabilityMode       string             // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string             // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc     func() string
	Name                   string // "The name of the MQTT device_tracker."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       string // "The payload that represents the available state."
	PayloadHome            string // "The payload value that represents the 'home' state for the device."
	PayloadNotAvailable    string // "The payload that represents the unavailable state."
	PayloadNotHome         string // "The payload value that represents the 'not_home' state for the device."
	PayloadReset           string // "The payload value that will have the device's location automatically derived from Home Assistant's zones."
	Qos                    int    // "The maximum QoS level of the state topic."
	SourceType             string // "Attribute of a device tracker that affects state when being used to track a [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`, or `bluetooth_le`."
	StateFunc              func() string
	UniqueId               string // "An ID that uniquely identifies this device_tracker. If two device_trackers have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that returns a device tracker state."
}

func NewDeviceTrackerOptions() *DeviceTrackerOptions {
	return &DeviceTrackerOptions{}
}
func (o *DeviceTrackerOptions) GetStates() *DeviceTrackerState {
	return &o.States
}
func (o *DeviceTrackerOptions) SetAvailabilityMode(mode string) *DeviceTrackerOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *DeviceTrackerOptions) SetAvailabilityTemplate(template string) *DeviceTrackerOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *DeviceTrackerOptions) SetAvailabilityFunc(f func() string) *DeviceTrackerOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *DeviceTrackerOptions) SetIcon(icon string) *DeviceTrackerOptions {
	o.Icon = icon
	return o
}
func (o *DeviceTrackerOptions) SetJsonAttributesTemplate(template string) *DeviceTrackerOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *DeviceTrackerOptions) SetJsonAttributesFunc(f func() string) *DeviceTrackerOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *DeviceTrackerOptions) SetName(name string) *DeviceTrackerOptions {
	o.Name = name
	return o
}
func (o *DeviceTrackerOptions) SetObjectId(id string) *DeviceTrackerOptions {
	o.ObjectId = id
	return o
}
func (o *DeviceTrackerOptions) SetPayloadAvailable(available string) *DeviceTrackerOptions {
	o.PayloadAvailable = available
	return o
}
func (o *DeviceTrackerOptions) SetPayloadHome(home string) *DeviceTrackerOptions {
	o.PayloadHome = home
	return o
}
func (o *DeviceTrackerOptions) SetPayloadNotAvailable(available string) *DeviceTrackerOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *DeviceTrackerOptions) SetPayloadNotHome(home string) *DeviceTrackerOptions {
	o.PayloadNotHome = home
	return o
}
func (o *DeviceTrackerOptions) SetPayloadReset(reset string) *DeviceTrackerOptions {
	o.PayloadReset = reset
	return o
}
func (o *DeviceTrackerOptions) SetQos(qos int) *DeviceTrackerOptions {
	o.Qos = qos
	return o
}
func (o *DeviceTrackerOptions) SetSourceType(t string) *DeviceTrackerOptions {
	o.SourceType = t
	return o
}
func (o *DeviceTrackerOptions) SetStateFunc(f func() string) *DeviceTrackerOptions {
	o.StateFunc = f
	return o
}
func (o *DeviceTrackerOptions) SetUniqueId(id string) *DeviceTrackerOptions {
	o.UniqueId = id
	return o
}
func (o *DeviceTrackerOptions) SetValueTemplate(template string) *DeviceTrackerOptions {
	o.ValueTemplate = template
	return o
}
