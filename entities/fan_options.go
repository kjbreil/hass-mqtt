package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type FanOptions struct {
	states                     FanState // External state update location
	availabilityMode           string   // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate       string   // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc           func() string
	commandTemplate            string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	commandFunc                func(mqtt.Message, mqtt.Client)
	enabledByDefault           bool   // "Flag which defines if the entity should be enabled when first added."
	encoding                   string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory             string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	icon                       string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate     string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc         func() string
	name                       string // "The name of the fan."
	objectId                   string // "Used instead of `name` for automatic generation of `entity_id`"
	optimistic                 bool   // "Flag that defines if fan works in optimistic mode"
	oscillationCommandTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `oscillation_command_topic`."
	oscillationCommandFunc     func(mqtt.Message, mqtt.Client)
	oscillationStateFunc       func() string
	oscillationValueTemplate   string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the oscillation."
	payloadAvailable           string // "The payload that represents the available state."
	payloadNotAvailable        string // "The payload that represents the unavailable state."
	payloadOff                 string // "The payload that represents the stop state."
	payloadOn                  string // "The payload that represents the running state."
	payloadOscillationOff      string // "The payload that represents the oscillation off state."
	payloadOscillationOn       string // "The payload that represents the oscillation on state."
	payloadResetPercentage     string // "A special payload that resets the `percentage` state attribute to `None` when received at the `percentage_state_topic`."
	payloadResetPresetMode     string // "A special payload that resets the `preset_mode` state attribute to `None` when received at the `preset_mode_state_topic`."
	percentageCommandTemplate  string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `percentage_command_topic`."
	percentageCommandFunc      func(mqtt.Message, mqtt.Client)
	percentageStateFunc        func() string
	percentageValueTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `percentage` value from the payload received on `percentage_state_topic`."
	presetModeCommandTemplate  string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `preset_mode_command_topic`."
	presetModeCommandFunc      func(mqtt.Message, mqtt.Client)
	presetModeStateFunc        func() string
	presetModeValueTemplate    string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	presetModes                ([]string) // "List of preset modes this fan is capable of running at. Common examples include `auto`, `smart`, `whoosh`, `eco` and `breeze`."
	qos                        int        // "The maximum QoS level of the state topic."
	retain                     bool       // "If the published message should have the retain flag on or not."
	speedRangeMax              int        // "The maximum of numeric output range (representing 100 %). The number of speeds within the `speed_range` / `100` will determine the `percentage_step`."
	speedRangeMin              int        // "The minimum of numeric output range (`off` not included, so `speed_range_min` - `1` represents 0 %). The number of speeds within the speed_range / 100 will determine the `percentage_step`."
	stateFunc                  func() string
	stateValueTemplate         string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state."
	uniqueId                   string // "An ID that uniquely identifies this fan. If two fans have the same unique ID, Home Assistant will raise an exception."
}

