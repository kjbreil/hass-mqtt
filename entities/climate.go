package entities

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	strcase "github.com/iancoleman/strcase"
	common "github.com/kjbreil/hass-mqtt/common"
	"log"
	"reflect"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Climate struct {
	ActionTemplate                 *string `json:"action_template,omitempty"` // "A template to render the value received on the `action_topic` with."
	ActionTopic                    *string `json:"action_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the current action. If this is set, the climate graph uses the value received as data source. Valid values: `off`, `heating`, `cooling`, `drying`, `idle`, `fan`."
	actionFunc                     func(mqtt.Message, mqtt.Client)
	AuxCommandTopic                *string `json:"aux_command_topic,omitempty"` // "The MQTT topic to publish commands to switch auxiliary heat."
	auxCommandFunc                 func(mqtt.Message, mqtt.Client)
	AuxStateTemplate               *string `json:"aux_state_template,omitempty"` // "A template to render the value received on the `aux_state_topic` with."
	AuxStateTopic                  *string `json:"aux_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the auxiliary heat mode. If this is not set, the auxiliary heat mode works in optimistic mode (see below)."
	auxStateFunc                   func() string
	AvailabilityMode               *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate           *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic              *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc               func() string
	CurrentHumidityTemplate        *string `json:"current_humidity_template,omitempty"` // "A template with which the value received on `current_humidity_topic` will be rendered."
	CurrentHumidityTopic           *string `json:"current_humidity_topic,omitempty"`    // "The MQTT topic on which to listen for the current humidity. A `\"None\"` value received will reset the current temperature. Empty values (`'''`) will be ignored."
	currentHumidityFunc            func() string
	CurrentTemperatureTemplate     *string `json:"current_temperature_template,omitempty"` // "A template with which the value received on `current_temperature_topic` will be rendered."
	CurrentTemperatureTopic        *string `json:"current_temperature_topic,omitempty"`    // "The MQTT topic on which to listen for the current temperature. A `\"None\"` value received will reset the current humidity. Empty values (`'''`) will be ignored."
	currentTemperatureFunc         func() string
	Device                         Device  `json:"device,omitempty"`                    // Device configuration parameters
	EnabledByDefault               *bool   `json:"enabled_by_default,omitempty"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding                       *string `json:"encoding,omitempty"`                  // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                 *string `json:"entity_category,omitempty"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	FanModeCommandTemplate         *string `json:"fan_mode_command_template,omitempty"` // "A template to render the value sent to the `fan_mode_command_topic` with."
	FanModeCommandTopic            *string `json:"fan_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the fan mode."
	fanModeCommandFunc             func(mqtt.Message, mqtt.Client)
	FanModeStateTemplate           *string `json:"fan_mode_state_template,omitempty"` // "A template to render the value received on the `fan_mode_state_topic` with."
	FanModeStateTopic              *string `json:"fan_mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC fan mode. If this is not set, the fan mode works in optimistic mode (see below)."
	fanModeStateFunc               func() string
	FanModes                       *([]string) `json:"fan_modes,omitempty"`                // "A list of supported fan modes."
	Icon                           *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Initial                        *int        `json:"initial,omitempty"`                  // "Set the initial target temperature."
	JsonAttributesTemplate         *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic            *string     `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc             func() string
	MaxHumidity                    *int     `json:"max_humidity,omitempty"`          // "The minimum target humidity percentage that can be set."
	MaxTemp                        *float64 `json:"max_temp,omitempty"`              // "Maximum set point available."
	MinHumidity                    *int     `json:"min_humidity,omitempty"`          // "The maximum target humidity percentage that can be set."
	MinTemp                        *float64 `json:"min_temp,omitempty"`              // "Minimum set point available."
	ModeCommandTemplate            *string  `json:"mode_command_template,omitempty"` // "A template to render the value sent to the `mode_command_topic` with."
	ModeCommandTopic               *string  `json:"mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the HVAC operation mode. Use with `mode_command_template` if you only want to publish the power state."
	modeCommandFunc                func(mqtt.Message, mqtt.Client)
	ModeStateTemplate              *string `json:"mode_state_template,omitempty"` // "A template to render the value received on the `mode_state_topic` with."
	ModeStateTopic                 *string `json:"mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC operation mode. If this is not set, the operation mode works in optimistic mode (see below)."
	modeStateFunc                  func() string
	Modes                          *([]string) `json:"modes,omitempty"`                        // "A list of supported modes. Needs to be a subset of the default values."
	Name                           *string     `json:"name,omitempty"`                         // "The name of the HVAC."
	ObjectId                       *string     `json:"object_id,omitempty"`                    // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                     *bool       `json:"optimistic,omitempty"`                   // "Flag that defines if the climate works in optimistic mode"
	PayloadAvailable               *string     `json:"payload_available,omitempty"`            // "The payload that represents the available state."
	PayloadNotAvailable            *string     `json:"payload_not_available,omitempty"`        // "The payload that represents the unavailable state."
	PayloadOff                     *string     `json:"payload_off,omitempty"`                  // "The payload that represents disabled state."
	PayloadOn                      *string     `json:"payload_on,omitempty"`                   // "The payload that represents enabled state."
	Precision                      *float64    `json:"precision,omitempty"`                    // "The desired precision for this device. Can be used to match your actual thermostat's precision. Supported values are `0.1`, `0.5` and `1.0`."
	PresetModeCommandTemplate      *string     `json:"preset_mode_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommandTopic         *string     `json:"preset_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the preset mode."
	presetModeCommandFunc          func(mqtt.Message, mqtt.Client)
	PresetModeStateTopic           *string `json:"preset_mode_state_topic,omitempty"` // "The MQTT topic subscribed to receive climate speed based on presets. When preset 'none' is received or `None` the `preset_mode` will be reset."
	presetModeStateFunc            func() string
	PresetModeValueTemplate        *string     `json:"preset_mode_value_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                    *([]string) `json:"preset_modes,omitempty"`                // "List of preset modes this climate is supporting. Common examples include `eco`, `away`, `boost`, `comfort`, `home`, `sleep` and `activity`."
	Qos                            *int        `json:"qos,omitempty"`                         // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                         *bool       `json:"retain,omitempty"`                      // "Defines if published messages should have the retain flag set."
	SwingModeCommandTemplate       *string     `json:"swing_mode_command_template,omitempty"` // "A template to render the value sent to the `swing_mode_command_topic` with."
	SwingModeCommandTopic          *string     `json:"swing_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the swing mode."
	swingModeCommandFunc           func(mqtt.Message, mqtt.Client)
	SwingModeStateTemplate         *string `json:"swing_mode_state_template,omitempty"` // "A template to render the value received on the `swing_mode_state_topic` with."
	SwingModeStateTopic            *string `json:"swing_mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC swing mode. If this is not set, the swing mode works in optimistic mode (see below)."
	swingModeStateFunc             func() string
	SwingModes                     *([]string) `json:"swing_modes,omitempty"`                      // "A list of supported swing modes."
	TargetHumidityCommandTemplate  *string     `json:"target_humidity_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `target_humidity_command_topic`."
	TargetHumidityCommandTopic     *string     `json:"target_humidity_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the target humidity."
	targetHumidityCommandFunc      func(mqtt.Message, mqtt.Client)
	TargetHumidityStateTemplate    *string `json:"target_humidity_state_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the climate `target_humidity` state."
	TargetHumidityStateTopic       *string `json:"target_humidity_state_topic,omitempty"`    // "The MQTT topic subscribed to receive the target humidity. If this is not set, the target humidity works in optimistic mode (see below). A `\"None\"` value received will reset the target humidity. Empty values (`'''`) will be ignored."
	targetHumidityStateFunc        func() string
	TempStep                       *float64 `json:"temp_step,omitempty"`                    // "Step size for temperature set point."
	TemperatureCommandTemplate     *string  `json:"temperature_command_template,omitempty"` // "A template to render the value sent to the `temperature_command_topic` with."
	TemperatureCommandTopic        *string  `json:"temperature_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the target temperature."
	temperatureCommandFunc         func(mqtt.Message, mqtt.Client)
	TemperatureHighCommandTemplate *string `json:"temperature_high_command_template,omitempty"` // "A template to render the value sent to the `temperature_high_command_topic` with."
	TemperatureHighCommandTopic    *string `json:"temperature_high_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the high target temperature."
	temperatureHighCommandFunc     func(mqtt.Message, mqtt.Client)
	TemperatureHighStateTemplate   *string `json:"temperature_high_state_template,omitempty"` // "A template to render the value received on the `temperature_high_state_topic` with. A `\"None\"` value received will reset the temperature high set point. Empty values (`'''`) will be ignored."
	TemperatureHighStateTopic      *string `json:"temperature_high_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes in the target high temperature. If this is not set, the target high temperature works in optimistic mode (see below)."
	temperatureHighStateFunc       func() string
	TemperatureLowCommandTemplate  *string `json:"temperature_low_command_template,omitempty"` // "A template to render the value sent to the `temperature_low_command_topic` with."
	TemperatureLowCommandTopic     *string `json:"temperature_low_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the target low temperature."
	temperatureLowCommandFunc      func(mqtt.Message, mqtt.Client)
	TemperatureLowStateTemplate    *string `json:"temperature_low_state_template,omitempty"` // "A template to render the value received on the `temperature_low_state_topic` with. A `\"None\"` value received will reset the temperature low set point. Empty values (`'''`) will be ignored."
	TemperatureLowStateTopic       *string `json:"temperature_low_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes in the target low temperature. If this is not set, the target low temperature works in optimistic mode (see below)."
	temperatureLowStateFunc        func() string
	TemperatureStateTemplate       *string `json:"temperature_state_template,omitempty"` // "A template to render the value received on the `temperature_state_topic` with."
	TemperatureStateTopic          *string `json:"temperature_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes in the target temperature. If this is not set, the target temperature works in optimistic mode (see below). A `\"None\"` value received will reset the temperature set point. Empty values (`'''`) will be ignored."
	temperatureStateFunc           func() string
	TemperatureUnit                *string       `json:"temperature_unit,omitempty"` // "Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit."
	UniqueId                       *string       `json:"unique_id,omitempty"`        // "An ID that uniquely identifies this HVAC device. If two HVAC devices have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate                  *string       `json:"value_template,omitempty"`   // "Default template to render the payloads on *all* `*_state_topic`s with."
	MQTT                           *MQTTFields   `json:"-"`                          // MQTT configuration parameters
	states                         climateState  // Internal Holder of States
	States                         *ClimateState `json:"-"` // External state update location
}

