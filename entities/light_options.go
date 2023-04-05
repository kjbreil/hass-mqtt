package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type LightOptions struct {
	States                    LightState // External state update location
	AvailabilityMode          string     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate      string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc          func() string
	BrightnessCommandTemplate string // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `brightness_command_topic`. Available variables: `value`."
	BrightnessCommandFunc     func(mqtt.Message, mqtt.Client)
	BrightnessScale           int // "Defines the maximum brightness value (i.e., 100%) of the MQTT device."
	BrightnessStateFunc       func() string
	BrightnessValueTemplate   string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the brightness value."
	ColorModeStateFunc        func() string
	ColorModeValueTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the color mode."
	ColorTempCommandTemplate  string // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `color_temp_command_topic`. Available variables: `value`."
	ColorTempCommandFunc      func(mqtt.Message, mqtt.Client)
	ColorTempStateFunc        func() string
	ColorTempValueTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the color temperature value."
	CommandFunc               func(mqtt.Message, mqtt.Client)
	EffectCommandTemplate     string // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `effect_command_topic`. Available variables: `value`."
	EffectCommandFunc         func(mqtt.Message, mqtt.Client)
	EffectList                ([]string) // "The list of effects the light supports."
	EffectStateFunc           func() string
	EffectValueTemplate       string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the effect value."
	EnabledByDefault          bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding                  string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory            string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	HsCommandTemplate         string // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `hs_command_topic`. Available variables: `hue` and `sat`."
	HsCommandFunc             func(mqtt.Message, mqtt.Client)
	HsStateFunc               func() string
	HsValueTemplate           string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the HS value."
	Icon                      string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc        func() string
	MaxMireds                 int    // "The maximum color temperature in mireds."
	MinMireds                 int    // "The minimum color temperature in mireds."
	Name                      string // "The name of the light."
	ObjectId                  string // "Used instead of `name` for automatic generation of `entity_id`"
	OnCommandType             string // "Defines when on the payload_on is sent. Using `last` (the default) will send any style (brightness, color, etc) topics first and then a `payload_on` to the `command_topic`. Using `first` will send the `payload_on` and then any style topics. Using `brightness` will only send brightness commands instead of the `payload_on` to turn the light on."
	Optimistic                bool   // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable          string // "The payload that represents the available state."
	PayloadNotAvailable       string // "The payload that represents the unavailable state."
	PayloadOff                string // "The payload that represents disabled state."
	PayloadOn                 string // "The payload that represents enabled state."
	Qos                       int    // "The maximum QoS level of the state topic."
	Retain                    bool   // "If the published message should have the retain flag on or not."
	RgbCommandTemplate        string // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgb_command_topic`. Available variables: `red`, `green` and `blue`."
	RgbCommandFunc            func(mqtt.Message, mqtt.Client)
	RgbStateFunc              func() string
	RgbValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGB value."
	RgbwCommandTemplate       string // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgbw_command_topic`. Available variables: `red`, `green`, `blue` and `white`."
	RgbwCommandFunc           func(mqtt.Message, mqtt.Client)
	RgbwStateFunc             func() string
	RgbwValueTemplate         string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGBW value."
	RgbwwCommandTemplate      string // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgbww_command_topic`. Available variables: `red`, `green`, `blue`, `cold_white` and `warm_white`."
	RgbwwCommandFunc          func(mqtt.Message, mqtt.Client)
	RgbwwStateFunc            func() string
	RgbwwValueTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGBWW value."
	Schema                    string // "The schema to use. Must be `default` or omitted to select the default schema."
	StateFunc                 func() string
	StateValueTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the state value. The template should match the payload `on` and `off` values, so if your light uses `power on` to turn on, your `state_value_template` string should return `power on` when the switch is on. For example if the message is just `on`, your `state_value_template` should be `power {{ value }}`."
	UniqueId                  string // "An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception."
	WhiteCommandFunc          func(mqtt.Message, mqtt.Client)
	WhiteScale                int    // "Defines the maximum white level (i.e., 100%) of the MQTT device."
	XyCommandTemplate         string // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `xy_command_topic`. Available variables: `x` and `y`."
	XyCommandFunc             func(mqtt.Message, mqtt.Client)
	XyStateFunc               func() string
	XyValueTemplate           string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the XY value."
}

