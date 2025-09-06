package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type VacuumOptions struct {
	states                 VacuumState // External state update location
	availabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	commandFunc            func(mqtt.Message, mqtt.Client)
	encoding               string     // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	fanSpeedList           ([]string) // "List of possible fan speeds for the vacuum."
	jsonAttributesTemplate string     // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name of the vacuum. Can be set to `null` if only the device name is relevant."
	objectId               string // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	payloadAvailable       string // "The payload that represents the available state."
	payloadCleanSpot       string // "The payload to send to the `command_topic` to begin a spot cleaning cycle."
	payloadLocate          string // "The payload to send to the `command_topic` to locate the vacuum (typically plays a song)."
	payloadNotAvailable    string // "The payload that represents the unavailable state."
	payloadPause           string // "The payload to send to the `command_topic` to pause the vacuum."
	payloadReturnToBase    string // "The payload to send to the `command_topic` to tell the vacuum to return to base."
	payloadStart           string // "The payload to send to the `command_topic` to begin the cleaning cycle."
	payloadStop            string // "The payload to send to the `command_topic` to stop cleaning."
	platform               string // "Must be `vacuum`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                    int    // "The maximum QoS level to be used when receiving and publishing messages."
	retain                 bool   // "If the published message should have the retain flag on or not."
	sendCommandFunc        func(mqtt.Message, mqtt.Client)
	setFanSpeedFunc        func(mqtt.Message, mqtt.Client)
	stateFunc              func() string
	supportedFeatures      ([]string) // "List of features that the vacuum supports (possible values are `start`, `stop`, `pause`, `return_home`, `status`, `locate`, `clean_spot`, `fan_speed`, `send_command`)."
	uniqueId               string     // "An ID that uniquely identifies this vacuum. If two vacuums have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
}

func NewVacuumOptions() *VacuumOptions {
	return &VacuumOptions{}
}
func (o *VacuumOptions) States() *VacuumState {
	return &o.states
}
func (o *VacuumOptions) AvailabilityMode(mode string) *VacuumOptions {
	o.availabilityMode = mode
	return o
}
func (o *VacuumOptions) AvailabilityTemplate(template string) *VacuumOptions {
	o.availabilityTemplate = template
	return o
}
func (o *VacuumOptions) AvailabilityFunc(f func() string) *VacuumOptions {
	o.availabilityFunc = f
	return o
}
func (o *VacuumOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *VacuumOptions {
	o.commandFunc = f
	return o
}
func (o *VacuumOptions) Encoding(encoding string) *VacuumOptions {
	o.encoding = encoding
	return o
}
func (o *VacuumOptions) FanSpeedList(list []string) *VacuumOptions {
	o.fanSpeedList = list
	return o
}
func (o *VacuumOptions) JsonAttributesTemplate(template string) *VacuumOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *VacuumOptions) JsonAttributesFunc(f func() string) *VacuumOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *VacuumOptions) Name(name string) *VacuumOptions {
	o.name = name
	return o
}
func (o *VacuumOptions) ObjectId(id string) *VacuumOptions {
	o.objectId = id
	return o
}
func (o *VacuumOptions) PayloadAvailable(available string) *VacuumOptions {
	o.payloadAvailable = available
	return o
}
func (o *VacuumOptions) PayloadCleanSpot(spot string) *VacuumOptions {
	o.payloadCleanSpot = spot
	return o
}
func (o *VacuumOptions) PayloadLocate(locate string) *VacuumOptions {
	o.payloadLocate = locate
	return o
}
func (o *VacuumOptions) PayloadNotAvailable(available string) *VacuumOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *VacuumOptions) PayloadPause(pause string) *VacuumOptions {
	o.payloadPause = pause
	return o
}
func (o *VacuumOptions) PayloadReturnToBase(base string) *VacuumOptions {
	o.payloadReturnToBase = base
	return o
}
func (o *VacuumOptions) PayloadStart(start string) *VacuumOptions {
	o.payloadStart = start
	return o
}
func (o *VacuumOptions) PayloadStop(stop string) *VacuumOptions {
	o.payloadStop = stop
	return o
}
func (o *VacuumOptions) Platform(platform string) *VacuumOptions {
	o.platform = platform
	return o
}
func (o *VacuumOptions) Qos(qos int) *VacuumOptions {
	o.qos = qos
	return o
}
func (o *VacuumOptions) Retain(retain bool) *VacuumOptions {
	o.retain = retain
	return o
}
func (o *VacuumOptions) SendCommandFunc(f func(mqtt.Message, mqtt.Client)) *VacuumOptions {
	o.sendCommandFunc = f
	return o
}
func (o *VacuumOptions) SetFanSpeedFunc(f func(mqtt.Message, mqtt.Client)) *VacuumOptions {
	o.setFanSpeedFunc = f
	return o
}
func (o *VacuumOptions) StateFunc(f func() string) *VacuumOptions {
	o.stateFunc = f
	return o
}
func (o *VacuumOptions) SupportedFeatures(features []string) *VacuumOptions {
	o.supportedFeatures = features
	return o
}
func (o *VacuumOptions) UniqueId(id string) *VacuumOptions {
	o.uniqueId = id
	return o
}