func NewClimate(o *ClimateOptions) *Climate {
	var c Climate

	c.States = &o.States
	if !reflect.ValueOf(o.ActionTemplate).IsZero() {
		c.ActionTemplate = &o.ActionTemplate
	}
	if !reflect.ValueOf(o.ActionFunc).IsZero() {
		c.actionFunc = o.ActionFunc
	}
	if !reflect.ValueOf(o.AuxCommandFunc).IsZero() {
		c.auxCommandFunc = o.AuxCommandFunc
	}
	if !reflect.ValueOf(o.AuxStateTemplate).IsZero() {
		c.AuxStateTemplate = &o.AuxStateTemplate
	}
	if !reflect.ValueOf(o.AuxStateFunc).IsZero() {
		c.auxStateFunc = o.AuxStateFunc
	}
	if !reflect.ValueOf(o.AvailabilityMode).IsZero() {
		c.AvailabilityMode = &o.AvailabilityMode
	}
	if !reflect.ValueOf(o.AvailabilityTemplate).IsZero() {
		c.AvailabilityTemplate = &o.AvailabilityTemplate
	}
	if !reflect.ValueOf(o.AvailabilityFunc).IsZero() {
		c.availabilityFunc = o.AvailabilityFunc
	}
	if !reflect.ValueOf(o.CurrentHumidityTemplate).IsZero() {
		c.CurrentHumidityTemplate = &o.CurrentHumidityTemplate
	}
	if !reflect.ValueOf(o.CurrentHumidityFunc).IsZero() {
		c.currentHumidityFunc = o.CurrentHumidityFunc
	}
	if !reflect.ValueOf(o.CurrentTemperatureTemplate).IsZero() {
		c.CurrentTemperatureTemplate = &o.CurrentTemperatureTemplate
	}
	if !reflect.ValueOf(o.CurrentTemperatureFunc).IsZero() {
		c.currentTemperatureFunc = o.CurrentTemperatureFunc
	}
	if !reflect.ValueOf(o.EnabledByDefault).IsZero() {
		c.EnabledByDefault = &o.EnabledByDefault
	}
	if !reflect.ValueOf(o.Encoding).IsZero() {
		c.Encoding = &o.Encoding
	}
	if !reflect.ValueOf(o.EntityCategory).IsZero() {
		c.EntityCategory = &o.EntityCategory
	}
	if !reflect.ValueOf(o.FanModeCommandTemplate).IsZero() {
		c.FanModeCommandTemplate = &o.FanModeCommandTemplate
	}
	if !reflect.ValueOf(o.FanModeCommandFunc).IsZero() {
		c.fanModeCommandFunc = o.FanModeCommandFunc
	}
	if !reflect.ValueOf(o.FanModeStateTemplate).IsZero() {
		c.FanModeStateTemplate = &o.FanModeStateTemplate
	}
	if !reflect.ValueOf(o.FanModeStateFunc).IsZero() {
		c.fanModeStateFunc = o.FanModeStateFunc
	}
	if !reflect.ValueOf(o.FanModes).IsZero() {
		c.FanModes = &o.FanModes
	}
	if !reflect.ValueOf(o.Icon).IsZero() {
		c.Icon = &o.Icon
	}
	if !reflect.ValueOf(o.Initial).IsZero() {
		c.Initial = &o.Initial
	}
	if !reflect.ValueOf(o.JsonAttributesTemplate).IsZero() {
		c.JsonAttributesTemplate = &o.JsonAttributesTemplate
	}
	if !reflect.ValueOf(o.JsonAttributesFunc).IsZero() {
		c.jsonAttributesFunc = o.JsonAttributesFunc
	}
	if !reflect.ValueOf(o.MaxHumidity).IsZero() {
		c.MaxHumidity = &o.MaxHumidity
	}
	if !reflect.ValueOf(o.MaxTemp).IsZero() {
		c.MaxTemp = &o.MaxTemp
	}
	if !reflect.ValueOf(o.MinHumidity).IsZero() {
		c.MinHumidity = &o.MinHumidity
	}
	if !reflect.ValueOf(o.MinTemp).IsZero() {
		c.MinTemp = &o.MinTemp
	}
	if !reflect.ValueOf(o.ModeCommandTemplate).IsZero() {
		c.ModeCommandTemplate = &o.ModeCommandTemplate
	}
	if !reflect.ValueOf(o.ModeCommandFunc).IsZero() {
		c.modeCommandFunc = o.ModeCommandFunc
	}
	if !reflect.ValueOf(o.ModeStateTemplate).IsZero() {
		c.ModeStateTemplate = &o.ModeStateTemplate
	}
	if !reflect.ValueOf(o.ModeStateFunc).IsZero() {
		c.modeStateFunc = o.ModeStateFunc
	}
	if !reflect.ValueOf(o.Modes).IsZero() {
		c.Modes = &o.Modes
	}
	if !reflect.ValueOf(o.Name).IsZero() {
		c.Name = &o.Name
	}
	if !reflect.ValueOf(o.ObjectId).IsZero() {
		c.ObjectId = &o.ObjectId
	}
	if !reflect.ValueOf(o.Optimistic).IsZero() {
		c.Optimistic = &o.Optimistic
	}
	if !reflect.ValueOf(o.PayloadAvailable).IsZero() {
		c.PayloadAvailable = &o.PayloadAvailable
	}
	if !reflect.ValueOf(o.PayloadNotAvailable).IsZero() {
		c.PayloadNotAvailable = &o.PayloadNotAvailable
	}
	if !reflect.ValueOf(o.PayloadOff).IsZero() {
		c.PayloadOff = &o.PayloadOff
	}
	if !reflect.ValueOf(o.PayloadOn).IsZero() {
		c.PayloadOn = &o.PayloadOn
	}
	if !reflect.ValueOf(o.Precision).IsZero() {
		c.Precision = &o.Precision
	}
	if !reflect.ValueOf(o.PresetModeCommandTemplate).IsZero() {
		c.PresetModeCommandTemplate = &o.PresetModeCommandTemplate
	}
	if !reflect.ValueOf(o.PresetModeCommandFunc).IsZero() {
		c.presetModeCommandFunc = o.PresetModeCommandFunc
	}
	if !reflect.ValueOf(o.PresetModeStateFunc).IsZero() {
		c.presetModeStateFunc = o.PresetModeStateFunc
	}
	if !reflect.ValueOf(o.PresetModeValueTemplate).IsZero() {
		c.PresetModeValueTemplate = &o.PresetModeValueTemplate
	}
	if !reflect.ValueOf(o.PresetModes).IsZero() {
		c.PresetModes = &o.PresetModes
	}
	if !reflect.ValueOf(o.Qos).IsZero() {
		c.Qos = &o.Qos
	}
	if !reflect.ValueOf(o.Retain).IsZero() {
		c.Retain = &o.Retain
	}
	if !reflect.ValueOf(o.SwingModeCommandTemplate).IsZero() {
		c.SwingModeCommandTemplate = &o.SwingModeCommandTemplate
	}
	if !reflect.ValueOf(o.SwingModeCommandFunc).IsZero() {
		c.swingModeCommandFunc = o.SwingModeCommandFunc
	}
	if !reflect.ValueOf(o.SwingModeStateTemplate).IsZero() {
		c.SwingModeStateTemplate = &o.SwingModeStateTemplate
	}
	if !reflect.ValueOf(o.SwingModeStateFunc).IsZero() {
		c.swingModeStateFunc = o.SwingModeStateFunc
	}
	if !reflect.ValueOf(o.SwingModes).IsZero() {
		c.SwingModes = &o.SwingModes
	}
	if !reflect.ValueOf(o.TargetHumidityCommandTemplate).IsZero() {
		c.TargetHumidityCommandTemplate = &o.TargetHumidityCommandTemplate
	}
	if !reflect.ValueOf(o.TargetHumidityCommandFunc).IsZero() {
		c.targetHumidityCommandFunc = o.TargetHumidityCommandFunc
	}
	if !reflect.ValueOf(o.TargetHumidityStateTemplate).IsZero() {
		c.TargetHumidityStateTemplate = &o.TargetHumidityStateTemplate
	}
	if !reflect.ValueOf(o.TargetHumidityStateFunc).IsZero() {
		c.targetHumidityStateFunc = o.TargetHumidityStateFunc
	}
	if !reflect.ValueOf(o.TempStep).IsZero() {
		c.TempStep = &o.TempStep
	}
	if !reflect.ValueOf(o.TemperatureCommandTemplate).IsZero() {
		c.TemperatureCommandTemplate = &o.TemperatureCommandTemplate
	}
	if !reflect.ValueOf(o.TemperatureCommandFunc).IsZero() {
		c.temperatureCommandFunc = o.TemperatureCommandFunc
	}
	if !reflect.ValueOf(o.TemperatureHighCommandTemplate).IsZero() {
		c.TemperatureHighCommandTemplate = &o.TemperatureHighCommandTemplate
	}
	if !reflect.ValueOf(o.TemperatureHighCommandFunc).IsZero() {
		c.temperatureHighCommandFunc = o.TemperatureHighCommandFunc
	}
	if !reflect.ValueOf(o.TemperatureHighStateTemplate).IsZero() {
		c.TemperatureHighStateTemplate = &o.TemperatureHighStateTemplate
	}
	if !reflect.ValueOf(o.TemperatureHighStateFunc).IsZero() {
		c.temperatureHighStateFunc = o.TemperatureHighStateFunc
	}
	if !reflect.ValueOf(o.TemperatureLowCommandTemplate).IsZero() {
		c.TemperatureLowCommandTemplate = &o.TemperatureLowCommandTemplate
	}
	if !reflect.ValueOf(o.TemperatureLowCommandFunc).IsZero() {
		c.temperatureLowCommandFunc = o.TemperatureLowCommandFunc
	}
	if !reflect.ValueOf(o.TemperatureLowStateTemplate).IsZero() {
		c.TemperatureLowStateTemplate = &o.TemperatureLowStateTemplate
	}
	if !reflect.ValueOf(o.TemperatureLowStateFunc).IsZero() {
		c.temperatureLowStateFunc = o.TemperatureLowStateFunc
	}
	if !reflect.ValueOf(o.TemperatureStateTemplate).IsZero() {
		c.TemperatureStateTemplate = &o.TemperatureStateTemplate
	}
	if !reflect.ValueOf(o.TemperatureStateFunc).IsZero() {
		c.temperatureStateFunc = o.TemperatureStateFunc
	}
	if !reflect.ValueOf(o.TemperatureUnit).IsZero() {
		c.TemperatureUnit = &o.TemperatureUnit
	}
	if !reflect.ValueOf(o.UniqueId).IsZero() {
		c.UniqueId = &o.UniqueId
	}
	if !reflect.ValueOf(o.ValueTemplate).IsZero() {
		c.ValueTemplate = &o.ValueTemplate
	}
	return &c
}