func NewFanOptions() *FanOptions {
	return &FanOptions{}
}
func (o *FanOptions) States() *FanState {
	return &o.states
}
func (o *FanOptions) AvailabilityMode(mode string) *FanOptions {
	o.availabilityMode = mode
	return o
}
func (o *FanOptions) AvailabilityTemplate(template string) *FanOptions {
	o.availabilityTemplate = template
	return o
}
func (o *FanOptions) AvailabilityFunc(f func() string) *FanOptions {
	o.availabilityFunc = f
	return o
}
func (o *FanOptions) CommandTemplate(template string) *FanOptions {
	o.commandTemplate = template
	return o
}
func (o *FanOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *FanOptions {
	o.commandFunc = f
	return o
}
func (o *FanOptions) EnabledByDefault(d bool) *FanOptions {
	o.enabledByDefault = d
	return o
}
func (o *FanOptions) Encoding(encoding string) *FanOptions {
	o.encoding = encoding
	return o
}
func (o *FanOptions) EntityCategory(category string) *FanOptions {
	o.entityCategory = category
	return o
}
func (o *FanOptions) Icon(icon string) *FanOptions {
	o.icon = icon
	return o
}
func (o *FanOptions) JsonAttributesTemplate(template string) *FanOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *FanOptions) JsonAttributesFunc(f func() string) *FanOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *FanOptions) Name(name string) *FanOptions {
	o.name = name
	return o
}
func (o *FanOptions) ObjectId(id string) *FanOptions {
	o.objectId = id
	return o
}
func (o *FanOptions) Optimistic(optimistic bool) *FanOptions {
	o.optimistic = optimistic
	return o
}
func (o *FanOptions) OscillationCommandTemplate(template string) *FanOptions {
	o.oscillationCommandTemplate = template
	return o
}
func (o *FanOptions) OscillationCommandFunc(f func(mqtt.Message, mqtt.Client)) *FanOptions {
	o.oscillationCommandFunc = f
	return o
}
func (o *FanOptions) OscillationStateFunc(f func() string) *FanOptions {
	o.oscillationStateFunc = f
	return o
}
func (o *FanOptions) EnableOscillation() *FanOptions {
	o.oscillationStateFunc = func() string {
		return o.states.Oscillation
	}
	o.oscillationCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Oscillation = string(message.Payload())
	}
	return o
}
func (o *FanOptions) OscillationValueTemplate(template string) *FanOptions {
	o.oscillationValueTemplate = template
	return o
}
func (o *FanOptions) PayloadAvailable(available string) *FanOptions {
	o.payloadAvailable = available
	return o
}
func (o *FanOptions) PayloadNotAvailable(available string) *FanOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *FanOptions) PayloadOff(off string) *FanOptions {
	o.payloadOff = off
	return o
}
func (o *FanOptions) PayloadOn(on string) *FanOptions {
	o.payloadOn = on
	return o
}
func (o *FanOptions) PayloadOscillationOff(off string) *FanOptions {
	o.payloadOscillationOff = off
	return o
}
func (o *FanOptions) PayloadOscillationOn(on string) *FanOptions {
	o.payloadOscillationOn = on
	return o
}
func (o *FanOptions) PayloadResetPercentage(percentage string) *FanOptions {
	o.payloadResetPercentage = percentage
	return o
}
func (o *FanOptions) PayloadResetPresetMode(mode string) *FanOptions {
	o.payloadResetPresetMode = mode
	return o
}
func (o *FanOptions) PercentageCommandTemplate(template string) *FanOptions {
	o.percentageCommandTemplate = template
	return o
}
func (o *FanOptions) PercentageCommandFunc(f func(mqtt.Message, mqtt.Client)) *FanOptions {
	o.percentageCommandFunc = f
	return o
}
func (o *FanOptions) PercentageStateFunc(f func() string) *FanOptions {
	o.percentageStateFunc = f
	return o
}
func (o *FanOptions) EnablePercentage() *FanOptions {
	o.percentageStateFunc = func() string {
		return o.states.Percentage
	}
	o.percentageCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Percentage = string(message.Payload())
	}
	return o
}
func (o *FanOptions) PercentageValueTemplate(template string) *FanOptions {
	o.percentageValueTemplate = template
	return o
}
func (o *FanOptions) PresetModeCommandTemplate(template string) *FanOptions {
	o.presetModeCommandTemplate = template
	return o
}
func (o *FanOptions) PresetModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *FanOptions {
	o.presetModeCommandFunc = f
	return o
}
func (o *FanOptions) PresetModeStateFunc(f func() string) *FanOptions {
	o.presetModeStateFunc = f
	return o
}
func (o *FanOptions) EnablePresetMode() *FanOptions {
	o.presetModeStateFunc = func() string {
		return o.states.PresetMode
	}
	o.presetModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.PresetMode = string(message.Payload())
	}
	return o
}
func (o *FanOptions) PresetModeValueTemplate(template string) *FanOptions {
	o.presetModeValueTemplate = template
	return o
}
func (o *FanOptions) PresetModes(modes []string) *FanOptions {
	o.presetModes = modes
	return o
}
func (o *FanOptions) Qos(qos int) *FanOptions {
	o.qos = qos
	return o
}
func (o *FanOptions) Retain(retain bool) *FanOptions {
	o.retain = retain
	return o
}
func (o *FanOptions) SpeedRangeMax(max int) *FanOptions {
	o.speedRangeMax = max
	return o
}
func (o *FanOptions) SpeedRangeMin(min int) *FanOptions {
	o.speedRangeMin = min
	return o
}
func (o *FanOptions) StateFunc(f func() string) *FanOptions {
	o.stateFunc = f
	return o
}
func (o *FanOptions) StateValueTemplate(template string) *FanOptions {
	o.stateValueTemplate = template
	return o
}
func (o *FanOptions) UniqueId(id string) *FanOptions {
	o.uniqueId = id
	return o
}
