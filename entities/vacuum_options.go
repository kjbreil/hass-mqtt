package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type VacuumOptions struct {
	States                 VacuumState // External state update location
	AvailabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	CommandFunc            func(mqtt.Message, mqtt.Client)
	Encoding               string     // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	FanSpeedList           ([]string) // "List of possible fan speeds for the vacuum."
	JsonAttributesTemplate string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc     func() string
	Name                   string // "The name of the vacuum."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       string // "The payload that represents the available state."
	PayloadCleanSpot       string // "The payload to send to the `command_topic` to begin a spot cleaning cycle."
	PayloadLocate          string // "The payload to send to the `command_topic` to locate the vacuum (typically plays a song)."
	PayloadNotAvailable    string // "The payload that represents the unavailable state."
	PayloadPause           string // "The payload to send to the `command_topic` to pause the vacuum."
	PayloadReturnToBase    string // "The payload to send to the `command_topic` to tell the vacuum to return to base."
	PayloadStart           string // "The payload to send to the `command_topic` to begin the cleaning cycle."
	PayloadStop            string // "The payload to send to the `command_topic` to stop cleaning."
	Qos                    int    // "The maximum QoS level of the state topic."
	Retain                 bool   // "If the published message should have the retain flag on or not."
	Schema                 string // "The schema to use. Must be `state` to select the state schema."
	SendCommandFunc        func(mqtt.Message, mqtt.Client)
	SetFanSpeedFunc        func(mqtt.Message, mqtt.Client)
	StateFunc              func() string
	SupportedFeatures      ([]string) // "List of features that the vacuum supports (possible values are `start`, `stop`, `pause`, `return_home`, `battery`, `status`, `locate`, `clean_spot`, `fan_speed`, `send_command`)."
	UniqueId               string     // "An ID that uniquely identifies this vacuum. If two vacuums have the same unique ID, Home Assistant will raise an exception."
}

func NewVacuumOptions() *VacuumOptions {
	return &VacuumOptions{}
}
func (o *VacuumOptions) GetStates() *VacuumState {
	return &o.States
}
func (o *VacuumOptions) SetAvailabilityMode(mode string) *VacuumOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *VacuumOptions) SetAvailabilityTemplate(template string) *VacuumOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *VacuumOptions) SetAvailabilityFunc(f func() string) *VacuumOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *VacuumOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *VacuumOptions {
	o.CommandFunc = f
	return o
}
func (o *VacuumOptions) SetEncoding(encoding string) *VacuumOptions {
	o.Encoding = encoding
	return o
}
func (o *VacuumOptions) SetFanSpeedList(list []string) *VacuumOptions {
	o.FanSpeedList = list
	return o
}
func (o *VacuumOptions) SetJsonAttributesTemplate(template string) *VacuumOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *VacuumOptions) SetJsonAttributesFunc(f func() string) *VacuumOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *VacuumOptions) SetName(name string) *VacuumOptions {
	o.Name = name
	return o
}
func (o *VacuumOptions) SetObjectId(id string) *VacuumOptions {
	o.ObjectId = id
	return o
}
func (o *VacuumOptions) SetPayloadAvailable(available string) *VacuumOptions {
	o.PayloadAvailable = available
	return o
}
func (o *VacuumOptions) SetPayloadCleanSpot(spot string) *VacuumOptions {
	o.PayloadCleanSpot = spot
	return o
}
func (o *VacuumOptions) SetPayloadLocate(locate string) *VacuumOptions {
	o.PayloadLocate = locate
	return o
}
func (o *VacuumOptions) SetPayloadNotAvailable(available string) *VacuumOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *VacuumOptions) SetPayloadPause(pause string) *VacuumOptions {
	o.PayloadPause = pause
	return o
}
func (o *VacuumOptions) SetPayloadReturnToBase(base string) *VacuumOptions {
	o.PayloadReturnToBase = base
	return o
}
func (o *VacuumOptions) SetPayloadStart(start string) *VacuumOptions {
	o.PayloadStart = start
	return o
}
func (o *VacuumOptions) SetPayloadStop(stop string) *VacuumOptions {
	o.PayloadStop = stop
	return o
}
func (o *VacuumOptions) SetQos(qos int) *VacuumOptions {
	o.Qos = qos
	return o
}
func (o *VacuumOptions) SetRetain(retain bool) *VacuumOptions {
	o.Retain = retain
	return o
}
func (o *VacuumOptions) SetSchema(schema string) *VacuumOptions {
	o.Schema = schema
	return o
}
func (o *VacuumOptions) SetSendCommandFunc(f func(mqtt.Message, mqtt.Client)) *VacuumOptions {
	o.SendCommandFunc = f
	return o
}
func (o *VacuumOptions) SetSetFanSpeedFunc(f func(mqtt.Message, mqtt.Client)) *VacuumOptions {
	o.SetFanSpeedFunc = f
	return o
}
func (o *VacuumOptions) SetStateFunc(f func() string) *VacuumOptions {
	o.StateFunc = f
	return o
}
func (o *VacuumOptions) SetSupportedFeatures(features []string) *VacuumOptions {
	o.SupportedFeatures = features
	return o
}
func (o *VacuumOptions) SetUniqueId(id string) *VacuumOptions {
	o.UniqueId = id
	return o
}