func NewLightOptions() *LightOptions {
	return &LightOptions{}
}
func (o *LightOptions) GetStates() *LightState {
	return &o.States
}
func (o *LightOptions) SetAvailabilityMode(mode string) *LightOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *LightOptions) SetAvailabilityTemplate(template string) *LightOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *LightOptions) SetAvailabilityFunc(f func() string) *LightOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *LightOptions) SetBrightnessCommandTemplate(template string) *LightOptions {
	o.BrightnessCommandTemplate = template
	return o
}
func (o *LightOptions) SetBrightnessCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.BrightnessCommandFunc = f
	return o
}
func (o *LightOptions) SetBrightnessScale(scale int) *LightOptions {
	o.BrightnessScale = scale
	return o
}
func (o *LightOptions) SetBrightnessStateFunc(f func() string) *LightOptions {
	o.BrightnessStateFunc = f
	return o
}
func (o *LightOptions) HasBrightness() *LightOptions {
	o.BrightnessStateFunc = func() string {
		return o.States.Brightness
	}
	o.BrightnessCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Brightness = string(message.Payload())
	}
	return o
}
func (o *LightOptions) SetBrightnessValueTemplate(template string) *LightOptions {
	o.BrightnessValueTemplate = template
	return o
}
func (o *LightOptions) SetColorModeStateFunc(f func() string) *LightOptions {
	o.ColorModeStateFunc = f
	return o
}
func (o *LightOptions) SetColorModeValueTemplate(template string) *LightOptions {
	o.ColorModeValueTemplate = template
	return o
}
func (o *LightOptions) SetColorTempCommandTemplate(template string) *LightOptions {
	o.ColorTempCommandTemplate = template
	return o
}
func (o *LightOptions) SetColorTempCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.ColorTempCommandFunc = f
	return o
}
func (o *LightOptions) SetColorTempStateFunc(f func() string) *LightOptions {
	o.ColorTempStateFunc = f
	return o
}
func (o *LightOptions) HasColorTemp() *LightOptions {
	o.ColorTempStateFunc = func() string {
		return o.States.ColorTemp
	}
	o.ColorTempCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.ColorTemp = string(message.Payload())
	}
	return o
}
func (o *LightOptions) SetColorTempValueTemplate(template string) *LightOptions {
	o.ColorTempValueTemplate = template
	return o
}
func (o *LightOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.CommandFunc = f
	return o
}
func (o *LightOptions) SetEffectCommandTemplate(template string) *LightOptions {
	o.EffectCommandTemplate = template
	return o
}
func (o *LightOptions) SetEffectCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.EffectCommandFunc = f
	return o
}
func (o *LightOptions) SetEffectList(list []string) *LightOptions {
	o.EffectList = list
	return o
}
func (o *LightOptions) SetEffectStateFunc(f func() string) *LightOptions {
	o.EffectStateFunc = f
	return o
}
func (o *LightOptions) HasEffect() *LightOptions {
	o.EffectStateFunc = func() string {
		return o.States.Effect
	}
	o.EffectCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Effect = string(message.Payload())
	}
	return o
}
func (o *LightOptions) SetEffectValueTemplate(template string) *LightOptions {
	o.EffectValueTemplate = template
	return o
}
func (o *LightOptions) SetEnabledByDefault(d bool) *LightOptions {
	o.EnabledByDefault = d
	return o
}
func (o *LightOptions) SetEncoding(encoding string) *LightOptions {
	o.Encoding = encoding
	return o
}
func (o *LightOptions) SetEntityCategory(category string) *LightOptions {
	o.EntityCategory = category
	return o
}
func (o *LightOptions) SetHsCommandTemplate(template string) *LightOptions {
	o.HsCommandTemplate = template
	return o
}
func (o *LightOptions) SetHsCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.HsCommandFunc = f
	return o
}
func (o *LightOptions) SetHsStateFunc(f func() string) *LightOptions {
	o.HsStateFunc = f
	return o
}
func (o *LightOptions) HasHs() *LightOptions {
	o.HsStateFunc = func() string {
		return o.States.Hs
	}
	o.HsCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Hs = string(message.Payload())
	}
	return o
}
func (o *LightOptions) SetHsValueTemplate(template string) *LightOptions {
	o.HsValueTemplate = template
	return o
}
func (o *LightOptions) SetIcon(icon string) *LightOptions {
	o.Icon = icon
	return o
}
func (o *LightOptions) SetJsonAttributesTemplate(template string) *LightOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *LightOptions) SetJsonAttributesFunc(f func() string) *LightOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *LightOptions) SetMaxMireds(mireds int) *LightOptions {
	o.MaxMireds = mireds
	return o
}
func (o *LightOptions) SetMinMireds(mireds int) *LightOptions {
	o.MinMireds = mireds
	return o
}
func (o *LightOptions) SetName(name string) *LightOptions {
	o.Name = name
	return o
}
func (o *LightOptions) SetObjectId(id string) *LightOptions {
	o.ObjectId = id
	return o
}
func (o *LightOptions) SetOnCommandType(t string) *LightOptions {
	o.OnCommandType = t
	return o
}
func (o *LightOptions) SetOptimistic(optimistic bool) *LightOptions {
	o.Optimistic = optimistic
	return o
}
func (o *LightOptions) SetPayloadAvailable(available string) *LightOptions {
	o.PayloadAvailable = available
	return o
}
func (o *LightOptions) SetPayloadNotAvailable(available string) *LightOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *LightOptions) SetPayloadOff(off string) *LightOptions {
	o.PayloadOff = off
	return o
}
func (o *LightOptions) SetPayloadOn(on string) *LightOptions {
	o.PayloadOn = on
	return o
}
func (o *LightOptions) SetQos(qos int) *LightOptions {
	o.Qos = qos
	return o
}
func (o *LightOptions) SetRetain(retain bool) *LightOptions {
	o.Retain = retain
	return o
}
func (o *LightOptions) SetRgbCommandTemplate(template string) *LightOptions {
	o.RgbCommandTemplate = template
	return o
}
func (o *LightOptions) SetRgbCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.RgbCommandFunc = f
	return o
}
func (o *LightOptions) SetRgbStateFunc(f func() string) *LightOptions {
	o.RgbStateFunc = f
	return o
}
func (o *LightOptions) HasRgb() *LightOptions {
	o.RgbStateFunc = func() string {
		return o.States.Rgb
	}
	o.RgbCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Rgb = string(message.Payload())
	}
	return o
}
func (o *LightOptions) SetRgbValueTemplate(template string) *LightOptions {
	o.RgbValueTemplate = template
	return o
}
func (o *LightOptions) SetRgbwCommandTemplate(template string) *LightOptions {
	o.RgbwCommandTemplate = template
	return o
}
func (o *LightOptions) SetRgbwCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.RgbwCommandFunc = f
	return o
}
func (o *LightOptions) SetRgbwStateFunc(f func() string) *LightOptions {
	o.RgbwStateFunc = f
	return o
}
func (o *LightOptions) HasRgbw() *LightOptions {
	o.RgbwStateFunc = func() string {
		return o.States.Rgbw
	}
	o.RgbwCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Rgbw = string(message.Payload())
	}
	return o
}
func (o *LightOptions) SetRgbwValueTemplate(template string) *LightOptions {
	o.RgbwValueTemplate = template
	return o
}
func (o *LightOptions) SetRgbwwCommandTemplate(template string) *LightOptions {
	o.RgbwwCommandTemplate = template
	return o
}
func (o *LightOptions) SetRgbwwCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.RgbwwCommandFunc = f
	return o
}
func (o *LightOptions) SetRgbwwStateFunc(f func() string) *LightOptions {
	o.RgbwwStateFunc = f
	return o
}
func (o *LightOptions) HasRgbww() *LightOptions {
	o.RgbwwStateFunc = func() string {
		return o.States.Rgbww
	}
	o.RgbwwCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Rgbww = string(message.Payload())
	}
	return o
}
func (o *LightOptions) SetRgbwwValueTemplate(template string) *LightOptions {
	o.RgbwwValueTemplate = template
	return o
}
func (o *LightOptions) SetSchema(schema string) *LightOptions {
	o.Schema = schema
	return o
}
func (o *LightOptions) SetStateFunc(f func() string) *LightOptions {
	o.StateFunc = f
	return o
}
func (o *LightOptions) SetStateValueTemplate(template string) *LightOptions {
	o.StateValueTemplate = template
	return o
}
func (o *LightOptions) SetUniqueId(id string) *LightOptions {
	o.UniqueId = id
	return o
}
func (o *LightOptions) SetWhiteCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.WhiteCommandFunc = f
	return o
}
func (o *LightOptions) SetWhiteScale(scale int) *LightOptions {
	o.WhiteScale = scale
	return o
}
func (o *LightOptions) SetXyCommandTemplate(template string) *LightOptions {
	o.XyCommandTemplate = template
	return o
}
func (o *LightOptions) SetXyCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.XyCommandFunc = f
	return o
}
func (o *LightOptions) SetXyStateFunc(f func() string) *LightOptions {
	o.XyStateFunc = f
	return o
}
func (o *LightOptions) HasXy() *LightOptions {
	o.XyStateFunc = func() string {
		return o.States.Xy
	}
	o.XyCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.States.Xy = string(message.Payload())
	}
	return o
}
func (o *LightOptions) SetXyValueTemplate(template string) *LightOptions {
	o.XyValueTemplate = template
	return o
}
