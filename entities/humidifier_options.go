package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type HumidifierOptions struct {
	States                        HumidifierState // External state update location
	AvailabilityMode              string          // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate          string          // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc              func() string
	CommandTemplate               string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandFunc                   func(mqtt.Message, mqtt.Client)
	DeviceClass                   string // "The device class of the MQTT device. Must be either `humidifier` or `dehumidifier`."
	EnabledByDefault              bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding                      string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                          string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc            func() string
	MaxHumidity                   int    // "The minimum target humidity percentage that can be set."
	MinHumidity                   int    // "The maximum target humidity percentage that can be set."
	ModeCommandTemplate           string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `mode_command_topic`."
	ModeCommandFunc               func(mqtt.Message, mqtt.Client)
	ModeStateTemplate             string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `mode` state."
	ModeStateFunc                 func() string
	Modes                         ([]string) // "List of available modes this humidifier is capable of running at. Common examples include `normal`, `eco`, `away`, `boost`, `comfort`, `home`, `sleep`, `auto` and `baby`. These examples offer built-in translations but other custom modes are allowed as well.  This attribute ust be configured together with the `mode_command_topic` attribute."
	Name                          string     // "The name of the humidifier."
	ObjectId                      string     // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                    bool       // "Flag that defines if humidifier works in optimistic mode"
	PayloadAvailable              string     // "The payload that represents the available state."
	PayloadNotAvailable           string     // "The payload that represents the unavailable state."
	PayloadOff                    string     // "The payload that represents the stop state."
	PayloadOn                     string     // "The payload that represents the running state."
	PayloadResetHumidity          string     // "A special payload that resets the `target_humidity` state attribute to `None` when received at the `target_humidity_state_topic`."
	PayloadResetMode              string     // "A special payload that resets the `mode` state attribute to `None` when received at the `mode_state_topic`."
	Qos                           int        // "The maximum QoS level of the state topic."
	Retain                        bool       // "If the published message should have the retain flag on or not."
	StateFunc                     func() string
	StateValueTemplate            string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state."
	TargetHumidityCommandTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `target_humidity_command_topic`."
	TargetHumidityCommandFunc     func(mqtt.Message, mqtt.Client)
	TargetHumidityStateTemplate   string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `target_humidity` state."
	TargetHumidityStateFunc       func() string
	UniqueId                      string // "An ID that uniquely identifies this humidifier. If two humidifiers have the same unique ID, Home Assistant will raise an exception."
}

