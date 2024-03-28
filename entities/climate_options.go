package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type ClimateOptions struct {
	states                         ClimateState // External state update location
	actionTemplate                 string       // "A template to render the value received on the `action_topic` with."
	actionFunc                     func() string
	auxCommandFunc                 func(mqtt.Message, mqtt.Client)
	auxStateTemplate               string // "A template to render the value received on the `aux_state_topic` with."
	auxStateFunc                   func() string
	availabilityMode               string // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate           string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc               func() string
	currentHumidityTemplate        string // "A template with which the value received on `current_humidity_topic` will be rendered."
	currentHumidityFunc            func() string
	currentTemperatureTemplate     string // "A template with which the value received on `current_temperature_topic` will be rendered."
	currentTemperatureFunc         func() string
	enabledByDefault               bool   // "Flag which defines if the entity should be enabled when first added."
	encoding                       string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory                 string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	fanModeCommandTemplate         string // "A template to render the value sent to the `fan_mode_command_topic` with."
	fanModeCommandFunc             func(mqtt.Message, mqtt.Client)
	fanModeStateTemplate           string // "A template to render the value received on the `fan_mode_state_topic` with."
	fanModeStateFunc               func() string
	fanModes                       ([]string) // "A list of supported fan modes."
	icon                           string     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	initial                        int        // "Set the initial target temperature."
	jsonAttributesTemplate         string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc             func() string
	maxHumidity                    int     // "The minimum target humidity percentage that can be set."
	maxTemp                        float64 // "Maximum set point available."
	minHumidity                    int     // "The maximum target humidity percentage that can be set."
	minTemp                        float64 // "Minimum set point available."
	modeCommandTemplate            string  // "A template to render the value sent to the `mode_command_topic` with."
	modeCommandFunc                func(mqtt.Message, mqtt.Client)
	modeStateTemplate              string // "A template to render the value received on the `mode_state_topic` with."
	modeStateFunc                  func() string
	modes                          ([]string) // "A list of supported modes. Needs to be a subset of the default values."
	name                           string     // "The name of the HVAC."
	objectId                       string     // "Used instead of `name` for automatic generation of `entity_id`"
	optimistic                     bool       // "Flag that defines if the climate works in optimistic mode"
	payloadAvailable               string     // "The payload that represents the available state."
	payloadNotAvailable            string     // "The payload that represents the unavailable state."
	payloadOff                     string     // "The payload that represents disabled state."
	payloadOn                      string     // "The payload that represents enabled state."
	precision                      float64    // "The desired precision for this device. Can be used to match your actual thermostat's precision. Supported values are `0.1`, `0.5` and `1.0`."
	presetModeCommandTemplate      string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `preset_mode_command_topic`."
	presetModeCommandFunc          func(mqtt.Message, mqtt.Client)
	presetModeStateFunc            func() string
	presetModeValueTemplate        string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	presetModes                    ([]string) // "List of preset modes this climate is supporting. Common examples include `eco`, `away`, `boost`, `comfort`, `home`, `sleep` and `activity`."
	qos                            int        // "The maximum QoS level to be used when receiving and publishing messages."
	retain                         bool       // "Defines if published messages should have the retain flag set."
	swingModeCommandTemplate       string     // "A template to render the value sent to the `swing_mode_command_topic` with."
	swingModeCommandFunc           func(mqtt.Message, mqtt.Client)
	swingModeStateTemplate         string // "A template to render the value received on the `swing_mode_state_topic` with."
	swingModeStateFunc             func() string
	swingModes                     ([]string) // "A list of supported swing modes."
	targetHumidityCommandTemplate  string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `target_humidity_command_topic`."
	targetHumidityCommandFunc      func(mqtt.Message, mqtt.Client)
	targetHumidityStateTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the climate `target_humidity` state."
	targetHumidityStateFunc        func() string
	tempStep                       float64 // "Step size for temperature set point."
	temperatureCommandTemplate     string  // "A template to render the value sent to the `temperature_command_topic` with."
	temperatureCommandFunc         func(mqtt.Message, mqtt.Client)
	temperatureHighCommandTemplate string // "A template to render the value sent to the `temperature_high_command_topic` with."
	temperatureHighCommandFunc     func(mqtt.Message, mqtt.Client)
	temperatureHighStateTemplate   string // "A template to render the value received on the `temperature_high_state_topic` with. A `\"None\"` value received will reset the temperature high set point. Empty values (`'''`) will be ignored."
	temperatureHighStateFunc       func() string
	temperatureLowCommandTemplate  string // "A template to render the value sent to the `temperature_low_command_topic` with."
	temperatureLowCommandFunc      func(mqtt.Message, mqtt.Client)
	temperatureLowStateTemplate    string // "A template to render the value received on the `temperature_low_state_topic` with. A `\"None\"` value received will reset the temperature low set point. Empty values (`'''`) will be ignored."
	temperatureLowStateFunc        func() string
	temperatureStateTemplate       string // "A template to render the value received on the `temperature_state_topic` with."
	temperatureStateFunc           func() string
	temperatureUnit                string // "Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit."
	uniqueId                       string // "An ID that uniquely identifies this HVAC device. If two HVAC devices have the same unique ID, Home Assistant will raise an exception."
	valueTemplate                  string // "Default template to render the payloads on *all* `*_state_topic`s with."
}

