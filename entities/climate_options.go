package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type ClimateOptions struct {
	States                         ClimateState // External state update location
	ActionTemplate                 string       // "A template to render the value received on the `action_topic` with."
	ActionFunc                     func(mqtt.Message, mqtt.Client)
	AuxCommandFunc                 func(mqtt.Message, mqtt.Client)
	AuxStateTemplate               string // "A template to render the value received on the `aux_state_topic` with."
	AuxStateFunc                   func() string
	AvailabilityMode               string // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate           string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc               func() string
	CurrentHumidityTemplate        string // "A template with which the value received on `current_humidity_topic` will be rendered."
	CurrentHumidityFunc            func() string
	CurrentTemperatureTemplate     string // "A template with which the value received on `current_temperature_topic` will be rendered."
	CurrentTemperatureFunc         func() string
	EnabledByDefault               bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding                       string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                 string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	FanModeCommandTemplate         string // "A template to render the value sent to the `fan_mode_command_topic` with."
	FanModeCommandFunc             func(mqtt.Message, mqtt.Client)
	FanModeStateTemplate           string // "A template to render the value received on the `fan_mode_state_topic` with."
	FanModeStateFunc               func() string
	FanModes                       ([]string) // "A list of supported fan modes."
	Icon                           string     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Initial                        int        // "Set the initial target temperature."
	JsonAttributesTemplate         string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc             func() string
	MaxHumidity                    int     // "The minimum target humidity percentage that can be set."
	MaxTemp                        float64 // "Maximum set point available."
	MinHumidity                    int     // "The maximum target humidity percentage that can be set."
	MinTemp                        float64 // "Minimum set point available."
	ModeCommandTemplate            string  // "A template to render the value sent to the `mode_command_topic` with."
	ModeCommandFunc                func(mqtt.Message, mqtt.Client)
	ModeStateTemplate              string // "A template to render the value received on the `mode_state_topic` with."
	ModeStateFunc                  func() string
	Modes                          ([]string) // "A list of supported modes. Needs to be a subset of the default values."
	Name                           string     // "The name of the HVAC."
	ObjectId                       string     // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                     bool       // "Flag that defines if the climate works in optimistic mode"
	PayloadAvailable               string     // "The payload that represents the available state."
	PayloadNotAvailable            string     // "The payload that represents the unavailable state."
	PayloadOff                     string     // "The payload that represents disabled state."
	PayloadOn                      string     // "The payload that represents enabled state."
	Precision                      float64    // "The desired precision for this device. Can be used to match your actual thermostat's precision. Supported values are `0.1`, `0.5` and `1.0`."
	PresetModeCommandTemplate      string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommandFunc          func(mqtt.Message, mqtt.Client)
	PresetModeStateFunc            func() string
	PresetModeValueTemplate        string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                    ([]string) // "List of preset modes this climate is supporting. Common examples include `eco`, `away`, `boost`, `comfort`, `home`, `sleep` and `activity`."
	Qos                            int        // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                         bool       // "Defines if published messages should have the retain flag set."
	SwingModeCommandTemplate       string     // "A template to render the value sent to the `swing_mode_command_topic` with."
	SwingModeCommandFunc           func(mqtt.Message, mqtt.Client)
	SwingModeStateTemplate         string // "A template to render the value received on the `swing_mode_state_topic` with."
	SwingModeStateFunc             func() string
	SwingModes                     ([]string) // "A list of supported swing modes."
	TargetHumidityCommandTemplate  string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `target_humidity_command_topic`."
	TargetHumidityCommandFunc      func(mqtt.Message, mqtt.Client)
	TargetHumidityStateTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the climate `target_humidity` state."
	TargetHumidityStateFunc        func() string
	TempStep                       float64 // "Step size for temperature set point."
	TemperatureCommandTemplate     string  // "A template to render the value sent to the `temperature_command_topic` with."
	TemperatureCommandFunc         func(mqtt.Message, mqtt.Client)
	TemperatureHighCommandTemplate string // "A template to render the value sent to the `temperature_high_command_topic` with."
	TemperatureHighCommandFunc     func(mqtt.Message, mqtt.Client)
	TemperatureHighStateTemplate   string // "A template to render the value received on the `temperature_high_state_topic` with. A `\"None\"` value received will reset the temperature high set point. Empty values (`'''`) will be ignored."
	TemperatureHighStateFunc       func() string
	TemperatureLowCommandTemplate  string // "A template to render the value sent to the `temperature_low_command_topic` with."
	TemperatureLowCommandFunc      func(mqtt.Message, mqtt.Client)
	TemperatureLowStateTemplate    string // "A template to render the value received on the `temperature_low_state_topic` with. A `\"None\"` value received will reset the temperature low set point. Empty values (`'''`) will be ignored."
	TemperatureLowStateFunc        func() string
	TemperatureStateTemplate       string // "A template to render the value received on the `temperature_state_topic` with."
	TemperatureStateFunc           func() string
	TemperatureUnit                string // "Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit."
	UniqueId                       string // "An ID that uniquely identifies this HVAC device. If two HVAC devices have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate                  string // "Default template to render the payloads on *all* `*_state_topic`s with."
}

