package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type AlarmControlPanelOptions struct {
	states                 AlarmControlPanelState // External state update location
	availabilityMode       string                 // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string                 // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	code                   string // "If defined, specifies a code to enable or disable the alarm in the frontend. Note that the code is validated locally and blocks sending MQTT messages to the remote device. For remote code validation, the code can be configured to either of the special values `REMOTE_CODE` (numeric code) or `REMOTE_CODE_TEXT` (text code). In this case, local code validation is bypassed but the frontend will still show a numeric or text code dialog. Use `command_template` to send the code to the remote device. Example configurations for remote code validation [can be found here](#configurations-with-remote-code-validation)."
	codeArmRequired        bool   // "If true the code is required to arm the alarm. If false the code is not validated."
	codeDisarmRequired     bool   // "If true the code is required to disarm the alarm. If false the code is not validated."
	codeTriggerRequired    bool   // "If true the code is required to trigger the alarm. If false the code is not validated."
	commandTemplate        string // "The [template](/docs/configuration/templating/#using-command-templates-with-mqtt) used for the command payload. Available variables: `action` and `code`."
	commandFunc            func(mqtt.Message, mqtt.Client)
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	entityPicture          string // "Picture URL for the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name of the alarm. Can be set to `null` if only the device name is relevant."
	objectId               string // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	payloadArmAway         string // "The payload to set armed-away mode on your Alarm Panel."
	payloadArmCustomBypass string // "The payload to set armed-custom-bypass mode on your Alarm Panel."
	payloadArmHome         string // "The payload to set armed-home mode on your Alarm Panel."
	payloadArmNight        string // "The payload to set armed-night mode on your Alarm Panel."
	payloadArmVacation     string // "The payload to set armed-vacation mode on your Alarm Panel."
	payloadAvailable       string // "The payload that represents the available state."
	payloadDisarm          string // "The payload to disarm your Alarm Panel."
	payloadNotAvailable    string // "The payload that represents the unavailable state."
	payloadTrigger         string // "The payload to trigger the alarm on your Alarm Panel."
	platform               string // "Must be `alarm_control_panel`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                    int    // "The maximum QoS level to be used when receiving and publishing messages."
	retain                 bool   // "If the published message should have the retain flag on or not."
	stateFunc              func() string
	supportedFeatures      ([]string) // "A list of features that the alarm control panel supports. The available list options are `arm_home`, `arm_away`, `arm_night`, `arm_vacation`, `arm_custom_bypass`, and `trigger`."
	uniqueId               string     // "An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	valueTemplate          string     // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the value."
}

