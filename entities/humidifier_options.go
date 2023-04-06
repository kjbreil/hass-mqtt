package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type HumidifierOptions struct {
	states                        HumidifierState // External state update location
	availabilityMode              string          // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate          string          // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc              func() string
	commandTemplate               string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	commandFunc                   func(mqtt.Message, mqtt.Client)
	deviceClass                   string // "The device class of the MQTT device. Must be either `humidifier` or `dehumidifier`."
	enabledByDefault              bool   // "Flag which defines if the entity should be enabled when first added."
	encoding                      string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory                string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	icon                          string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc            func() string
	maxHumidity                   int    // "The minimum target humidity percentage that can be set."
	minHumidity                   int    // "The maximum target humidity percentage that can be set."
	modeCommandTemplate           string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `mode_command_topic`."
	modeCommandFunc               func(mqtt.Message, mqtt.Client)
	modeStateTemplate             string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `mode` state."
	modeStateFunc                 func() string
	modes                         ([]string) // "List of available modes this humidifier is capable of running at. Common examples include `normal`, `eco`, `away`, `boost`, `comfort`, `home`, `sleep`, `auto` and `baby`. These examples offer built-in translations but other custom modes are allowed as well.  This attribute ust be configured together with the `mode_command_topic` attribute."
	name                          string     // "The name of the humidifier."
	objectId                      string     // "Used instead of `name` for automatic generation of `entity_id`"
	optimistic                    bool       // "Flag that defines if humidifier works in optimistic mode"
	payloadAvailable              string     // "The payload that represents the available state."
	payloadNotAvailable           string     // "The payload that represents the unavailable state."
	payloadOff                    string     // "The payload that represents the stop state."
	payloadOn                     string     // "The payload that represents the running state."
	payloadResetHumidity          string     // "A special payload that resets the `target_humidity` state attribute to `None` when received at the `target_humidity_state_topic`."
	payloadResetMode              string     // "A special payload that resets the `mode` state attribute to `None` when received at the `mode_state_topic`."
	qos                           int        // "The maximum QoS level of the state topic."
	retain                        bool       // "If the published message should have the retain flag on or not."
	stateFunc                     func() string
	stateValueTemplate            string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state."
	targetHumidityCommandTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `target_humidity_command_topic`."
	targetHumidityCommandFunc     func(mqtt.Message, mqtt.Client)
	targetHumidityStateTemplate   string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `target_humidity` state."
	targetHumidityStateFunc       func() string
	uniqueId                      string // "An ID that uniquely identifies this humidifier. If two humidifiers have the same unique ID, Home Assistant will raise an exception."
}

