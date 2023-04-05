package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type AlarmControlPanelOptions struct {
	States                 AlarmControlPanelState // External state update location
	AvailabilityMode       string                 // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string                 // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	Code                   string // "If defined, specifies a code to enable or disable the alarm in the frontend. Note that the code is validated locally and blocks sending MQTT messages to the remote device. For remote code validation, the code can be configured to either of the special values `REMOTE_CODE` (numeric code) or `REMOTE_CODE_TEXT` (text code). In this case, local code validation is bypassed but the frontend will still show a numeric or text code dialog. Use `command_template` to send the code to the remote device. Example configurations for remote code validation [can be found here](./#configurations-with-remote-code-validation)."
	CodeArmRequired        bool   // "If true the code is required to arm the alarm. If false the code is not validated."
	CodeDisarmRequired     bool   // "If true the code is required to disarm the alarm. If false the code is not validated."
	CodeTriggerRequired    bool   // "If true the code is required to trigger the alarm. If false the code is not validated."
	CommandTemplate        string // "The [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) used for the command payload. Available variables: `action` and `code`."
	CommandFunc            func(mqtt.Message, mqtt.Client)
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc     func() string
	Name                   string // "The name of the alarm."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadArmAway         string // "The payload to set armed-away mode on your Alarm Panel."
	PayloadArmCustomBypass string // "The payload to set armed-custom-bypass mode on your Alarm Panel."
	PayloadArmHome         string // "The payload to set armed-home mode on your Alarm Panel."
	PayloadArmNight        string // "The payload to set armed-night mode on your Alarm Panel."
	PayloadArmVacation     string // "The payload to set armed-vacation mode on your Alarm Panel."
	PayloadAvailable       string // "The payload that represents the available state."
	PayloadDisarm          string // "The payload to disarm your Alarm Panel."
	PayloadNotAvailable    string // "The payload that represents the unavailable state."
	PayloadTrigger         string // "The payload to trigger the alarm on your Alarm Panel."
	Qos                    int    // "The maximum QoS level of the state topic."
	Retain                 bool   // "If the published message should have the retain flag on or not."
	StateFunc              func() string
	UniqueId               string // "An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value."
}

func NewAlarmControlPanelOptions() *AlarmControlPanelOptions {
	return &AlarmControlPanelOptions{}
}
func (o *AlarmControlPanelOptions) GetStates() *AlarmControlPanelState {
	return &o.States
}
func (o *AlarmControlPanelOptions) SetAvailabilityMode(mode string) *AlarmControlPanelOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *AlarmControlPanelOptions) SetAvailabilityTemplate(template string) *AlarmControlPanelOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *AlarmControlPanelOptions) SetAvailabilityFunc(f func() string) *AlarmControlPanelOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *AlarmControlPanelOptions) SetCode(code string) *AlarmControlPanelOptions {
	o.Code = code
	return o
}
func (o *AlarmControlPanelOptions) SetCodeArmRequired(required bool) *AlarmControlPanelOptions {
	o.CodeArmRequired = required
	return o
}
func (o *AlarmControlPanelOptions) SetCodeDisarmRequired(required bool) *AlarmControlPanelOptions {
	o.CodeDisarmRequired = required
	return o
}
func (o *AlarmControlPanelOptions) SetCodeTriggerRequired(required bool) *AlarmControlPanelOptions {
	o.CodeTriggerRequired = required
	return o
}
func (o *AlarmControlPanelOptions) SetCommandTemplate(template string) *AlarmControlPanelOptions {
	o.CommandTemplate = template
	return o
}
func (o *AlarmControlPanelOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *AlarmControlPanelOptions {
	o.CommandFunc = f
	return o
}
func (o *AlarmControlPanelOptions) SetEnabledByDefault(d bool) *AlarmControlPanelOptions {
	o.EnabledByDefault = d
	return o
}
func (o *AlarmControlPanelOptions) SetEncoding(encoding string) *AlarmControlPanelOptions {
	o.Encoding = encoding
	return o
}
func (o *AlarmControlPanelOptions) SetEntityCategory(category string) *AlarmControlPanelOptions {
	o.EntityCategory = category
	return o
}
func (o *AlarmControlPanelOptions) SetIcon(icon string) *AlarmControlPanelOptions {
	o.Icon = icon
	return o
}
func (o *AlarmControlPanelOptions) SetJsonAttributesTemplate(template string) *AlarmControlPanelOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *AlarmControlPanelOptions) SetJsonAttributesFunc(f func() string) *AlarmControlPanelOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *AlarmControlPanelOptions) SetName(name string) *AlarmControlPanelOptions {
	o.Name = name
	return o
}
func (o *AlarmControlPanelOptions) SetObjectId(id string) *AlarmControlPanelOptions {
	o.ObjectId = id
	return o
}
func (o *AlarmControlPanelOptions) SetPayloadArmAway(away string) *AlarmControlPanelOptions {
	o.PayloadArmAway = away
	return o
}
func (o *AlarmControlPanelOptions) SetPayloadArmCustomBypass(bypass string) *AlarmControlPanelOptions {
	o.PayloadArmCustomBypass = bypass
	return o
}
func (o *AlarmControlPanelOptions) SetPayloadArmHome(home string) *AlarmControlPanelOptions {
	o.PayloadArmHome = home
	return o
}
func (o *AlarmControlPanelOptions) SetPayloadArmNight(night string) *AlarmControlPanelOptions {
	o.PayloadArmNight = night
	return o
}
func (o *AlarmControlPanelOptions) SetPayloadArmVacation(vacation string) *AlarmControlPanelOptions {
	o.PayloadArmVacation = vacation
	return o
}
func (o *AlarmControlPanelOptions) SetPayloadAvailable(available string) *AlarmControlPanelOptions {
	o.PayloadAvailable = available
	return o
}
func (o *AlarmControlPanelOptions) SetPayloadDisarm(disarm string) *AlarmControlPanelOptions {
	o.PayloadDisarm = disarm
	return o
}
func (o *AlarmControlPanelOptions) SetPayloadNotAvailable(available string) *AlarmControlPanelOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *AlarmControlPanelOptions) SetPayloadTrigger(trigger string) *AlarmControlPanelOptions {
	o.PayloadTrigger = trigger
	return o
}
func (o *AlarmControlPanelOptions) SetQos(qos int) *AlarmControlPanelOptions {
	o.Qos = qos
	return o
}
func (o *AlarmControlPanelOptions) SetRetain(retain bool) *AlarmControlPanelOptions {
	o.Retain = retain
	return o
}
func (o *AlarmControlPanelOptions) SetStateFunc(f func() string) *AlarmControlPanelOptions {
	o.StateFunc = f
	return o
}
func (o *AlarmControlPanelOptions) SetUniqueId(id string) *AlarmControlPanelOptions {
	o.UniqueId = id
	return o
}
func (o *AlarmControlPanelOptions) SetValueTemplate(template string) *AlarmControlPanelOptions {
	o.ValueTemplate = template
	return o
}
