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
type Light struct {
	AvailabilityMode          *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate      *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic         *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc          func() string
	BrightnessCommandTemplate *string `json:"brightness_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `brightness_command_topic`. Available variables: `value`."
	BrightnessCommandTopic    *string `json:"brightness_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light’s brightness."
	brightnessCommandFunc     func(mqtt.Message, mqtt.Client)
	BrightnessScale           *int    `json:"brightness_scale,omitempty"`       // "Defines the maximum brightness value (i.e., 100%) of the MQTT device."
	BrightnessStateTopic      *string `json:"brightness_state_topic,omitempty"` // "The MQTT topic subscribed to receive brightness state updates."
	brightnessStateFunc       func() string
	BrightnessValueTemplate   *string `json:"brightness_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the brightness value."
	ColorModeStateTopic       *string `json:"color_mode_state_topic,omitempty"`    // "The MQTT topic subscribed to receive color mode updates. If this is not configured, `color_mode` will be automatically set according to the last received valid color or color temperature"
	colorModeStateFunc        func() string
	ColorModeValueTemplate    *string `json:"color_mode_value_template,omitempty"`   // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the color mode."
	ColorTempCommandTemplate  *string `json:"color_temp_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `color_temp_command_topic`. Available variables: `value`."
	ColorTempCommandTopic     *string `json:"color_temp_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light’s color temperature state. The color temperature command slider has a range of 153 to 500 mireds (micro reciprocal degrees)."
	colorTempCommandFunc      func(mqtt.Message, mqtt.Client)
	ColorTempStateTopic       *string `json:"color_temp_state_topic,omitempty"` // "The MQTT topic subscribed to receive color temperature state updates."
	colorTempStateFunc        func() string
	ColorTempValueTemplate    *string `json:"color_temp_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the color temperature value."
	CommandTopic              *string `json:"command_topic,omitempty"`             // "The MQTT topic to publish commands to change the switch state."
	commandFunc               func(mqtt.Message, mqtt.Client)
	Device                    Device  `json:"device,omitempty"`                  // Device configuration parameters
	EffectCommandTemplate     *string `json:"effect_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `effect_command_topic`. Available variables: `value`."
	EffectCommandTopic        *string `json:"effect_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light's effect state."
	effectCommandFunc         func(mqtt.Message, mqtt.Client)
	EffectList                *([]string) `json:"effect_list,omitempty"`        // "The list of effects the light supports."
	EffectStateTopic          *string     `json:"effect_state_topic,omitempty"` // "The MQTT topic subscribed to receive effect state updates."
	effectStateFunc           func() string
	EffectValueTemplate       *string `json:"effect_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the effect value."
	EnabledByDefault          *bool   `json:"enabled_by_default,omitempty"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding                  *string `json:"encoding,omitempty"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory            *string `json:"entity_category,omitempty"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	HsCommandTemplate         *string `json:"hs_command_template,omitempty"`   // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `hs_command_topic`. Available variables: `hue` and `sat`."
	HsCommandTopic            *string `json:"hs_command_topic,omitempty"`      // "The MQTT topic to publish commands to change the light's color state in HS format (Hue Saturation). Range for Hue: 0° .. 360°, Range of Saturation: 0..100. Note: Brightness is sent separately in the `brightness_command_topic`."
	hsCommandFunc             func(mqtt.Message, mqtt.Client)
	HsStateTopic              *string `json:"hs_state_topic,omitempty"` // "The MQTT topic subscribed to receive color state updates in HS format. The expected payload is the hue and saturation values separated by commas, for example, `359.5,100.0`. Note: Brightness is received separately in the `brightness_state_topic`."
	hsStateFunc               func() string
	HsValueTemplate           *string `json:"hs_value_template,omitempty"`        // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the HS value."
	Icon                      *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate    *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic       *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc        func() string
	MaxMireds                 *int    `json:"max_mireds,omitempty"`            // "The maximum color temperature in mireds."
	MinMireds                 *int    `json:"min_mireds,omitempty"`            // "The minimum color temperature in mireds."
	Name                      *string `json:"name,omitempty"`                  // "The name of the light."
	ObjectId                  *string `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	OnCommandType             *string `json:"on_command_type,omitempty"`       // "Defines when on the payload_on is sent. Using `last` (the default) will send any style (brightness, color, etc) topics first and then a `payload_on` to the `command_topic`. Using `first` will send the `payload_on` and then any style topics. Using `brightness` will only send brightness commands instead of the `payload_on` to turn the light on."
	Optimistic                *bool   `json:"optimistic,omitempty"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable          *string `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable       *string `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOff                *string `json:"payload_off,omitempty"`           // "The payload that represents disabled state."
	PayloadOn                 *string `json:"payload_on,omitempty"`            // "The payload that represents enabled state."
	Qos                       *int    `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic."
	Retain                    *bool   `json:"retain,omitempty"`                // "If the published message should have the retain flag on or not."
	RgbCommandTemplate        *string `json:"rgb_command_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgb_command_topic`. Available variables: `red`, `green` and `blue`."
	RgbCommandTopic           *string `json:"rgb_command_topic,omitempty"`     // "The MQTT topic to publish commands to change the light's RGB state."
	rgbCommandFunc            func(mqtt.Message, mqtt.Client)
	RgbStateTopic             *string `json:"rgb_state_topic,omitempty"` // "The MQTT topic subscribed to receive RGB state updates. The expected payload is the RGB values separated by commas, for example, `255,0,127`."
	rgbStateFunc              func() string
	RgbValueTemplate          *string `json:"rgb_value_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGB value."
	RgbwCommandTemplate       *string `json:"rgbw_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgbw_command_topic`. Available variables: `red`, `green`, `blue` and `white`."
	RgbwCommandTopic          *string `json:"rgbw_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light's RGBW state."
	rgbwCommandFunc           func(mqtt.Message, mqtt.Client)
	RgbwStateTopic            *string `json:"rgbw_state_topic,omitempty"` // "The MQTT topic subscribed to receive RGBW state updates. The expected payload is the RGBW values separated by commas, for example, `255,0,127,64`."
	rgbwStateFunc             func() string
	RgbwValueTemplate         *string `json:"rgbw_value_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGBW value."
	RgbwwCommandTemplate      *string `json:"rgbww_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgbww_command_topic`. Available variables: `red`, `green`, `blue`, `cold_white` and `warm_white`."
	RgbwwCommandTopic         *string `json:"rgbww_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light's RGBWW state."
	rgbwwCommandFunc          func(mqtt.Message, mqtt.Client)
	RgbwwStateTopic           *string `json:"rgbww_state_topic,omitempty"` // "The MQTT topic subscribed to receive RGBWW state updates. The expected payload is the RGBWW values separated by commas, for example, `255,0,127,64,32`."
	rgbwwStateFunc            func() string
	RgbwwValueTemplate        *string `json:"rgbww_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGBWW value."
	Schema                    *string `json:"schema,omitempty"`               // "The schema to use. Must be `default` or omitted to select the default schema."
	StateTopic                *string `json:"state_topic,omitempty"`          // "The MQTT topic subscribed to receive state updates."
	stateFunc                 func() string
	StateValueTemplate        *string `json:"state_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the state value. The template should match the payload `on` and `off` values, so if your light uses `power on` to turn on, your `state_value_template` string should return `power on` when the switch is on. For example if the message is just `on`, your `state_value_template` should be `power {{ value }}`."
	UniqueId                  *string `json:"unique_id,omitempty"`            // "An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception."
	WhiteCommandTopic         *string `json:"white_command_topic,omitempty"`  // "The MQTT topic to publish commands to change the light to white mode with a given brightness."
	whiteCommandFunc          func(mqtt.Message, mqtt.Client)
	WhiteScale                *int    `json:"white_scale,omitempty"`         // "Defines the maximum white level (i.e., 100%) of the MQTT device."
	XyCommandTemplate         *string `json:"xy_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `xy_command_topic`. Available variables: `x` and `y`."
	XyCommandTopic            *string `json:"xy_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light's XY state."
	xyCommandFunc             func(mqtt.Message, mqtt.Client)
	XyStateTopic              *string `json:"xy_state_topic,omitempty"` // "The MQTT topic subscribed to receive XY state updates. The expected payload is the X and Y color values separated by commas, for example, `0.675,0.322`."
	xyStateFunc               func() string
	XyValueTemplate           *string     `json:"xy_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the XY value."
	MQTT                      *MQTTFields `json:"-"`                           // MQTT configuration parameters
	states                    lightState  // Internal Holder of States
	States                    *LightState `json:"-"` // External state update location
}