func NewClimateOptions() *ClimateOptions {
	return &ClimateOptions{}
}
func (o *ClimateOptions) GetStates() *ClimateState {
	return &o.States
}
func (o *ClimateOptions) SetActionTemplate(template string) *ClimateOptions {
	o.ActionTemplate = template
	return o
}
func (o *ClimateOptions) SetActionFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.ActionFunc = f
	return o
}
func (o *ClimateOptions) SetAuxCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.AuxCommandFunc = f
	return o
}
func (o *ClimateOptions) SetAuxStateTemplate(template string) *ClimateOptions {
	o.AuxStateTemplate = template
	return o
}
func (o *ClimateOptions) SetAuxStateFunc(f func() string) *ClimateOptions {
	o.AuxStateFunc = f
	return o
}
func (o *ClimateOptions) HasAux() *ClimateOptions {
	o.AuxStateFunc = func() string {
		return o.States.Aux
	}
	o.AuxCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Aux = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SetAvailabilityMode(mode string) *ClimateOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *ClimateOptions) SetAvailabilityTemplate(template string) *ClimateOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *ClimateOptions) SetAvailabilityFunc(f func() string) *ClimateOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *ClimateOptions) SetCurrentHumidityTemplate(template string) *ClimateOptions {
	o.CurrentHumidityTemplate = template
	return o
}
func (o *ClimateOptions) SetCurrentHumidityFunc(f func() string) *ClimateOptions {
	o.CurrentHumidityFunc = f
	return o
}
func (o *ClimateOptions) SetCurrentTemperatureTemplate(template string) *ClimateOptions {
	o.CurrentTemperatureTemplate = template
	return o
}
func (o *ClimateOptions) SetCurrentTemperatureFunc(f func() string) *ClimateOptions {
	o.CurrentTemperatureFunc = f
	return o
}
func (o *ClimateOptions) SetEnabledByDefault(d bool) *ClimateOptions {
	o.EnabledByDefault = d
	return o
}
func (o *ClimateOptions) SetEncoding(encoding string) *ClimateOptions {
	o.Encoding = encoding
	return o
}
func (o *ClimateOptions) SetEntityCategory(category string) *ClimateOptions {
	o.EntityCategory = category
	return o
}
func (o *ClimateOptions) SetFanModeCommandTemplate(template string) *ClimateOptions {
	o.FanModeCommandTemplate = template
	return o
}
func (o *ClimateOptions) SetFanModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.FanModeCommandFunc = f
	return o
}
func (o *ClimateOptions) SetFanModeStateTemplate(template string) *ClimateOptions {
	o.FanModeStateTemplate = template
	return o
}
func (o *ClimateOptions) SetFanModeStateFunc(f func() string) *ClimateOptions {
	o.FanModeStateFunc = f
	return o
}
func (o *ClimateOptions) HasFanMode() *ClimateOptions {
	o.FanModeStateFunc = func() string {
		return o.States.FanMode
	}
	o.FanModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.FanMode = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SetFanModes(modes []string) *ClimateOptions {
	o.FanModes = modes
	return o
}
func (o *ClimateOptions) SetIcon(icon string) *ClimateOptions {
	o.Icon = icon
	return o
}
func (o *ClimateOptions) SetInitial(initial int) *ClimateOptions {
	o.Initial = initial
	return o
}
func (o *ClimateOptions) SetJsonAttributesTemplate(template string) *ClimateOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *ClimateOptions) SetJsonAttributesFunc(f func() string) *ClimateOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *ClimateOptions) SetMaxHumidity(humidity int) *ClimateOptions {
	o.MaxHumidity = humidity
	return o
}
func (o *ClimateOptions) SetMaxTemp(temp float64) *ClimateOptions {
	o.MaxTemp = temp
	return o
}
func (o *ClimateOptions) SetMinHumidity(humidity int) *ClimateOptions {
	o.MinHumidity = humidity
	return o
}
func (o *ClimateOptions) SetMinTemp(temp float64) *ClimateOptions {
	o.MinTemp = temp
	return o
}
func (o *ClimateOptions) SetModeCommandTemplate(template string) *ClimateOptions {
	o.ModeCommandTemplate = template
	return o
}
func (o *ClimateOptions) SetModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.ModeCommandFunc = f
	return o
}
func (o *ClimateOptions) SetModeStateTemplate(template string) *ClimateOptions {
	o.ModeStateTemplate = template
	return o
}
func (o *ClimateOptions) SetModeStateFunc(f func() string) *ClimateOptions {
	o.ModeStateFunc = f
	return o
}
func (o *ClimateOptions) HasMode() *ClimateOptions {
	o.ModeStateFunc = func() string {
		return o.States.Mode
	}
	o.ModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Mode = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SetModes(modes []string) *ClimateOptions {
	o.Modes = modes
	return o
}
func (o *ClimateOptions) SetName(name string) *ClimateOptions {
	o.Name = name
	return o
}
func (o *ClimateOptions) SetObjectId(id string) *ClimateOptions {
	o.ObjectId = id
	return o
}
func (o *ClimateOptions) SetOptimistic(optimistic bool) *ClimateOptions {
	o.Optimistic = optimistic
	return o
}
func (o *ClimateOptions) SetPayloadAvailable(available string) *ClimateOptions {
	o.PayloadAvailable = available
	return o
}
func (o *ClimateOptions) SetPayloadNotAvailable(available string) *ClimateOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *ClimateOptions) SetPayloadOff(off string) *ClimateOptions {
	o.PayloadOff = off
	return o
}
func (o *ClimateOptions) SetPayloadOn(on string) *ClimateOptions {
	o.PayloadOn = on
	return o
}
func (o *ClimateOptions) SetPrecision(precision float64) *ClimateOptions {
	o.Precision = precision
	return o
}
func (o *ClimateOptions) SetPresetModeCommandTemplate(template string) *ClimateOptions {
	o.PresetModeCommandTemplate = template
	return o
}
func (o *ClimateOptions) SetPresetModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.PresetModeCommandFunc = f
	return o
}
func (o *ClimateOptions) SetPresetModeStateFunc(f func() string) *ClimateOptions {
	o.PresetModeStateFunc = f
	return o
}
func (o *ClimateOptions) HasPresetMode() *ClimateOptions {
	o.PresetModeStateFunc = func() string {
		return o.States.PresetMode
	}
	o.PresetModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.PresetMode = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SetPresetModeValueTemplate(template string) *ClimateOptions {
	o.PresetModeValueTemplate = template
	return o
}
func (o *ClimateOptions) SetPresetModes(modes []string) *ClimateOptions {
	o.PresetModes = modes
	return o
}
func (o *ClimateOptions) SetQos(qos int) *ClimateOptions {
	o.Qos = qos
	return o
}
func (o *ClimateOptions) SetRetain(retain bool) *ClimateOptions {
	o.Retain = retain
	return o
}
func (o *ClimateOptions) SetSwingModeCommandTemplate(template string) *ClimateOptions {
	o.SwingModeCommandTemplate = template
	return o
}
func (o *ClimateOptions) SetSwingModeCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.SwingModeCommandFunc = f
	return o
}
func (o *ClimateOptions) SetSwingModeStateTemplate(template string) *ClimateOptions {
	o.SwingModeStateTemplate = template
	return o
}
func (o *ClimateOptions) SetSwingModeStateFunc(f func() string) *ClimateOptions {
	o.SwingModeStateFunc = f
	return o
}
func (o *ClimateOptions) HasSwingMode() *ClimateOptions {
	o.SwingModeStateFunc = func() string {
		return o.States.SwingMode
	}
	o.SwingModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.SwingMode = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SetSwingModes(modes []string) *ClimateOptions {
	o.SwingModes = modes
	return o
}
func (o *ClimateOptions) SetTargetHumidityCommandTemplate(template string) *ClimateOptions {
	o.TargetHumidityCommandTemplate = template
	return o
}
func (o *ClimateOptions) SetTargetHumidityCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.TargetHumidityCommandFunc = f
	return o
}
func (o *ClimateOptions) SetTargetHumidityStateTemplate(template string) *ClimateOptions {
	o.TargetHumidityStateTemplate = template
	return o
}
func (o *ClimateOptions) SetTargetHumidityStateFunc(f func() string) *ClimateOptions {
	o.TargetHumidityStateFunc = f
	return o
}
func (o *ClimateOptions) HasTargetHumidity() *ClimateOptions {
	o.TargetHumidityStateFunc = func() string {
		return o.States.TargetHumidity
	}
	o.TargetHumidityCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.TargetHumidity = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SetTempStep(step float64) *ClimateOptions {
	o.TempStep = step
	return o
}
func (o *ClimateOptions) SetTemperatureCommandTemplate(template string) *ClimateOptions {
	o.TemperatureCommandTemplate = template
	return o
}
func (o *ClimateOptions) SetTemperatureCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.TemperatureCommandFunc = f
	return o
}
func (o *ClimateOptions) SetTemperatureHighCommandTemplate(template string) *ClimateOptions {
	o.TemperatureHighCommandTemplate = template
	return o
}
func (o *ClimateOptions) SetTemperatureHighCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.TemperatureHighCommandFunc = f
	return o
}
func (o *ClimateOptions) SetTemperatureHighStateTemplate(template string) *ClimateOptions {
	o.TemperatureHighStateTemplate = template
	return o
}
func (o *ClimateOptions) SetTemperatureHighStateFunc(f func() string) *ClimateOptions {
	o.TemperatureHighStateFunc = f
	return o
}
func (o *ClimateOptions) HasTemperatureHigh() *ClimateOptions {
	o.TemperatureHighStateFunc = func() string {
		return o.States.TemperatureHigh
	}
	o.TemperatureHighCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.TemperatureHigh = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SetTemperatureLowCommandTemplate(template string) *ClimateOptions {
	o.TemperatureLowCommandTemplate = template
	return o
}
func (o *ClimateOptions) SetTemperatureLowCommandFunc(f func(mqtt.Message, mqtt.Client)) *ClimateOptions {
	o.TemperatureLowCommandFunc = f
	return o
}
func (o *ClimateOptions) SetTemperatureLowStateTemplate(template string) *ClimateOptions {
	o.TemperatureLowStateTemplate = template
	return o
}
func (o *ClimateOptions) SetTemperatureLowStateFunc(f func() string) *ClimateOptions {
	o.TemperatureLowStateFunc = f
	return o
}
func (o *ClimateOptions) HasTemperatureLow() *ClimateOptions {
	o.TemperatureLowStateFunc = func() string {
		return o.States.TemperatureLow
	}
	o.TemperatureLowCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.TemperatureLow = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SetTemperatureStateTemplate(template string) *ClimateOptions {
	o.TemperatureStateTemplate = template
	return o
}
func (o *ClimateOptions) SetTemperatureStateFunc(f func() string) *ClimateOptions {
	o.TemperatureStateFunc = f
	return o
}
func (o *ClimateOptions) HasTemperature() *ClimateOptions {
	o.TemperatureStateFunc = func() string {
		return o.States.Temperature
	}
	o.TemperatureCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Temperature = string(message.Payload())
	}
	return o
}
func (o *ClimateOptions) SetTemperatureUnit(unit string) *ClimateOptions {
	o.TemperatureUnit = unit
	return o
}
func (o *ClimateOptions) SetUniqueId(id string) *ClimateOptions {
	o.UniqueId = id
	return o
}
func (o *ClimateOptions) SetValueTemplate(template string) *ClimateOptions {
	o.ValueTemplate = template
	return o
}