func NewHumidifierOptions() *HumidifierOptions {
	return &HumidifierOptions{}
}
func (o *HumidifierOptions) GetStates() *HumidifierState {
	return &o.States
}
func (o *HumidifierOptions) SetAvailabilityMode(mode string) *HumidifierOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *HumidifierOptions) SetAvailabilityTemplate(template string) *HumidifierOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *HumidifierOptions) SetAvailabilityFunc(f func() string) *HumidifierOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *HumidifierOptions) SetCommandTemplate(template string) *HumidifierOptions {
	o.CommandTemplate = template
	return o
}
func (o *HumidifierOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *HumidifierOptions {
	o.CommandFunc = f
	return o
}
func (o *HumidifierOptions) SetDeviceClass(class string) *HumidifierOptions {
	o.DeviceClass = class
	return o
}
func (o *HumidifierOptions) SetEnabledByDefault(d bool) *HumidifierOptions {
	o.EnabledByDefault = d
	return o
}
func (o *HumidifierOptions) SetEncoding(encoding string) *HumidifierOptions {
	o.Encoding = encoding
	return o
}
func (o *HumidifierOptions) SetEntityCategory(category string) *HumidifierOptions {
	o.EntityCategory = category
	return o
}
func (o *HumidifierOptions) SetIcon(icon string) *HumidifierOptions {
	o.Icon = icon
	return o
}
func (o *HumidifierOptions) SetJsonAttributesTemplate(template string) *HumidifierOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *HumidifierOptions) SetJsonAttributesFunc(f func() string) *HumidifierOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *HumidifierOptions) SetMaxHumidity(humidity int) *HumidifierOptions {
	o.MaxHumidity = humidity
	return o
}
func (o *HumidifierOptions) SetMinHumidity(humidity int) *HumidifierOptions {
	o.MinHumidity = humidity
	return o
}
func (o *HumidifierOptions) SetModeCommandTemplate(template string) *HumidifierOptions {
	o.ModeCommandTemplate = template
	return o
}
func (o *HumidifierOptions) SetModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *HumidifierOptions {
	o.ModeCommandFunc = f
	return o
}
func (o *HumidifierOptions) SetModeStateTemplate(template string) *HumidifierOptions {
	o.ModeStateTemplate = template
	return o
}
func (o *HumidifierOptions) SetModeStateFunc(f func() string) *HumidifierOptions {
	o.ModeStateFunc = f
	return o
}
func (o *HumidifierOptions) HasMode() *HumidifierOptions {
	o.ModeStateFunc = func() string {
		return o.States.Mode
	}
	o.ModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Mode = string(message.Payload())
	}
	return o
}
func (o *HumidifierOptions) SetModes(modes []string) *HumidifierOptions {
	o.Modes = modes
	return o
}
func (o *HumidifierOptions) SetName(name string) *HumidifierOptions {
	o.Name = name
	return o
}
func (o *HumidifierOptions) SetObjectId(id string) *HumidifierOptions {
	o.ObjectId = id
	return o
}
func (o *HumidifierOptions) SetOptimistic(optimistic bool) *HumidifierOptions {
	o.Optimistic = optimistic
	return o
}
func (o *HumidifierOptions) SetPayloadAvailable(available string) *HumidifierOptions {
	o.PayloadAvailable = available
	return o
}
func (o *HumidifierOptions) SetPayloadNotAvailable(available string) *HumidifierOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *HumidifierOptions) SetPayloadOff(off string) *HumidifierOptions {
	o.PayloadOff = off
	return o
}
func (o *HumidifierOptions) SetPayloadOn(on string) *HumidifierOptions {
	o.PayloadOn = on
	return o
}
func (o *HumidifierOptions) SetPayloadResetHumidity(humidity string) *HumidifierOptions {
	o.PayloadResetHumidity = humidity
	return o
}
func (o *HumidifierOptions) SetPayloadResetMode(mode string) *HumidifierOptions {
	o.PayloadResetMode = mode
	return o
}
func (o *HumidifierOptions) SetQos(qos int) *HumidifierOptions {
	o.Qos = qos
	return o
}
func (o *HumidifierOptions) SetRetain(retain bool) *HumidifierOptions {
	o.Retain = retain
	return o
}
func (o *HumidifierOptions) SetStateFunc(f func() string) *HumidifierOptions {
	o.StateFunc = f
	return o
}
func (o *HumidifierOptions) SetStateValueTemplate(template string) *HumidifierOptions {
	o.StateValueTemplate = template
	return o
}
func (o *HumidifierOptions) SetTargetHumidityCommandTemplate(template string) *HumidifierOptions {
	o.TargetHumidityCommandTemplate = template
	return o
}
func (o *HumidifierOptions) SetTargetHumidityCommandFunc(f func(mqtt.Message, mqtt.Client)) *HumidifierOptions {
	o.TargetHumidityCommandFunc = f
	return o
}
func (o *HumidifierOptions) SetTargetHumidityStateTemplate(template string) *HumidifierOptions {
	o.TargetHumidityStateTemplate = template
	return o
}
func (o *HumidifierOptions) SetTargetHumidityStateFunc(f func() string) *HumidifierOptions {
	o.TargetHumidityStateFunc = f
	return o
}
func (o *HumidifierOptions) HasTargetHumidity() *HumidifierOptions {
	o.TargetHumidityStateFunc = func() string {
		return o.States.TargetHumidity
	}
	o.TargetHumidityCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.TargetHumidity = string(message.Payload())
	}
	return o
}
func (o *HumidifierOptions) SetUniqueId(id string) *HumidifierOptions {
	o.UniqueId = id
	return o
}
