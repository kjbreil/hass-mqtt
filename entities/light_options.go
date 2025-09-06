package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type LightOptions struct {
	states                    LightState // External state update location
	availabilityMode          string     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate      string     // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc          func() string
	brightnessCommandTemplate string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to compose message which will be sent to `brightness_command_topic`. Available variables: `value`."
	brightnessCommandFunc     func(mqtt.Message, mqtt.Client)
	brightnessScale           int // "Defines the maximum brightness value (i.e., 100%) of the MQTT device."
	brightnessStateFunc       func() string
	brightnessValueTemplate   string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the brightness value."
	colorModeStateFunc        func() string
	colorModeValueTemplate    string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the color mode."
	colorTempCommandTemplate  string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to compose message which will be sent to `color_temp_command_topic`. Available variables: `value`."
	colorTempCommandFunc      func(mqtt.Message, mqtt.Client)
	colorTempKelvin           bool // "When set to `true`, `color_temp_command_topic` will publish color mode updates in Kelvin and process `color_temp_state_topic` will process state updates in Kelvin. When not set the `color_temp` values are converted to mireds."
	colorTempStateFunc        func() string
	colorTempValueTemplate    string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the color temperature value."
	commandFunc               func(mqtt.Message, mqtt.Client)
	effectCommandTemplate     string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to compose message which will be sent to `effect_command_topic`. Available variables: `value`."
	effectCommandFunc         func(mqtt.Message, mqtt.Client)
	effectList                ([]string) // "The list of effects the light supports."
	effectStateFunc           func() string
	effectValueTemplate       string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the effect value."
	enabledByDefault          bool   // "Flag which defines if the entity should be enabled when first added."
	encoding                  string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory            string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	entityPicture             string // "Picture URL for the entity."
	hsCommandTemplate         string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to compose message which will be sent to `hs_command_topic`. Available variables: `hue` and `sat`."
	hsCommandFunc             func(mqtt.Message, mqtt.Client)
	hsStateFunc               func() string
	hsValueTemplate           string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the HS value."
	icon                      string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate    string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc        func() string
	maxKelvin                 int    // "The maximum color temperature in Kelvin."
	maxMireds                 int    // "The maximum color temperature in mireds."
	minKelvin                 int    // "The minimum color temperature in Kelvin."
	minMireds                 int    // "The minimum color temperature in mireds."
	name                      string // "The name of the light. Can be set to `null` if only the device name is relevant."
	objectId                  string // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	onCommandType             string // "Defines when on the payload_on is sent. Using `last` (the default) will send any style (brightness, color, etc) topics first and then a `payload_on` to the `command_topic`. Using `first` will send the `payload_on` and then any style topics. Using `brightness` will only send brightness commands instead of the `payload_on` to turn the light on."
	optimistic                bool   // "Flag that defines if switch works in optimistic mode."
	payloadAvailable          string // "The payload that represents the available state."
	payloadNotAvailable       string // "The payload that represents the unavailable state."
	payloadOff                string // "The payload that represents the off state."
	payloadOn                 string // "The payload that represents the on state."
	platform                  string // "Must be `light`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                       int    // "The maximum QoS level to be used when receiving and publishing messages."
	retain                    bool   // "If the published message should have the retain flag on or not."
	rgbCommandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to compose message which will be sent to `rgb_command_topic`. Available variables: `red`, `green` and `blue`."
	rgbCommandFunc            func(mqtt.Message, mqtt.Client)
	rgbStateFunc              func() string
	rgbValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the RGB value."
	rgbwCommandTemplate       string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to compose message which will be sent to `rgbw_command_topic`. Available variables: `red`, `green`, `blue` and `white`."
	rgbwCommandFunc           func(mqtt.Message, mqtt.Client)
	rgbwStateFunc             func() string
	rgbwValueTemplate         string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the RGBW value."
	rgbwwCommandTemplate      string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to compose message which will be sent to `rgbww_command_topic`. Available variables: `red`, `green`, `blue`, `cold_white` and `warm_white`."
	rgbwwCommandFunc          func(mqtt.Message, mqtt.Client)
	rgbwwStateFunc            func() string
	rgbwwValueTemplate        string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the RGBWW value."
	schema                    string // "The schema to use. Must be `basic` or omitted to select the default schema."
	stateFunc                 func() string
	stateValueTemplate        string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the state value. The template should return the values defined by `payload_on` (defaults to \"ON\") and `payload_off` (defaults to \"OFF\") settings, or \"None\"."
	uniqueId                  string // "An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	whiteCommandFunc          func(mqtt.Message, mqtt.Client)
	whiteScale                int    // "Defines the maximum white level (i.e., 100%) of the MQTT device."
	xyCommandTemplate         string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to compose message which will be sent to `xy_command_topic`. Available variables: `x` and `y`."
	xyCommandFunc             func(mqtt.Message, mqtt.Client)
	xyStateFunc               func() string
	xyValueTemplate           string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the XY value."
}

func NewLightOptions() *LightOptions {
	return &LightOptions{}
}
func (o *LightOptions) States() *LightState {
	return &o.states
}
func (o *LightOptions) AvailabilityMode(mode string) *LightOptions {
	o.availabilityMode = mode
	return o
}
func (o *LightOptions) AvailabilityTemplate(template string) *LightOptions {
	o.availabilityTemplate = template
	return o
}
func (o *LightOptions) AvailabilityFunc(f func() string) *LightOptions {
	o.availabilityFunc = f
	return o
}
func (o *LightOptions) BrightnessCommandTemplate(template string) *LightOptions {
	o.brightnessCommandTemplate = template
	return o
}
func (o *LightOptions) BrightnessCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.brightnessCommandFunc = f
	return o
}
func (o *LightOptions) BrightnessScale(scale int) *LightOptions {
	o.brightnessScale = scale
	return o
}
func (o *LightOptions) BrightnessStateFunc(f func() string) *LightOptions {
	o.brightnessStateFunc = f
	return o
}
func (o *LightOptions) EnableBrightness() *LightOptions {
	o.brightnessStateFunc = func() string {
		return o.states.Brightness
	}
	o.brightnessCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Brightness = string(message.Payload())
	}
	return o
}
func (o *LightOptions) BrightnessValueTemplate(template string) *LightOptions {
	o.brightnessValueTemplate = template
	return o
}
func (o *LightOptions) ColorModeStateFunc(f func() string) *LightOptions {
	o.colorModeStateFunc = f
	return o
}
func (o *LightOptions) ColorModeValueTemplate(template string) *LightOptions {
	o.colorModeValueTemplate = template
	return o
}
func (o *LightOptions) ColorTempCommandTemplate(template string) *LightOptions {
	o.colorTempCommandTemplate = template
	return o
}
func (o *LightOptions) ColorTempCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.colorTempCommandFunc = f
	return o
}
func (o *LightOptions) ColorTempKelvin(kelvin bool) *LightOptions {
	o.colorTempKelvin = kelvin
	return o
}
func (o *LightOptions) ColorTempStateFunc(f func() string) *LightOptions {
	o.colorTempStateFunc = f
	return o
}
func (o *LightOptions) EnableColorTemp() *LightOptions {
	o.colorTempStateFunc = func() string {
		return o.states.ColorTemp
	}
	o.colorTempCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.ColorTemp = string(message.Payload())
	}
	return o
}
func (o *LightOptions) ColorTempValueTemplate(template string) *LightOptions {
	o.colorTempValueTemplate = template
	return o
}
func (o *LightOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.commandFunc = f
	return o
}
func (o *LightOptions) EffectCommandTemplate(template string) *LightOptions {
	o.effectCommandTemplate = template
	return o
}
func (o *LightOptions) EffectCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.effectCommandFunc = f
	return o
}
func (o *LightOptions) EffectList(list []string) *LightOptions {
	o.effectList = list
	return o
}
func (o *LightOptions) EffectStateFunc(f func() string) *LightOptions {
	o.effectStateFunc = f
	return o
}
func (o *LightOptions) EnableEffect() *LightOptions {
	o.effectStateFunc = func() string {
		return o.states.Effect
	}
	o.effectCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Effect = string(message.Payload())
	}
	return o
}
func (o *LightOptions) EffectValueTemplate(template string) *LightOptions {
	o.effectValueTemplate = template
	return o
}
func (o *LightOptions) EnabledByDefault(d bool) *LightOptions {
	o.enabledByDefault = d
	return o
}
func (o *LightOptions) Encoding(encoding string) *LightOptions {
	o.encoding = encoding
	return o
}
func (o *LightOptions) EntityCategory(category string) *LightOptions {
	o.entityCategory = category
	return o
}
func (o *LightOptions) EntityPicture(picture string) *LightOptions {
	o.entityPicture = picture
	return o
}
func (o *LightOptions) HsCommandTemplate(template string) *LightOptions {
	o.hsCommandTemplate = template
	return o
}
func (o *LightOptions) HsCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.hsCommandFunc = f
	return o
}
func (o *LightOptions) HsStateFunc(f func() string) *LightOptions {
	o.hsStateFunc = f
	return o
}
func (o *LightOptions) EnableHs() *LightOptions {
	o.hsStateFunc = func() string {
		return o.states.Hs
	}
	o.hsCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Hs = string(message.Payload())
	}
	return o
}
func (o *LightOptions) HsValueTemplate(template string) *LightOptions {
	o.hsValueTemplate = template
	return o
}
func (o *LightOptions) Icon(icon string) *LightOptions {
	o.icon = icon
	return o
}
func (o *LightOptions) JsonAttributesTemplate(template string) *LightOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *LightOptions) JsonAttributesFunc(f func() string) *LightOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *LightOptions) MaxKelvin(kelvin int) *LightOptions {
	o.maxKelvin = kelvin
	return o
}
func (o *LightOptions) MaxMireds(mireds int) *LightOptions {
	o.maxMireds = mireds
	return o
}
func (o *LightOptions) MinKelvin(kelvin int) *LightOptions {
	o.minKelvin = kelvin
	return o
}
func (o *LightOptions) MinMireds(mireds int) *LightOptions {
	o.minMireds = mireds
	return o
}
func (o *LightOptions) Name(name string) *LightOptions {
	o.name = name
	return o
}
func (o *LightOptions) ObjectId(id string) *LightOptions {
	o.objectId = id
	return o
}
func (o *LightOptions) OnCommandType(t string) *LightOptions {
	o.onCommandType = t
	return o
}
func (o *LightOptions) Optimistic(optimistic bool) *LightOptions {
	o.optimistic = optimistic
	return o
}
func (o *LightOptions) PayloadAvailable(available string) *LightOptions {
	o.payloadAvailable = available
	return o
}
func (o *LightOptions) PayloadNotAvailable(available string) *LightOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *LightOptions) PayloadOff(off string) *LightOptions {
	o.payloadOff = off
	return o
}
func (o *LightOptions) PayloadOn(on string) *LightOptions {
	o.payloadOn = on
	return o
}
func (o *LightOptions) Platform(platform string) *LightOptions {
	o.platform = platform
	return o
}
func (o *LightOptions) Qos(qos int) *LightOptions {
	o.qos = qos
	return o
}
func (o *LightOptions) Retain(retain bool) *LightOptions {
	o.retain = retain
	return o
}
func (o *LightOptions) RgbCommandTemplate(template string) *LightOptions {
	o.rgbCommandTemplate = template
	return o
}
func (o *LightOptions) RgbCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.rgbCommandFunc = f
	return o
}
func (o *LightOptions) RgbStateFunc(f func() string) *LightOptions {
	o.rgbStateFunc = f
	return o
}
func (o *LightOptions) EnableRgb() *LightOptions {
	o.rgbStateFunc = func() string {
		return o.states.Rgb
	}
	o.rgbCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Rgb = string(message.Payload())
	}
	return o
}
func (o *LightOptions) RgbValueTemplate(template string) *LightOptions {
	o.rgbValueTemplate = template
	return o
}
func (o *LightOptions) RgbwCommandTemplate(template string) *LightOptions {
	o.rgbwCommandTemplate = template
	return o
}
func (o *LightOptions) RgbwCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.rgbwCommandFunc = f
	return o
}
func (o *LightOptions) RgbwStateFunc(f func() string) *LightOptions {
	o.rgbwStateFunc = f
	return o
}
func (o *LightOptions) EnableRgbw() *LightOptions {
	o.rgbwStateFunc = func() string {
		return o.states.Rgbw
	}
	o.rgbwCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Rgbw = string(message.Payload())
	}
	return o
}
func (o *LightOptions) RgbwValueTemplate(template string) *LightOptions {
	o.rgbwValueTemplate = template
	return o
}
func (o *LightOptions) RgbwwCommandTemplate(template string) *LightOptions {
	o.rgbwwCommandTemplate = template
	return o
}
func (o *LightOptions) RgbwwCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.rgbwwCommandFunc = f
	return o
}
func (o *LightOptions) RgbwwStateFunc(f func() string) *LightOptions {
	o.rgbwwStateFunc = f
	return o
}
func (o *LightOptions) EnableRgbww() *LightOptions {
	o.rgbwwStateFunc = func() string {
		return o.states.Rgbww
	}
	o.rgbwwCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Rgbww = string(message.Payload())
	}
	return o
}
func (o *LightOptions) RgbwwValueTemplate(template string) *LightOptions {
	o.rgbwwValueTemplate = template
	return o
}
func (o *LightOptions) Schema(schema string) *LightOptions {
	o.schema = schema
	return o
}
func (o *LightOptions) StateFunc(f func() string) *LightOptions {
	o.stateFunc = f
	return o
}
func (o *LightOptions) StateValueTemplate(template string) *LightOptions {
	o.stateValueTemplate = template
	return o
}
func (o *LightOptions) UniqueId(id string) *LightOptions {
	o.uniqueId = id
	return o
}
func (o *LightOptions) WhiteCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.whiteCommandFunc = f
	return o
}
func (o *LightOptions) WhiteScale(scale int) *LightOptions {
	o.whiteScale = scale
	return o
}
func (o *LightOptions) XyCommandTemplate(template string) *LightOptions {
	o.xyCommandTemplate = template
	return o
}
func (o *LightOptions) XyCommandFunc(f func(mqtt.Message, mqtt.Client)) *LightOptions {
	o.xyCommandFunc = f
	return o
}
func (o *LightOptions) XyStateFunc(f func() string) *LightOptions {
	o.xyStateFunc = f
	return o
}
func (o *LightOptions) EnableXy() *LightOptions {
	o.xyStateFunc = func() string {
		return o.states.Xy
	}
	o.xyCommandFunc = func(message mqtt.Message, client mqtt.Client) {
		o.states.Xy = string(message.Payload())
	}
	return o
}
func (o *LightOptions) XyValueTemplate(template string) *LightOptions {
	o.xyValueTemplate = template
	return o
}