type climateState struct {
	Aux                *string
	Availability       *string
	CurrentHumidity    *string
	CurrentTemperature *string
	FanMode            *string
	JsonAttributes     *string
	Mode               *string
	PresetMode         *string
	SwingMode          *string
	TargetHumidity     *string
	TemperatureHigh    *string
	TemperatureLow     *string
	Temperature        *string
}
type ClimateState struct {
	Aux                string
	CurrentHumidity    string
	CurrentTemperature string
	FanMode            string
	JsonAttributes     string
	Mode               string
	PresetMode         string
	SwingMode          string
	TargetHumidity     string
	TemperatureHigh    string
	TemperatureLow     string
	Temperature        string
}

func (d *Climate) SetAux(s string) {
	d.States.Aux = s
	d.UpdateState()
}
func (d *Climate) SetCurrentHumidity(s string) {
	d.States.CurrentHumidity = s
	d.UpdateState()
}
func (d *Climate) SetCurrentTemperature(s string) {
	d.States.CurrentTemperature = s
	d.UpdateState()
}
func (d *Climate) SetFanMode(s string) {
	d.States.FanMode = s
	d.UpdateState()
}
func (d *Climate) SetJsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Climate) SetMode(s string) {
	d.States.Mode = s
	d.UpdateState()
}
func (d *Climate) SetPresetMode(s string) {
	d.States.PresetMode = s
	d.UpdateState()
}
func (d *Climate) SetSwingMode(s string) {
	d.States.SwingMode = s
	d.UpdateState()
}
func (d *Climate) SetTargetHumidity(s string) {
	d.States.TargetHumidity = s
	d.UpdateState()
}
func (d *Climate) SetTemperatureHigh(s string) {
	d.States.TemperatureHigh = s
	d.UpdateState()
}
func (d *Climate) SetTemperatureLow(s string) {
	d.States.TemperatureLow = s
	d.UpdateState()
}
func (d *Climate) SetTemperature(s string) {
	d.States.Temperature = s
	d.UpdateState()
}
func (d *Climate) GetRawId() string {
	return "climate"
}
func (d *Climate) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Climate) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Climate) GetName() string {
	return *d.Name
}
func (d *Climate) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Climate) UpdateState() {
	if d.AuxStateTopic != nil {
		state := d.auxStateFunc()
		if d.states.Aux == nil || state != *d.states.Aux || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AuxStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Aux = &state
		}
	}
	if d.AvailabilityTopic != nil {
		state := d.availabilityFunc()
		if d.states.Availability == nil || state != *d.states.Availability || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
		}
	}
	if d.CurrentHumidityTopic != nil {
		state := d.currentHumidityFunc()
		if d.states.CurrentHumidity == nil || state != *d.states.CurrentHumidity || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.CurrentHumidityTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.CurrentHumidity = &state
		}
	}
	if d.CurrentTemperatureTopic != nil {
		state := d.currentTemperatureFunc()
		if d.states.CurrentTemperature == nil || state != *d.states.CurrentTemperature || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.CurrentTemperatureTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.CurrentTemperature = &state
		}
	}
	if d.FanModeStateTopic != nil {
		state := d.fanModeStateFunc()
		if d.states.FanMode == nil || state != *d.states.FanMode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.FanModeStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.FanMode = &state
		}
	}
	if d.JsonAttributesTopic != nil {
		state := d.jsonAttributesFunc()
		if d.states.JsonAttributes == nil || state != *d.states.JsonAttributes || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.JsonAttributesTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.JsonAttributes = &state
		}
	}
	if d.ModeStateTopic != nil {
		state := d.modeStateFunc()
		if d.states.Mode == nil || state != *d.states.Mode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ModeStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Mode = &state
		}
	}
	if d.PresetModeStateTopic != nil {
		state := d.presetModeStateFunc()
		if d.states.PresetMode == nil || state != *d.states.PresetMode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.PresetModeStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.PresetMode = &state
		}
	}
	if d.SwingModeStateTopic != nil {
		state := d.swingModeStateFunc()
		if d.states.SwingMode == nil || state != *d.states.SwingMode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.SwingModeStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.SwingMode = &state
		}
	}
	if d.TargetHumidityStateTopic != nil {
		state := d.targetHumidityStateFunc()
		if d.states.TargetHumidity == nil || state != *d.states.TargetHumidity || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TargetHumidityStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.TargetHumidity = &state
		}
	}
	if d.TemperatureHighStateTopic != nil {
		state := d.temperatureHighStateFunc()
		if d.states.TemperatureHigh == nil || state != *d.states.TemperatureHigh || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TemperatureHighStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.TemperatureHigh = &state
		}
	}
	if d.TemperatureLowStateTopic != nil {
		state := d.temperatureLowStateFunc()
		if d.states.TemperatureLow == nil || state != *d.states.TemperatureLow || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TemperatureLowStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.TemperatureLow = &state
		}
	}
	if d.TemperatureStateTopic != nil {
		state := d.temperatureStateFunc()
		if d.states.Temperature == nil || state != *d.states.Temperature || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TemperatureStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Temperature = &state
		}
	}
}
func (d *Climate) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.ActionTopic != nil {
		t := c.Subscribe(*d.ActionTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.AuxCommandTopic != nil {
		t := c.Subscribe(*d.AuxCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.FanModeCommandTopic != nil {
		t := c.Subscribe(*d.FanModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ModeCommandTopic != nil {
		t := c.Subscribe(*d.ModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PresetModeCommandTopic != nil {
		t := c.Subscribe(*d.PresetModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SwingModeCommandTopic != nil {
		t := c.Subscribe(*d.SwingModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TargetHumidityCommandTopic != nil {
		t := c.Subscribe(*d.TargetHumidityCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureCommandTopic != nil {
		t := c.Subscribe(*d.TemperatureCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureHighCommandTopic != nil {
		t := c.Subscribe(*d.TemperatureHighCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureLowCommandTopic != nil {
		t := c.Subscribe(*d.TemperatureLowCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	token := c.Publish(GetDiscoveryTopic(d), 2, true, message)
	token.WaitTimeout(common.WaitTimeout)
	d.availabilityFunc()
	d.UpdateState()
}
func (d *Climate) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, false, "offline")
	token.WaitTimeout(common.WaitTimeout)
	if d.ActionTopic != nil {
		t := c.Unsubscribe(*d.ActionTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.AuxCommandTopic != nil {
		t := c.Unsubscribe(*d.AuxCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.FanModeCommandTopic != nil {
		t := c.Unsubscribe(*d.FanModeCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ModeCommandTopic != nil {
		t := c.Unsubscribe(*d.ModeCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PresetModeCommandTopic != nil {
		t := c.Unsubscribe(*d.PresetModeCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SwingModeCommandTopic != nil {
		t := c.Unsubscribe(*d.SwingModeCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TargetHumidityCommandTopic != nil {
		t := c.Unsubscribe(*d.TargetHumidityCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureCommandTopic != nil {
		t := c.Unsubscribe(*d.TemperatureCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureHighCommandTopic != nil {
		t := c.Unsubscribe(*d.TemperatureHighCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureLowCommandTopic != nil {
		t := c.Unsubscribe(*d.TemperatureLowCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Climate) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Climate) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(0)
	}
	if d.Retain == nil {
		d.Retain = new(bool)
		*d.Retain = false
	}
	if d.UniqueId == nil {
		d.UniqueId = new(string)
		*d.UniqueId = strcase.ToDelimited(*d.Name, uint8(0x2d))
	}
	if d.availabilityFunc == nil {
		d.availabilityFunc = AvailabilityFunc
	}
	if d.MQTT == nil {
		d.MQTT = new(MQTTFields)
	}
	d.AddMessageHandler()
	d.populateTopics()
}
func (d *Climate) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Climate) populateTopics() {
	if d.actionFunc != nil {
		d.ActionTopic = new(string)
		*d.ActionTopic = GetTopic(d, "action_topic")
		common.TopicStore[*d.ActionTopic] = &d.actionFunc
	}
	if d.auxCommandFunc != nil {
		d.AuxCommandTopic = new(string)
		*d.AuxCommandTopic = GetTopic(d, "aux_command_topic")
		common.TopicStore[*d.AuxCommandTopic] = &d.auxCommandFunc
	}
	if d.auxStateFunc != nil {
		d.AuxStateTopic = new(string)
		*d.AuxStateTopic = GetTopic(d, "aux_state_topic")
	}
	if d.availabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.currentHumidityFunc != nil {
		d.CurrentHumidityTopic = new(string)
		*d.CurrentHumidityTopic = GetTopic(d, "current_humidity_topic")
	}
	if d.currentTemperatureFunc != nil {
		d.CurrentTemperatureTopic = new(string)
		*d.CurrentTemperatureTopic = GetTopic(d, "current_temperature_topic")
	}
	if d.fanModeCommandFunc != nil {
		d.FanModeCommandTopic = new(string)
		*d.FanModeCommandTopic = GetTopic(d, "fan_mode_command_topic")
		common.TopicStore[*d.FanModeCommandTopic] = &d.fanModeCommandFunc
	}
	if d.fanModeStateFunc != nil {
		d.FanModeStateTopic = new(string)
		*d.FanModeStateTopic = GetTopic(d, "fan_mode_state_topic")
	}
	if d.jsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
	}
	if d.modeCommandFunc != nil {
		d.ModeCommandTopic = new(string)
		*d.ModeCommandTopic = GetTopic(d, "mode_command_topic")
		common.TopicStore[*d.ModeCommandTopic] = &d.modeCommandFunc
	}
	if d.modeStateFunc != nil {
		d.ModeStateTopic = new(string)
		*d.ModeStateTopic = GetTopic(d, "mode_state_topic")
	}
	if d.presetModeCommandFunc != nil {
		d.PresetModeCommandTopic = new(string)
		*d.PresetModeCommandTopic = GetTopic(d, "preset_mode_command_topic")
		common.TopicStore[*d.PresetModeCommandTopic] = &d.presetModeCommandFunc
	}
	if d.presetModeStateFunc != nil {
		d.PresetModeStateTopic = new(string)
		*d.PresetModeStateTopic = GetTopic(d, "preset_mode_state_topic")
	}
	if d.swingModeCommandFunc != nil {
		d.SwingModeCommandTopic = new(string)
		*d.SwingModeCommandTopic = GetTopic(d, "swing_mode_command_topic")
		common.TopicStore[*d.SwingModeCommandTopic] = &d.swingModeCommandFunc
	}
	if d.swingModeStateFunc != nil {
		d.SwingModeStateTopic = new(string)
		*d.SwingModeStateTopic = GetTopic(d, "swing_mode_state_topic")
	}
	if d.targetHumidityCommandFunc != nil {
		d.TargetHumidityCommandTopic = new(string)
		*d.TargetHumidityCommandTopic = GetTopic(d, "target_humidity_command_topic")
		common.TopicStore[*d.TargetHumidityCommandTopic] = &d.targetHumidityCommandFunc
	}
	if d.targetHumidityStateFunc != nil {
		d.TargetHumidityStateTopic = new(string)
		*d.TargetHumidityStateTopic = GetTopic(d, "target_humidity_state_topic")
	}
	if d.temperatureCommandFunc != nil {
		d.TemperatureCommandTopic = new(string)
		*d.TemperatureCommandTopic = GetTopic(d, "temperature_command_topic")
		common.TopicStore[*d.TemperatureCommandTopic] = &d.temperatureCommandFunc
	}
	if d.temperatureHighCommandFunc != nil {
		d.TemperatureHighCommandTopic = new(string)
		*d.TemperatureHighCommandTopic = GetTopic(d, "temperature_high_command_topic")
		common.TopicStore[*d.TemperatureHighCommandTopic] = &d.temperatureHighCommandFunc
	}
	if d.temperatureHighStateFunc != nil {
		d.TemperatureHighStateTopic = new(string)
		*d.TemperatureHighStateTopic = GetTopic(d, "temperature_high_state_topic")
	}
	if d.temperatureLowCommandFunc != nil {
		d.TemperatureLowCommandTopic = new(string)
		*d.TemperatureLowCommandTopic = GetTopic(d, "temperature_low_command_topic")
		common.TopicStore[*d.TemperatureLowCommandTopic] = &d.temperatureLowCommandFunc
	}
	if d.temperatureLowStateFunc != nil {
		d.TemperatureLowStateTopic = new(string)
		*d.TemperatureLowStateTopic = GetTopic(d, "temperature_low_state_topic")
	}
	if d.temperatureStateFunc != nil {
		d.TemperatureStateTopic = new(string)
		*d.TemperatureStateTopic = GetTopic(d, "temperature_state_topic")
	}
}
func (d *Climate) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Climate) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