func NewLight(o *LightOptions) *Light {
	var l Light

	l.States = &o.States
	if !reflect.ValueOf(o.AvailabilityMode).IsZero() {
		l.AvailabilityMode = &o.AvailabilityMode
	}
	if !reflect.ValueOf(o.AvailabilityTemplate).IsZero() {
		l.AvailabilityTemplate = &o.AvailabilityTemplate
	}
	if !reflect.ValueOf(o.AvailabilityFunc).IsZero() {
		l.availabilityFunc = o.AvailabilityFunc
	}
	if !reflect.ValueOf(o.BrightnessCommandTemplate).IsZero() {
		l.BrightnessCommandTemplate = &o.BrightnessCommandTemplate
	}
	if !reflect.ValueOf(o.BrightnessCommandFunc).IsZero() {
		l.brightnessCommandFunc = o.BrightnessCommandFunc
	}
	if !reflect.ValueOf(o.BrightnessScale).IsZero() {
		l.BrightnessScale = &o.BrightnessScale
	}
	if !reflect.ValueOf(o.BrightnessStateFunc).IsZero() {
		l.brightnessStateFunc = o.BrightnessStateFunc
	}
	if !reflect.ValueOf(o.BrightnessValueTemplate).IsZero() {
		l.BrightnessValueTemplate = &o.BrightnessValueTemplate
	}
	if !reflect.ValueOf(o.ColorModeStateFunc).IsZero() {
		l.colorModeStateFunc = o.ColorModeStateFunc
	}
	if !reflect.ValueOf(o.ColorModeValueTemplate).IsZero() {
		l.ColorModeValueTemplate = &o.ColorModeValueTemplate
	}
	if !reflect.ValueOf(o.ColorTempCommandTemplate).IsZero() {
		l.ColorTempCommandTemplate = &o.ColorTempCommandTemplate
	}
	if !reflect.ValueOf(o.ColorTempCommandFunc).IsZero() {
		l.colorTempCommandFunc = o.ColorTempCommandFunc
	}
	if !reflect.ValueOf(o.ColorTempStateFunc).IsZero() {
		l.colorTempStateFunc = o.ColorTempStateFunc
	}
	if !reflect.ValueOf(o.ColorTempValueTemplate).IsZero() {
		l.ColorTempValueTemplate = &o.ColorTempValueTemplate
	}
	if !reflect.ValueOf(o.CommandFunc).IsZero() {
		l.commandFunc = o.CommandFunc
	} else {
		l.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.States.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.EffectCommandTemplate).IsZero() {
		l.EffectCommandTemplate = &o.EffectCommandTemplate
	}
	if !reflect.ValueOf(o.EffectCommandFunc).IsZero() {
		l.effectCommandFunc = o.EffectCommandFunc
	}
	if !reflect.ValueOf(o.EffectList).IsZero() {
		l.EffectList = &o.EffectList
	}
	if !reflect.ValueOf(o.EffectStateFunc).IsZero() {
		l.effectStateFunc = o.EffectStateFunc
	}
	if !reflect.ValueOf(o.EffectValueTemplate).IsZero() {
		l.EffectValueTemplate = &o.EffectValueTemplate
	}
	if !reflect.ValueOf(o.EnabledByDefault).IsZero() {
		l.EnabledByDefault = &o.EnabledByDefault
	}
	if !reflect.ValueOf(o.Encoding).IsZero() {
		l.Encoding = &o.Encoding
	}
	if !reflect.ValueOf(o.EntityCategory).IsZero() {
		l.EntityCategory = &o.EntityCategory
	}
	if !reflect.ValueOf(o.HsCommandTemplate).IsZero() {
		l.HsCommandTemplate = &o.HsCommandTemplate
	}
	if !reflect.ValueOf(o.HsCommandFunc).IsZero() {
		l.hsCommandFunc = o.HsCommandFunc
	}
	if !reflect.ValueOf(o.HsStateFunc).IsZero() {
		l.hsStateFunc = o.HsStateFunc
	}
	if !reflect.ValueOf(o.HsValueTemplate).IsZero() {
		l.HsValueTemplate = &o.HsValueTemplate
	}
	if !reflect.ValueOf(o.Icon).IsZero() {
		l.Icon = &o.Icon
	}
	if !reflect.ValueOf(o.JsonAttributesTemplate).IsZero() {
		l.JsonAttributesTemplate = &o.JsonAttributesTemplate
	}
	if !reflect.ValueOf(o.JsonAttributesFunc).IsZero() {
		l.jsonAttributesFunc = o.JsonAttributesFunc
	}
	if !reflect.ValueOf(o.MaxMireds).IsZero() {
		l.MaxMireds = &o.MaxMireds
	}
	if !reflect.ValueOf(o.MinMireds).IsZero() {
		l.MinMireds = &o.MinMireds
	}
	if !reflect.ValueOf(o.Name).IsZero() {
		l.Name = &o.Name
	}
	if !reflect.ValueOf(o.ObjectId).IsZero() {
		l.ObjectId = &o.ObjectId
	}
	if !reflect.ValueOf(o.OnCommandType).IsZero() {
		l.OnCommandType = &o.OnCommandType
	}
	if !reflect.ValueOf(o.Optimistic).IsZero() {
		l.Optimistic = &o.Optimistic
	}
	if !reflect.ValueOf(o.PayloadAvailable).IsZero() {
		l.PayloadAvailable = &o.PayloadAvailable
	}
	if !reflect.ValueOf(o.PayloadNotAvailable).IsZero() {
		l.PayloadNotAvailable = &o.PayloadNotAvailable
	}
	if !reflect.ValueOf(o.PayloadOff).IsZero() {
		l.PayloadOff = &o.PayloadOff
	}
	if !reflect.ValueOf(o.PayloadOn).IsZero() {
		l.PayloadOn = &o.PayloadOn
	}
	if !reflect.ValueOf(o.Qos).IsZero() {
		l.Qos = &o.Qos
	}
	if !reflect.ValueOf(o.Retain).IsZero() {
		l.Retain = &o.Retain
	}
	if !reflect.ValueOf(o.RgbCommandTemplate).IsZero() {
		l.RgbCommandTemplate = &o.RgbCommandTemplate
	}
	if !reflect.ValueOf(o.RgbCommandFunc).IsZero() {
		l.rgbCommandFunc = o.RgbCommandFunc
	}
	if !reflect.ValueOf(o.RgbStateFunc).IsZero() {
		l.rgbStateFunc = o.RgbStateFunc
	}
	if !reflect.ValueOf(o.RgbValueTemplate).IsZero() {
		l.RgbValueTemplate = &o.RgbValueTemplate
	}
	if !reflect.ValueOf(o.RgbwCommandTemplate).IsZero() {
		l.RgbwCommandTemplate = &o.RgbwCommandTemplate
	}
	if !reflect.ValueOf(o.RgbwCommandFunc).IsZero() {
		l.rgbwCommandFunc = o.RgbwCommandFunc
	}
	if !reflect.ValueOf(o.RgbwStateFunc).IsZero() {
		l.rgbwStateFunc = o.RgbwStateFunc
	}
	if !reflect.ValueOf(o.RgbwValueTemplate).IsZero() {
		l.RgbwValueTemplate = &o.RgbwValueTemplate
	}
	if !reflect.ValueOf(o.RgbwwCommandTemplate).IsZero() {
		l.RgbwwCommandTemplate = &o.RgbwwCommandTemplate
	}
	if !reflect.ValueOf(o.RgbwwCommandFunc).IsZero() {
		l.rgbwwCommandFunc = o.RgbwwCommandFunc
	}
	if !reflect.ValueOf(o.RgbwwStateFunc).IsZero() {
		l.rgbwwStateFunc = o.RgbwwStateFunc
	}
	if !reflect.ValueOf(o.RgbwwValueTemplate).IsZero() {
		l.RgbwwValueTemplate = &o.RgbwwValueTemplate
	}
	if !reflect.ValueOf(o.Schema).IsZero() {
		l.Schema = &o.Schema
	}
	if !reflect.ValueOf(o.StateFunc).IsZero() {
		l.stateFunc = o.StateFunc
	} else {
		l.stateFunc = func() string {
			return l.States.State
		}
	}
	if !reflect.ValueOf(o.StateValueTemplate).IsZero() {
		l.StateValueTemplate = &o.StateValueTemplate
	}
	if !reflect.ValueOf(o.UniqueId).IsZero() {
		l.UniqueId = &o.UniqueId
	}
	if !reflect.ValueOf(o.WhiteCommandFunc).IsZero() {
		l.whiteCommandFunc = o.WhiteCommandFunc
	}
	if !reflect.ValueOf(o.WhiteScale).IsZero() {
		l.WhiteScale = &o.WhiteScale
	}
	if !reflect.ValueOf(o.XyCommandTemplate).IsZero() {
		l.XyCommandTemplate = &o.XyCommandTemplate
	}
	if !reflect.ValueOf(o.XyCommandFunc).IsZero() {
		l.xyCommandFunc = o.XyCommandFunc
	}
	if !reflect.ValueOf(o.XyStateFunc).IsZero() {
		l.xyStateFunc = o.XyStateFunc
	}
	if !reflect.ValueOf(o.XyValueTemplate).IsZero() {
		l.XyValueTemplate = &o.XyValueTemplate
	}
	return &l
}

