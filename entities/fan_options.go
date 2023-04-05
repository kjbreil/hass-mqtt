package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type FanOptions struct {
	States                     FanState // External state update location
	AvailabilityMode           string   // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate       string   // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc           func() string
	CommandTemplate            string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandFunc                func(mqtt.Message, mqtt.Client)
	EnabledByDefault           bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding                   string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory             string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                       string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate     string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc         func() string
	Name                       string // "The name of the fan."
	ObjectId                   string // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                 bool   // "Flag that defines if fan works in optimistic mode"
	OscillationCommandTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `oscillation_command_topic`."
	OscillationCommandFunc     func(mqtt.Message, mqtt.Client)
	OscillationStateFunc       func() string
	OscillationValueTemplate   string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the oscillation."
	PayloadAvailable           string // "The payload that represents the available state."
	PayloadNotAvailable        string // "The payload that represents the unavailable state."
	PayloadOff                 string // "The payload that represents the stop state."
	PayloadOn                  string // "The payload that represents the running state."
	PayloadOscillationOff      string // "The payload that represents the oscillation off state."
	PayloadOscillationOn       string // "The payload that represents the oscillation on state."
	PayloadResetPercentage     string // "A special payload that resets the `percentage` state attribute to `None` when received at the `percentage_state_topic`."
	PayloadResetPresetMode     string // "A special payload that resets the `preset_mode` state attribute to `None` when received at the `preset_mode_state_topic`."
	PercentageCommandTemplate  string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `percentage_command_topic`."
	PercentageCommandFunc      func(mqtt.Message, mqtt.Client)
	PercentageStateFunc        func() string
	PercentageValueTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `percentage` value from the payload received on `percentage_state_topic`."
	PresetModeCommandTemplate  string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommandFunc      func(mqtt.Message, mqtt.Client)
	PresetModeStateFunc        func() string
	PresetModeValueTemplate    string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                ([]string) // "List of preset modes this fan is capable of running at. Common examples include `auto`, `smart`, `whoosh`, `eco` and `breeze`."
	Qos                        int        // "The maximum QoS level of the state topic."
	Retain                     bool       // "If the published message should have the retain flag on or not."
	SpeedRangeMax              int        // "The maximum of numeric output range (representing 100 %). The number of speeds within the `speed_range` / `100` will determine the `percentage_step`."
	SpeedRangeMin              int        // "The minimum of numeric output range (`off` not included, so `speed_range_min` - `1` represents 0 %). The number of speeds within the speed_range / 100 will determine the `percentage_step`."
	StateFunc                  func() string
	StateValueTemplate         string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state."
	UniqueId                   string // "An ID that uniquely identifies this fan. If two fans have the same unique ID, Home Assistant will raise an exception."
}

