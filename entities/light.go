package entities

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	strcase "github.com/iancoleman/strcase"
	common "github.com/kjbreil/hass-mqtt/common"
	"log"
	"reflect"
	"strings"
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

func NewLight(o *LightOptions) (*Light, error) {
	var l Light

	l.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		l.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		l.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		l.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.brightnessCommandTemplate).IsZero() {
		l.BrightnessCommandTemplate = &o.brightnessCommandTemplate
	}
	if !reflect.ValueOf(o.brightnessCommandFunc).IsZero() {
		l.brightnessCommandFunc = o.brightnessCommandFunc
	}
	if !reflect.ValueOf(o.brightnessScale).IsZero() {
		l.BrightnessScale = &o.brightnessScale
	}
	if !reflect.ValueOf(o.brightnessStateFunc).IsZero() {
		l.brightnessStateFunc = o.brightnessStateFunc
	}
	if !reflect.ValueOf(o.brightnessValueTemplate).IsZero() {
		l.BrightnessValueTemplate = &o.brightnessValueTemplate
	}
	if !reflect.ValueOf(o.colorModeStateFunc).IsZero() {
		l.colorModeStateFunc = o.colorModeStateFunc
	}
	if !reflect.ValueOf(o.colorModeValueTemplate).IsZero() {
		l.ColorModeValueTemplate = &o.colorModeValueTemplate
	}
	if !reflect.ValueOf(o.colorTempCommandTemplate).IsZero() {
		l.ColorTempCommandTemplate = &o.colorTempCommandTemplate
	}
	if !reflect.ValueOf(o.colorTempCommandFunc).IsZero() {
		l.colorTempCommandFunc = o.colorTempCommandFunc
	}
	if !reflect.ValueOf(o.colorTempStateFunc).IsZero() {
		l.colorTempStateFunc = o.colorTempStateFunc
	}
	if !reflect.ValueOf(o.colorTempValueTemplate).IsZero() {
		l.ColorTempValueTemplate = &o.colorTempValueTemplate
	}
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		l.commandFunc = o.commandFunc
	} else {
		l.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.effectCommandTemplate).IsZero() {
		l.EffectCommandTemplate = &o.effectCommandTemplate
	}
	if !reflect.ValueOf(o.effectCommandFunc).IsZero() {
		l.effectCommandFunc = o.effectCommandFunc
	}
	if !reflect.ValueOf(o.effectList).IsZero() {
		l.EffectList = &o.effectList
	}
	if !reflect.ValueOf(o.effectStateFunc).IsZero() {
		l.effectStateFunc = o.effectStateFunc
	}
	if !reflect.ValueOf(o.effectValueTemplate).IsZero() {
		l.EffectValueTemplate = &o.effectValueTemplate
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		l.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		l.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		l.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.hsCommandTemplate).IsZero() {
		l.HsCommandTemplate = &o.hsCommandTemplate
	}
	if !reflect.ValueOf(o.hsCommandFunc).IsZero() {
		l.hsCommandFunc = o.hsCommandFunc
	}
	if !reflect.ValueOf(o.hsStateFunc).IsZero() {
		l.hsStateFunc = o.hsStateFunc
	}
	if !reflect.ValueOf(o.hsValueTemplate).IsZero() {
		l.HsValueTemplate = &o.hsValueTemplate
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		l.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		l.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		l.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.maxMireds).IsZero() {
		l.MaxMireds = &o.maxMireds
	}
	if !reflect.ValueOf(o.minMireds).IsZero() {
		l.MinMireds = &o.minMireds
	}
	if !reflect.ValueOf(o.name).IsZero() {
		l.Name = &o.name
	} else {
		return nil, fmt.Errorf("name not defined")
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		l.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.onCommandType).IsZero() {
		l.OnCommandType = &o.onCommandType
	}
	if !reflect.ValueOf(o.optimistic).IsZero() {
		l.Optimistic = &o.optimistic
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		l.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		l.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadOff).IsZero() {
		l.PayloadOff = &o.payloadOff
	}
	if !reflect.ValueOf(o.payloadOn).IsZero() {
		l.PayloadOn = &o.payloadOn
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		l.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		l.Retain = &o.retain
	}
	if !reflect.ValueOf(o.rgbCommandTemplate).IsZero() {
		l.RgbCommandTemplate = &o.rgbCommandTemplate
	}
	if !reflect.ValueOf(o.rgbCommandFunc).IsZero() {
		l.rgbCommandFunc = o.rgbCommandFunc
	}
	if !reflect.ValueOf(o.rgbStateFunc).IsZero() {
		l.rgbStateFunc = o.rgbStateFunc
	}
	if !reflect.ValueOf(o.rgbValueTemplate).IsZero() {
		l.RgbValueTemplate = &o.rgbValueTemplate
	}
	if !reflect.ValueOf(o.rgbwCommandTemplate).IsZero() {
		l.RgbwCommandTemplate = &o.rgbwCommandTemplate
	}
	if !reflect.ValueOf(o.rgbwCommandFunc).IsZero() {
		l.rgbwCommandFunc = o.rgbwCommandFunc
	}
	if !reflect.ValueOf(o.rgbwStateFunc).IsZero() {
		l.rgbwStateFunc = o.rgbwStateFunc
	}
	if !reflect.ValueOf(o.rgbwValueTemplate).IsZero() {
		l.RgbwValueTemplate = &o.rgbwValueTemplate
	}
	if !reflect.ValueOf(o.rgbwwCommandTemplate).IsZero() {
		l.RgbwwCommandTemplate = &o.rgbwwCommandTemplate
	}
	if !reflect.ValueOf(o.rgbwwCommandFunc).IsZero() {
		l.rgbwwCommandFunc = o.rgbwwCommandFunc
	}
	if !reflect.ValueOf(o.rgbwwStateFunc).IsZero() {
		l.rgbwwStateFunc = o.rgbwwStateFunc
	}
	if !reflect.ValueOf(o.rgbwwValueTemplate).IsZero() {
		l.RgbwwValueTemplate = &o.rgbwwValueTemplate
	}
	if !reflect.ValueOf(o.schema).IsZero() {
		l.Schema = &o.schema
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		l.stateFunc = o.stateFunc
	} else {
		l.stateFunc = func() string {
			return l.States.State
		}
	}
	if !reflect.ValueOf(o.stateValueTemplate).IsZero() {
		l.StateValueTemplate = &o.stateValueTemplate
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		l.UniqueId = &o.uniqueId
	} else {
		uniqueId := strcase.ToDelimited(o.name, uint8(0x2d))
		uniqueId = strings.ReplaceAll(uniqueId, "'", "_")
		l.UniqueId = &uniqueId
	}
	if !reflect.ValueOf(o.whiteCommandFunc).IsZero() {
		l.whiteCommandFunc = o.whiteCommandFunc
	}
	if !reflect.ValueOf(o.whiteScale).IsZero() {
		l.WhiteScale = &o.whiteScale
	}
	if !reflect.ValueOf(o.xyCommandTemplate).IsZero() {
		l.XyCommandTemplate = &o.xyCommandTemplate
	}
	if !reflect.ValueOf(o.xyCommandFunc).IsZero() {
		l.xyCommandFunc = o.xyCommandFunc
	}
	if !reflect.ValueOf(o.xyStateFunc).IsZero() {
		l.xyStateFunc = o.xyStateFunc
	}
	if !reflect.ValueOf(o.xyValueTemplate).IsZero() {
		l.XyValueTemplate = &o.xyValueTemplate
	}
	return &l, nil
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