type lightState struct {
	Availability   *string
	Brightness     *string
	ColorMode      *string
	ColorTemp      *string
	Effect         *string
	Hs             *string
	JsonAttributes *string
	Rgb            *string
	Rgbw           *string
	Rgbww          *string
	State          *string
	Xy             *string
}
type LightState struct {
	Brightness     string
	ColorMode      string
	ColorTemp      string
	Effect         string
	Hs             string
	JsonAttributes string
	Rgb            string
	Rgbw           string
	Rgbww          string
	State          string
	Xy             string
}

func (d *Light) SetBrightness(s string) {
	d.States.Brightness = s
	d.UpdateState()
}
func (d *Light) SetColorMode(s string) {
	d.States.ColorMode = s
	d.UpdateState()
}
func (d *Light) SetColorTemp(s string) {
	d.States.ColorTemp = s
	d.UpdateState()
}
func (d *Light) SetEffect(s string) {
	d.States.Effect = s
	d.UpdateState()
}
func (d *Light) SetHs(s string) {
	d.States.Hs = s
	d.UpdateState()
}
func (d *Light) SetJsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Light) SetRgb(s string) {
	d.States.Rgb = s
	d.UpdateState()
}
func (d *Light) SetRgbw(s string) {
	d.States.Rgbw = s
	d.UpdateState()
}
func (d *Light) SetRgbww(s string) {
	d.States.Rgbww = s
	d.UpdateState()
}
func (d *Light) SetState(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Light) SetXy(s string) {
	d.States.Xy = s
	d.UpdateState()
}
func (d *Light) GetRawId() string {
	return "light"
}
func (d *Light) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Light) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Light) GetName() string {
	return *d.Name
}
func (d *Light) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Light) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.availabilityFunc()
		if d.states.Availability == nil || state != *d.states.Availability || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
		}
	}
	if d.BrightnessStateTopic != nil {
		state := d.brightnessStateFunc()
		if d.states.Brightness == nil || state != *d.states.Brightness || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.BrightnessStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Brightness = &state
		}
	}
	if d.ColorModeStateTopic != nil {
		state := d.colorModeStateFunc()
		if d.states.ColorMode == nil || state != *d.states.ColorMode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ColorModeStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.ColorMode = &state
		}
	}
	if d.ColorTempStateTopic != nil {
		state := d.colorTempStateFunc()
		if d.states.ColorTemp == nil || state != *d.states.ColorTemp || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ColorTempStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.ColorTemp = &state
		}
	}
	if d.EffectStateTopic != nil {
		state := d.effectStateFunc()
		if d.states.Effect == nil || state != *d.states.Effect || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.EffectStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Effect = &state
		}
	}
	if d.HsStateTopic != nil {
		state := d.hsStateFunc()
		if d.states.Hs == nil || state != *d.states.Hs || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.HsStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Hs = &state
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
	if d.RgbStateTopic != nil {
		state := d.rgbStateFunc()
		if d.states.Rgb == nil || state != *d.states.Rgb || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.RgbStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Rgb = &state
		}
	}
	if d.RgbwStateTopic != nil {
		state := d.rgbwStateFunc()
		if d.states.Rgbw == nil || state != *d.states.Rgbw || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.RgbwStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Rgbw = &state
		}
	}
	if d.RgbwwStateTopic != nil {
		state := d.rgbwwStateFunc()
		if d.states.Rgbww == nil || state != *d.states.Rgbww || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.RgbwwStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Rgbww = &state
		}
	}
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
	if d.XyStateTopic != nil {
		state := d.xyStateFunc()
		if d.states.Xy == nil || state != *d.states.Xy || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.XyStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Xy = &state
		}
	}
}
func (d *Light) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.BrightnessCommandTopic != nil {
		t := c.Subscribe(*d.BrightnessCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ColorTempCommandTopic != nil {
		t := c.Subscribe(*d.ColorTempCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.CommandTopic != nil {
		t := c.Subscribe(*d.CommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.EffectCommandTopic != nil {
		t := c.Subscribe(*d.EffectCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.HsCommandTopic != nil {
		t := c.Subscribe(*d.HsCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbCommandTopic != nil {
		t := c.Subscribe(*d.RgbCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbwCommandTopic != nil {
		t := c.Subscribe(*d.RgbwCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbwwCommandTopic != nil {
		t := c.Subscribe(*d.RgbwwCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.WhiteCommandTopic != nil {
		t := c.Subscribe(*d.WhiteCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.XyCommandTopic != nil {
		t := c.Subscribe(*d.XyCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d *Light) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, false, "offline")
	token.WaitTimeout(common.WaitTimeout)
	if d.BrightnessCommandTopic != nil {
		t := c.Unsubscribe(*d.BrightnessCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ColorTempCommandTopic != nil {
		t := c.Unsubscribe(*d.ColorTempCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.CommandTopic != nil {
		t := c.Unsubscribe(*d.CommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.EffectCommandTopic != nil {
		t := c.Unsubscribe(*d.EffectCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.HsCommandTopic != nil {
		t := c.Unsubscribe(*d.HsCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbCommandTopic != nil {
		t := c.Unsubscribe(*d.RgbCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbwCommandTopic != nil {
		t := c.Unsubscribe(*d.RgbwCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbwwCommandTopic != nil {
		t := c.Unsubscribe(*d.RgbwwCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.WhiteCommandTopic != nil {
		t := c.Unsubscribe(*d.WhiteCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.XyCommandTopic != nil {
		t := c.Unsubscribe(*d.XyCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Light) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Light) Initialize() {
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
func (d *Light) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Light) populateTopics() {
	if d.availabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.brightnessCommandFunc != nil {
		d.BrightnessCommandTopic = new(string)
		*d.BrightnessCommandTopic = GetTopic(d, "brightness_command_topic")
		common.TopicStore[*d.BrightnessCommandTopic] = &d.brightnessCommandFunc
	}
	if d.brightnessStateFunc != nil {
		d.BrightnessStateTopic = new(string)
		*d.BrightnessStateTopic = GetTopic(d, "brightness_state_topic")
	}
	if d.colorModeStateFunc != nil {
		d.ColorModeStateTopic = new(string)
		*d.ColorModeStateTopic = GetTopic(d, "color_mode_state_topic")
	}
	if d.colorTempCommandFunc != nil {
		d.ColorTempCommandTopic = new(string)
		*d.ColorTempCommandTopic = GetTopic(d, "color_temp_command_topic")
		common.TopicStore[*d.ColorTempCommandTopic] = &d.colorTempCommandFunc
	}
	if d.colorTempStateFunc != nil {
		d.ColorTempStateTopic = new(string)
		*d.ColorTempStateTopic = GetTopic(d, "color_temp_state_topic")
	}
	if d.commandFunc != nil {
		d.CommandTopic = new(string)
		*d.CommandTopic = GetTopic(d, "command_topic")
		common.TopicStore[*d.CommandTopic] = &d.commandFunc
	}
	if d.effectCommandFunc != nil {
		d.EffectCommandTopic = new(string)
		*d.EffectCommandTopic = GetTopic(d, "effect_command_topic")
		common.TopicStore[*d.EffectCommandTopic] = &d.effectCommandFunc
	}
	if d.effectStateFunc != nil {
		d.EffectStateTopic = new(string)
		*d.EffectStateTopic = GetTopic(d, "effect_state_topic")
	}
	if d.hsCommandFunc != nil {
		d.HsCommandTopic = new(string)
		*d.HsCommandTopic = GetTopic(d, "hs_command_topic")
		common.TopicStore[*d.HsCommandTopic] = &d.hsCommandFunc
	}
	if d.hsStateFunc != nil {
		d.HsStateTopic = new(string)
		*d.HsStateTopic = GetTopic(d, "hs_state_topic")
	}
	if d.jsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
	}
	if d.rgbCommandFunc != nil {
		d.RgbCommandTopic = new(string)
		*d.RgbCommandTopic = GetTopic(d, "rgb_command_topic")
		common.TopicStore[*d.RgbCommandTopic] = &d.rgbCommandFunc
	}
	if d.rgbStateFunc != nil {
		d.RgbStateTopic = new(string)
		*d.RgbStateTopic = GetTopic(d, "rgb_state_topic")
	}
	if d.rgbwCommandFunc != nil {
		d.RgbwCommandTopic = new(string)
		*d.RgbwCommandTopic = GetTopic(d, "rgbw_command_topic")
		common.TopicStore[*d.RgbwCommandTopic] = &d.rgbwCommandFunc
	}
	if d.rgbwStateFunc != nil {
		d.RgbwStateTopic = new(string)
		*d.RgbwStateTopic = GetTopic(d, "rgbw_state_topic")
	}
	if d.rgbwwCommandFunc != nil {
		d.RgbwwCommandTopic = new(string)
		*d.RgbwwCommandTopic = GetTopic(d, "rgbww_command_topic")
		common.TopicStore[*d.RgbwwCommandTopic] = &d.rgbwwCommandFunc
	}
	if d.rgbwwStateFunc != nil {
		d.RgbwwStateTopic = new(string)
		*d.RgbwwStateTopic = GetTopic(d, "rgbww_state_topic")
	}
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.whiteCommandFunc != nil {
		d.WhiteCommandTopic = new(string)
		*d.WhiteCommandTopic = GetTopic(d, "white_command_topic")
		common.TopicStore[*d.WhiteCommandTopic] = &d.whiteCommandFunc
	}
	if d.xyCommandFunc != nil {
		d.XyCommandTopic = new(string)
		*d.XyCommandTopic = GetTopic(d, "xy_command_topic")
		common.TopicStore[*d.XyCommandTopic] = &d.xyCommandFunc
	}
	if d.xyStateFunc != nil {
		d.XyStateTopic = new(string)
		*d.XyStateTopic = GetTopic(d, "xy_state_topic")
	}
}
func (d *Light) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Light) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