func NewHumidifierOptions() *HumidifierOptions {
	return &HumidifierOptions{}
}
func (o *HumidifierOptions) States() *HumidifierState {
	return &o.states
}
func (o *HumidifierOptions) AvailabilityMode(mode string) *HumidifierOptions {
	o.availabilityMode = mode
	return o
}
func (o *HumidifierOptions) AvailabilityTemplate(template string) *HumidifierOptions {
	o.availabilityTemplate = template
	return o
}
func (o *HumidifierOptions) AvailabilityFunc(f func() string) *HumidifierOptions {
	o.availabilityFunc = f
	return o
}
func (o *HumidifierOptions) CommandTemplate(template string) *HumidifierOptions {
	o.commandTemplate = template
	return o
}
func (o *HumidifierOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *HumidifierOptions {
	o.commandFunc = f
	return o
}
func (o *HumidifierOptions) DeviceClass(class string) *HumidifierOptions {
	o.deviceClass = class
	return o
}
func (o *HumidifierOptions) EnabledByDefault(d bool) *HumidifierOptions {
	o.enabledByDefault = d
	return o
}
func (o *HumidifierOptions) Encoding(encoding string) *HumidifierOptions {
	o.encoding = encoding
	return o
}
func (o *HumidifierOptions) EntityCategory(category string) *HumidifierOptions {
	o.entityCategory = category
	return o
}
func (o *HumidifierOptions) Icon(icon string) *HumidifierOptions {
	o.icon = icon
	return o
}
func (o *HumidifierOptions) JsonAttributesTemplate(template string) *HumidifierOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *HumidifierOptions) JsonAttributesFunc(f func() string) *HumidifierOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *HumidifierOptions) MaxHumidity(humidity int) *HumidifierOptions {
	o.maxHumidity = humidity
	return o
}
func (o *HumidifierOptions) MinHumidity(humidity int) *HumidifierOptions {
	o.minHumidity = humidity
	return o
}
func (o *HumidifierOptions) ModeCommandTemplate(template string) *HumidifierOptions {
	o.modeCommandTemplate = template
	return o
}
func (o *HumidifierOptions) ModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *HumidifierOptions {
	o.modeCommandFunc = f
	return o
}
func (o *HumidifierOptions) ModeStateTemplate(template string) *HumidifierOptions {
	o.modeStateTemplate = template
	return o
}
func (o *HumidifierOptions) ModeStateFunc(f func() string) *HumidifierOptions {
	o.modeStateFunc = f
	return o
}
func (o *HumidifierOptions) EnableMode() *HumidifierOptions {
	o.modeStateFunc = func() string {
		return o.states.Mode
	}
	o.modeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Mode = string(message.Payload())
	}
	return o
}
func (o *HumidifierOptions) Modes(modes []string) *HumidifierOptions {
	o.modes = modes
	return o
}
func (o *HumidifierOptions) Name(name string) *HumidifierOptions {
	o.name = name
	return o
}
func (o *HumidifierOptions) ObjectId(id string) *HumidifierOptions {
	o.objectId = id
	return o
}
func (o *HumidifierOptions) Optimistic(optimistic bool) *HumidifierOptions {
	o.optimistic = optimistic
	return o
}
func (o *HumidifierOptions) PayloadAvailable(available string) *HumidifierOptions {
	o.payloadAvailable = available
	return o
}
func (o *HumidifierOptions) PayloadNotAvailable(available string) *HumidifierOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *HumidifierOptions) PayloadOff(off string) *HumidifierOptions {
	o.payloadOff = off
	return o
}
func (o *HumidifierOptions) PayloadOn(on string) *HumidifierOptions {
	o.payloadOn = on
	return o
}
func (o *HumidifierOptions) PayloadResetHumidity(humidity string) *HumidifierOptions {
	o.payloadResetHumidity = humidity
	return o
}
func (o *HumidifierOptions) PayloadResetMode(mode string) *HumidifierOptions {
	o.payloadResetMode = mode
	return o
}
func (o *HumidifierOptions) Qos(qos int) *HumidifierOptions {
	o.qos = qos
	return o
}
func (o *HumidifierOptions) Retain(retain bool) *HumidifierOptions {
	o.retain = retain
	return o
}
func (o *HumidifierOptions) StateFunc(f func() string) *HumidifierOptions {
	o.stateFunc = f
	return o
}
func (o *HumidifierOptions) StateValueTemplate(template string) *HumidifierOptions {
	o.stateValueTemplate = template
	return o
}
func (o *HumidifierOptions) TargetHumidityCommandTemplate(template string) *HumidifierOptions {
	o.targetHumidityCommandTemplate = template
	return o
}
func (o *HumidifierOptions) TargetHumidityCommandFunc(f func(mqtt.Message, mqtt.Client)) *HumidifierOptions {
	o.targetHumidityCommandFunc = f
	return o
}
func (o *HumidifierOptions) TargetHumidityStateTemplate(template string) *HumidifierOptions {
	o.targetHumidityStateTemplate = template
	return o
}
func (o *HumidifierOptions) TargetHumidityStateFunc(f func() string) *HumidifierOptions {
	o.targetHumidityStateFunc = f
	return o
}
func (o *HumidifierOptions) EnableTargetHumidity() *HumidifierOptions {
	o.targetHumidityStateFunc = func() string {
		return o.states.TargetHumidity
	}
	o.targetHumidityCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.TargetHumidity = string(message.Payload())
	}
	return o
}
func (o *HumidifierOptions) UniqueId(id string) *HumidifierOptions {
	o.uniqueId = id
	return o
}