func (d *Light) Brightness(s string) {
	d.States.Brightness = s
	d.UpdateState()
}
func (d *Light) ColorMode(s string) {
	d.States.ColorMode = s
	d.UpdateState()
}
func (d *Light) ColorTemp(s string) {
	d.States.ColorTemp = s
	d.UpdateState()
}
func (d *Light) Effect(s string) {
	d.States.Effect = s
	d.UpdateState()
}
func (d *Light) Hs(s string) {
	d.States.Hs = s
	d.UpdateState()
}
func (d *Light) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Light) Rgb(s string) {
	d.States.Rgb = s
	d.UpdateState()
}
func (d *Light) Rgbw(s string) {
	d.States.Rgbw = s
	d.UpdateState()
}
func (d *Light) Rgbww(s string) {
	d.States.Rgbww = s
	d.UpdateState()
}
func (d *Light) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Light) Xy(s string) {
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
func (d *Light) GetDomainEntity() string {
	return fmt.Sprintf("light.%s", strings.ReplaceAll(*d.UniqueId, "-", "_"))
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
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
		}
	}
	if d.BrightnessStateTopic != nil {
		state := d.brightnessStateFunc()
		if d.states.Brightness == nil || state != *d.states.Brightness || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.BrightnessStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Brightness = &state
		}
	}
	if d.ColorModeStateTopic != nil {
		state := d.colorModeStateFunc()
		if d.states.ColorMode == nil || state != *d.states.ColorMode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ColorModeStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.ColorMode = &state
		}
	}
	if d.ColorTempStateTopic != nil {
		state := d.colorTempStateFunc()
		if d.states.ColorTemp == nil || state != *d.states.ColorTemp || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ColorTempStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.ColorTemp = &state
		}
	}
	if d.EffectStateTopic != nil {
		state := d.effectStateFunc()
		if d.states.Effect == nil || state != *d.states.Effect || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.EffectStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Effect = &state
		}
	}
	if d.HsStateTopic != nil {
		state := d.hsStateFunc()
		if d.states.Hs == nil || state != *d.states.Hs || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.HsStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Hs = &state
		}
	}
	if d.JsonAttributesTopic != nil {
		state := d.jsonAttributesFunc()
		if d.states.JsonAttributes == nil || state != *d.states.JsonAttributes || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.JsonAttributesTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.JsonAttributes = &state
		}
	}
	if d.RgbStateTopic != nil {
		state := d.rgbStateFunc()
		if d.states.Rgb == nil || state != *d.states.Rgb || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.RgbStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Rgb = &state
		}
	}
	if d.RgbwStateTopic != nil {
		state := d.rgbwStateFunc()
		if d.states.Rgbw == nil || state != *d.states.Rgbw || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.RgbwStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Rgbw = &state
		}
	}
	if d.RgbwwStateTopic != nil {
		state := d.rgbwwStateFunc()
		if d.states.Rgbww == nil || state != *d.states.Rgbww || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.RgbwwStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Rgbww = &state
		}
	}
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
	if d.XyStateTopic != nil {
		state := d.xyStateFunc()
		if d.states.Xy == nil || state != *d.states.Xy || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.XyStateTopic, byte(*d.Qos), true, state)
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