func NewClimateOptions() *ClimateOptions {
	return &ClimateOptions{}
}
func (o *ClimateOptions) States() *ClimateState {
	return &o.states
}
func (o *ClimateOptions) ActionTemplate(template string) *ClimateOptions {
	o.actionTemplate = template
	return o
}
func (o *ClimateOptions) ActionFunc(f func() string) *ClimateOptions {
	o.actionFunc = f
	return o
}
func (o *ClimateOptions) AuxCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.auxCommandFunc = f
	return o
}
func (o *ClimateOptions) AuxStateTemplate(template string) *ClimateOptions {
	o.auxStateTemplate = template
	return o
}
func (o *ClimateOptions) AuxStateFunc(f func() string) *ClimateOptions {
	o.auxStateFunc = f
	return o
}
func (o *ClimateOptions) EnableAux() *ClimateOptions {
	o.auxStateFunc = func() string {
		return o.states.Aux
	}
	o.auxCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Aux = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) AvailabilityMode(mode string) *ClimateOptions {
	o.availabilityMode = mode
	return o
}
func (o *ClimateOptions) AvailabilityTemplate(template string) *ClimateOptions {
	o.availabilityTemplate = template
	return o
}
func (o *ClimateOptions) AvailabilityFunc(f func() string) *ClimateOptions {
	o.availabilityFunc = f
	return o
}
func (o *ClimateOptions) CurrentHumidityTemplate(template string) *ClimateOptions {
	o.currentHumidityTemplate = template
	return o
}
func (o *ClimateOptions) CurrentHumidityFunc(f func() string) *ClimateOptions {
	o.currentHumidityFunc = f
	return o
}
func (o *ClimateOptions) CurrentTemperatureTemplate(template string) *ClimateOptions {
	o.currentTemperatureTemplate = template
	return o
}
func (o *ClimateOptions) CurrentTemperatureFunc(f func() string) *ClimateOptions {
	o.currentTemperatureFunc = f
	return o
}
func (o *ClimateOptions) EnabledByDefault(d bool) *ClimateOptions {
	o.enabledByDefault = d
	return o
}
func (o *ClimateOptions) Encoding(encoding string) *ClimateOptions {
	o.encoding = encoding
	return o
}
func (o *ClimateOptions) EntityCategory(category string) *ClimateOptions {
	o.entityCategory = category
	return o
}
func (o *ClimateOptions) FanModeCommandTemplate(template string) *ClimateOptions {
	o.fanModeCommandTemplate = template
	return o
}
func (o *ClimateOptions) FanModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.fanModeCommandFunc = f
	return o
}
func (o *ClimateOptions) FanModeStateTemplate(template string) *ClimateOptions {
	o.fanModeStateTemplate = template
	return o
}
func (o *ClimateOptions) FanModeStateFunc(f func() string) *ClimateOptions {
	o.fanModeStateFunc = f
	return o
}
func (o *ClimateOptions) EnableFanMode() *ClimateOptions {
	o.fanModeStateFunc = func() string {
		return o.states.FanMode
	}
	o.fanModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.FanMode = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) FanModes(modes []string) *ClimateOptions {
	o.fanModes = modes
	return o
}
func (o *ClimateOptions) Icon(icon string) *ClimateOptions {
	o.icon = icon
	return o
}
func (o *ClimateOptions) Initial(initial int) *ClimateOptions {
	o.initial = initial
	return o
}
func (o *ClimateOptions) JsonAttributesTemplate(template string) *ClimateOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *ClimateOptions) JsonAttributesFunc(f func() string) *ClimateOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *ClimateOptions) MaxHumidity(humidity int) *ClimateOptions {
	o.maxHumidity = humidity
	return o
}
func (o *ClimateOptions) MaxTemp(temp float64) *ClimateOptions {
	o.maxTemp = temp
	return o
}
func (o *ClimateOptions) MinHumidity(humidity int) *ClimateOptions {
	o.minHumidity = humidity
	return o
}
func (o *ClimateOptions) MinTemp(temp float64) *ClimateOptions {
	o.minTemp = temp
	return o
}
func (o *ClimateOptions) ModeCommandTemplate(template string) *ClimateOptions {
	o.modeCommandTemplate = template
	return o
}
func (o *ClimateOptions) ModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.modeCommandFunc = f
	return o
}
func (o *ClimateOptions) ModeStateTemplate(template string) *ClimateOptions {
	o.modeStateTemplate = template
	return o
}
func (o *ClimateOptions) ModeStateFunc(f func() string) *ClimateOptions {
	o.modeStateFunc = f
	return o
}
func (o *ClimateOptions) EnableMode() *ClimateOptions {
	o.modeStateFunc = func() string {
		return o.states.Mode
	}
	o.modeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Mode = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) Modes(modes []string) *ClimateOptions {
	o.modes = modes
	return o
}
func (o *ClimateOptions) Name(name string) *ClimateOptions {
	o.name = name
	return o
}
func (o *ClimateOptions) ObjectId(id string) *ClimateOptions {
	o.objectId = id
	return o
}
func (o *ClimateOptions) Optimistic(optimistic bool) *ClimateOptions {
	o.optimistic = optimistic
	return o
}
func (o *ClimateOptions) PayloadAvailable(available string) *ClimateOptions {
	o.payloadAvailable = available
	return o
}
func (o *ClimateOptions) PayloadNotAvailable(available string) *ClimateOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *ClimateOptions) PayloadOff(off string) *ClimateOptions {
	o.payloadOff = off
	return o
}
func (o *ClimateOptions) PayloadOn(on string) *ClimateOptions {
	o.payloadOn = on
	return o
}
func (o *ClimateOptions) Precision(precision float64) *ClimateOptions {
	o.precision = precision
	return o
}
func (o *ClimateOptions) PresetModeCommandTemplate(template string) *ClimateOptions {
	o.presetModeCommandTemplate = template
	return o
}
func (o *ClimateOptions) PresetModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.presetModeCommandFunc = f
	return o
}
func (o *ClimateOptions) PresetModeStateFunc(f func() string) *ClimateOptions {
	o.presetModeStateFunc = f
	return o
}
func (o *ClimateOptions) EnablePresetMode() *ClimateOptions {
	o.presetModeStateFunc = func() string {
		return o.states.PresetMode
	}
	o.presetModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.PresetMode = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) PresetModeValueTemplate(template string) *ClimateOptions {
	o.presetModeValueTemplate = template
	return o
}
func (o *ClimateOptions) PresetModes(modes []string) *ClimateOptions {
	o.presetModes = modes
	return o
}
func (o *ClimateOptions) Qos(qos int) *ClimateOptions {
	o.qos = qos
	return o
}
func (o *ClimateOptions) Retain(retain bool) *ClimateOptions {
	o.retain = retain
	return o
}
func (o *ClimateOptions) SwingModeCommandTemplate(template string) *ClimateOptions {
	o.swingModeCommandTemplate = template
	return o
}
func (o *ClimateOptions) SwingModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.swingModeCommandFunc = f
	return o
}
func (o *ClimateOptions) SwingModeStateTemplate(template string) *ClimateOptions {
	o.swingModeStateTemplate = template
	return o
}
func (o *ClimateOptions) SwingModeStateFunc(f func() string) *ClimateOptions {
	o.swingModeStateFunc = f
	return o
}
func (o *ClimateOptions) EnableSwingMode() *ClimateOptions {
	o.swingModeStateFunc = func() string {
		return o.states.SwingMode
	}
	o.swingModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.SwingMode = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SwingModes(modes []string) *ClimateOptions {
	o.swingModes = modes
	return o
}
func (o *ClimateOptions) TargetHumidityCommandTemplate(template string) *ClimateOptions {
	o.targetHumidityCommandTemplate = template
	return o
}
func (o *ClimateOptions) TargetHumidityCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.targetHumidityCommandFunc = f
	return o
}
func (o *ClimateOptions) TargetHumidityStateTemplate(template string) *ClimateOptions {
	o.targetHumidityStateTemplate = template
	return o
}
func (o *ClimateOptions) TargetHumidityStateFunc(f func() string) *ClimateOptions {
	o.targetHumidityStateFunc = f
	return o
}
func (o *ClimateOptions) EnableTargetHumidity() *ClimateOptions {
	o.targetHumidityStateFunc = func() string {
		return o.states.TargetHumidity
	}
	o.targetHumidityCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.TargetHumidity = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) TempStep(step float64) *ClimateOptions {
	o.tempStep = step
	return o
}
func (o *ClimateOptions) TemperatureCommandTemplate(template string) *ClimateOptions {
	o.temperatureCommandTemplate = template
	return o
}
func (o *ClimateOptions) TemperatureCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.temperatureCommandFunc = f
	return o
}
func (o *ClimateOptions) TemperatureHighCommandTemplate(template string) *ClimateOptions {
	o.temperatureHighCommandTemplate = template
	return o
}
func (o *ClimateOptions) TemperatureHighCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.temperatureHighCommandFunc = f
	return o
}
func (o *ClimateOptions) TemperatureHighStateTemplate(template string) *ClimateOptions {
	o.temperatureHighStateTemplate = template
	return o
}
func (o *ClimateOptions) TemperatureHighStateFunc(f func() string) *ClimateOptions {
	o.temperatureHighStateFunc = f
	return o
}
func (o *ClimateOptions) EnableTemperatureHigh() *ClimateOptions {
	o.temperatureHighStateFunc = func() string {
		return o.states.TemperatureHigh
	}
	o.temperatureHighCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.TemperatureHigh = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) TemperatureLowCommandTemplate(template string) *ClimateOptions {
	o.temperatureLowCommandTemplate = template
	return o
}
func (o *ClimateOptions) TemperatureLowCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.temperatureLowCommandFunc = f
	return o
}
func (o *ClimateOptions) TemperatureLowStateTemplate(template string) *ClimateOptions {
	o.temperatureLowStateTemplate = template
	return o
}
func (o *ClimateOptions) TemperatureLowStateFunc(f func() string) *ClimateOptions {
	o.temperatureLowStateFunc = f
	return o
}
func (o *ClimateOptions) EnableTemperatureLow() *ClimateOptions {
	o.temperatureLowStateFunc = func() string {
		return o.states.TemperatureLow
	}
	o.temperatureLowCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.TemperatureLow = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) TemperatureStateTemplate(template string) *ClimateOptions {
	o.temperatureStateTemplate = template
	return o
}
func (o *ClimateOptions) TemperatureStateFunc(f func() string) *ClimateOptions {
	o.temperatureStateFunc = f
	return o
}
func (o *ClimateOptions) EnableTemperature() *ClimateOptions {
	o.temperatureStateFunc = func() string {
		return o.states.Temperature
	}
	o.temperatureCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Temperature = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) TemperatureUnit(unit string) *ClimateOptions {
	o.temperatureUnit = unit
	return o
}
func (o *ClimateOptions) UniqueId(id string) *ClimateOptions {
	o.uniqueId = id
	return o
}
func (o *ClimateOptions) ValueTemplate(template string) *ClimateOptions {
	o.valueTemplate = template
	return o
}