func NewAlarmControlPanelOptions() *AlarmControlPanelOptions {
	return &AlarmControlPanelOptions{}
}
func (o *AlarmControlPanelOptions) States() *AlarmControlPanelState {
	return &o.states
}
func (o *AlarmControlPanelOptions) AvailabilityMode(mode string) *AlarmControlPanelOptions {
	o.availabilityMode = mode
	return o
}
func (o *AlarmControlPanelOptions) AvailabilityTemplate(template string) *AlarmControlPanelOptions {
	o.availabilityTemplate = template
	return o
}
func (o *AlarmControlPanelOptions) AvailabilityFunc(f func() string) *AlarmControlPanelOptions {
	o.availabilityFunc = f
	return o
}
func (o *AlarmControlPanelOptions) Code(code string) *AlarmControlPanelOptions {
	o.code = code
	return o
}
func (o *AlarmControlPanelOptions) CodeArmRequired(required bool) *AlarmControlPanelOptions {
	o.codeArmRequired = required
	return o
}
func (o *AlarmControlPanelOptions) CodeDisarmRequired(required bool) *AlarmControlPanelOptions {
	o.codeDisarmRequired = required
	return o
}
func (o *AlarmControlPanelOptions) CodeTriggerRequired(required bool) *AlarmControlPanelOptions {
	o.codeTriggerRequired = required
	return o
}
func (o *AlarmControlPanelOptions) CommandTemplate(template string) *AlarmControlPanelOptions {
	o.commandTemplate = template
	return o
}
func (o *AlarmControlPanelOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *AlarmControlPanelOptions {
	o.commandFunc = f
	return o
}
func (o *AlarmControlPanelOptions) EnabledByDefault(d bool) *AlarmControlPanelOptions {
	o.enabledByDefault = d
	return o
}
func (o *AlarmControlPanelOptions) Encoding(encoding string) *AlarmControlPanelOptions {
	o.encoding = encoding
	return o
}
func (o *AlarmControlPanelOptions) EntityCategory(category string) *AlarmControlPanelOptions {
	o.entityCategory = category
	return o
}
func (o *AlarmControlPanelOptions) EntityPicture(picture string) *AlarmControlPanelOptions {
	o.entityPicture = picture
	return o
}
func (o *AlarmControlPanelOptions) Icon(icon string) *AlarmControlPanelOptions {
	o.icon = icon
	return o
}
func (o *AlarmControlPanelOptions) JsonAttributesTemplate(template string) *AlarmControlPanelOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *AlarmControlPanelOptions) JsonAttributesFunc(f func() string) *AlarmControlPanelOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *AlarmControlPanelOptions) Name(name string) *AlarmControlPanelOptions {
	o.name = name
	return o
}
func (o *AlarmControlPanelOptions) ObjectId(id string) *AlarmControlPanelOptions {
	o.objectId = id
	return o
}
func (o *AlarmControlPanelOptions) PayloadArmAway(away string) *AlarmControlPanelOptions {
	o.payloadArmAway = away
	return o
}
func (o *AlarmControlPanelOptions) PayloadArmCustomBypass(bypass string) *AlarmControlPanelOptions {
	o.payloadArmCustomBypass = bypass
	return o
}
func (o *AlarmControlPanelOptions) PayloadArmHome(home string) *AlarmControlPanelOptions {
	o.payloadArmHome = home
	return o
}
func (o *AlarmControlPanelOptions) PayloadArmNight(night string) *AlarmControlPanelOptions {
	o.payloadArmNight = night
	return o
}
func (o *AlarmControlPanelOptions) PayloadArmVacation(vacation string) *AlarmControlPanelOptions {
	o.payloadArmVacation = vacation
	return o
}
func (o *AlarmControlPanelOptions) PayloadAvailable(available string) *AlarmControlPanelOptions {
	o.payloadAvailable = available
	return o
}
func (o *AlarmControlPanelOptions) PayloadDisarm(disarm string) *AlarmControlPanelOptions {
	o.payloadDisarm = disarm
	return o
}
func (o *AlarmControlPanelOptions) PayloadNotAvailable(available string) *AlarmControlPanelOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *AlarmControlPanelOptions) PayloadTrigger(trigger string) *AlarmControlPanelOptions {
	o.payloadTrigger = trigger
	return o
}
func (o *AlarmControlPanelOptions) Platform(platform string) *AlarmControlPanelOptions {
	o.platform = platform
	return o
}
func (o *AlarmControlPanelOptions) Qos(qos int) *AlarmControlPanelOptions {
	o.qos = qos
	return o
}
func (o *AlarmControlPanelOptions) Retain(retain bool) *AlarmControlPanelOptions {
	o.retain = retain
	return o
}
func (o *AlarmControlPanelOptions) StateFunc(f func() string) *AlarmControlPanelOptions {
	o.stateFunc = f
	return o
}
func (o *AlarmControlPanelOptions) SupportedFeatures(features []string) *AlarmControlPanelOptions {
	o.supportedFeatures = features
	return o
}
func (o *AlarmControlPanelOptions) UniqueId(id string) *AlarmControlPanelOptions {
	o.uniqueId = id
	return o
}
func (o *AlarmControlPanelOptions) ValueTemplate(template string) *AlarmControlPanelOptions {
	o.valueTemplate = template
	return o
}