func NewFanOptions() *FanOptions {
	return &FanOptions{}
}
func (o *FanOptions) GetStates() *FanState {
	return &o.States
}
func (o *FanOptions) SetAvailabilityMode(mode string) *FanOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *FanOptions) SetAvailabilityTemplate(template string) *FanOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *FanOptions) SetAvailabilityFunc(f func() string) *FanOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *FanOptions) SetCommandTemplate(template string) *FanOptions {
	o.CommandTemplate = template
	return o
}
func (o *FanOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *FanOptions {
	o.CommandFunc = f
	return o
}
func (o *FanOptions) SetEnabledByDefault(d bool) *FanOptions {
	o.EnabledByDefault = d
	return o
}
func (o *FanOptions) SetEncoding(encoding string) *FanOptions {
	o.Encoding = encoding
	return o
}
func (o *FanOptions) SetEntityCategory(category string) *FanOptions {
	o.EntityCategory = category
	return o
}
func (o *FanOptions) SetIcon(icon string) *FanOptions {
	o.Icon = icon
	return o
}
func (o *FanOptions) SetJsonAttributesTemplate(template string) *FanOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *FanOptions) SetJsonAttributesFunc(f func() string) *FanOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *FanOptions) SetName(name string) *FanOptions {
	o.Name = name
	return o
}
func (o *FanOptions) SetObjectId(id string) *FanOptions {
	o.ObjectId = id
	return o
}
func (o *FanOptions) SetOptimistic(optimistic bool) *FanOptions {
	o.Optimistic = optimistic
	return o
}
func (o *FanOptions) SetOscillationCommandTemplate(template string) *FanOptions {
	o.OscillationCommandTemplate = template
	return o
}
func (o *FanOptions) SetOscillationCommandFunc(f func(mqtt.Message, mqtt.Client)) *FanOptions {
	o.OscillationCommandFunc = f
	return o
}
func (o *FanOptions) SetOscillationStateFunc(f func() string) *FanOptions {
	o.OscillationStateFunc = f
	return o
}
func (o *FanOptions) HasOscillation() *FanOptions {
	o.OscillationStateFunc = func() string {
		return o.States.Oscillation
	}
	o.OscillationCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Oscillation = string(message.Payload())
	}
	return o
}
func (o *FanOptions) SetOscillationValueTemplate(template string) *FanOptions {
	o.OscillationValueTemplate = template
	return o
}
func (o *FanOptions) SetPayloadAvailable(available string) *FanOptions {
	o.PayloadAvailable = available
	return o
}
func (o *FanOptions) SetPayloadNotAvailable(available string) *FanOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *FanOptions) SetPayloadOff(off string) *FanOptions {
	o.PayloadOff = off
	return o
}
func (o *FanOptions) SetPayloadOn(on string) *FanOptions {
	o.PayloadOn = on
	return o
}
func (o *FanOptions) SetPayloadOscillationOff(off string) *FanOptions {
	o.PayloadOscillationOff = off
	return o
}
func (o *FanOptions) SetPayloadOscillationOn(on string) *FanOptions {
	o.PayloadOscillationOn = on
	return o
}
func (o *FanOptions) SetPayloadResetPercentage(percentage string) *FanOptions {
	o.PayloadResetPercentage = percentage
	return o
}
func (o *FanOptions) SetPayloadResetPresetMode(mode string) *FanOptions {
	o.PayloadResetPresetMode = mode
	return o
}
func (o *FanOptions) SetPercentageCommandTemplate(template string) *FanOptions {
	o.PercentageCommandTemplate = template
	return o
}
func (o *FanOptions) SetPercentageCommandFunc(f func(mqtt.Message, mqtt.Client)) *FanOptions {
	o.PercentageCommandFunc = f
	return o
}
func (o *FanOptions) SetPercentageStateFunc(f func() string) *FanOptions {
	o.PercentageStateFunc = f
	return o
}
func (o *FanOptions) HasPercentage() *FanOptions {
	o.PercentageStateFunc = func() string {
		return o.States.Percentage
	}
	o.PercentageCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Percentage = string(message.Payload())
	}
	return o
}
func (o *FanOptions) SetPercentageValueTemplate(template string) *FanOptions {
	o.PercentageValueTemplate = template
	return o
}
func (o *FanOptions) SetPresetModeCommandTemplate(template string) *FanOptions {
	o.PresetModeCommandTemplate = template
	return o
}
func (o *FanOptions) SetPresetModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *FanOptions {
	o.PresetModeCommandFunc = f
	return o
}
func (o *FanOptions) SetPresetModeStateFunc(f func() string) *FanOptions {
	o.PresetModeStateFunc = f
	return o
}
func (o *FanOptions) HasPresetMode() *FanOptions {
	o.PresetModeStateFunc = func() string {
		return o.States.PresetMode
	}
	o.PresetModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.PresetMode = string(message.Payload())
	}
	return o
}
func (o *FanOptions) SetPresetModeValueTemplate(template string) *FanOptions {
	o.PresetModeValueTemplate = template
	return o
}
func (o *FanOptions) SetPresetModes(modes []string) *FanOptions {
	o.PresetModes = modes
	return o
}
func (o *FanOptions) SetQos(qos int) *FanOptions {
	o.Qos = qos
	return o
}
func (o *FanOptions) SetRetain(retain bool) *FanOptions {
	o.Retain = retain
	return o
}
func (o *FanOptions) SetSpeedRangeMax(max int) *FanOptions {
	o.SpeedRangeMax = max
	return o
}
func (o *FanOptions) SetSpeedRangeMin(min int) *FanOptions {
	o.SpeedRangeMin = min
	return o
}
func (o *FanOptions) SetStateFunc(f func() string) *FanOptions {
	o.StateFunc = f
	return o
}
func (o *FanOptions) SetStateValueTemplate(template string) *FanOptions {
	o.StateValueTemplate = template
	return o
}
func (o *FanOptions) SetUniqueId(id string) *FanOptions {
	o.UniqueId = id
	return o
}
