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
type Climate struct {
	ActionTemplate                     *string `json:"action_template,omitempty"` // "A template to render the value received on the `action_topic` with."
	ActionTopic                        *string `json:"action_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the current action. If this is set, the climate graph uses the value received as data source. A \"None\" payload resets the current action state. An empty payload is ignored. Valid action values: `off`, `heating`, `cooling`, `drying`, `idle`, `fan`."
	actionFunc                         func() string
	AvailabilityMode                   *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate               *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic                  *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc                   func() string
	CurrentHumidityTemplate            *string `json:"current_humidity_template,omitempty"` // "A template with which the value received on `current_humidity_topic` will be rendered."
	CurrentHumidityTopic               *string `json:"current_humidity_topic,omitempty"`    // "The MQTT topic on which to listen for the current humidity. A `\"None\"` value received will reset the current humidity. Empty values (`'''`) will be ignored."
	currentHumidityFunc                func() string
	CurrentTemperatureTemplate         *string `json:"current_temperature_template,omitempty"` // "A template with which the value received on `current_temperature_topic` will be rendered."
	CurrentTemperatureTopic            *string `json:"current_temperature_topic,omitempty"`    // "The MQTT topic on which to listen for the current temperature. A `\"None\"` value received will reset the current temperature. Empty values (`'''`) will be ignored."
	currentTemperatureFunc             func() string
	Device                             Device  `json:"device,omitempty"`                    // Device configuration parameters
	EnabledByDefault                   *bool   `json:"enabled_by_default,omitempty"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding                           *string `json:"encoding,omitempty"`                  // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                     *string `json:"entity_category,omitempty"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	EntityPicture                      *string `json:"entity_picture,omitempty"`            // "Picture URL for the entity."
	FanModeCommandTemplate             *string `json:"fan_mode_command_template,omitempty"` // "A template to render the value sent to the `fan_mode_command_topic` with."
	FanModeCommandTopic                *string `json:"fan_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the fan mode."
	fanModeCommandFunc                 func(mqtt.Message, mqtt.Client)
	FanModeStateTemplate               *string `json:"fan_mode_state_template,omitempty"` // "A template to render the value received on the `fan_mode_state_topic` with."
	FanModeStateTopic                  *string `json:"fan_mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC fan mode. If this is not set, the fan mode works in optimistic mode (see below). A \"None\" payload resets the fan mode state. An empty payload is ignored."
	fanModeStateFunc                   func() string
	FanModes                           *([]string) `json:"fan_modes,omitempty"`                // "A list of supported fan modes."
	Icon                               *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Initial                            *float64    `json:"initial,omitempty"`                  // "Set the initial target temperature. The default value depends on the temperature unit and will be 21° or 69.8°F."
	JsonAttributesTemplate             *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic                *string     `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc                 func() string
	MaxHumidity                        *float64 `json:"max_humidity,omitempty"`          // "The minimum target humidity percentage that can be set."
	MaxTemp                            *float64 `json:"max_temp,omitempty"`              // "Maximum set point available. The default value depends on the temperature unit, and will be 35°C or 95°F."
	MinHumidity                        *float64 `json:"min_humidity,omitempty"`          // "The maximum target humidity percentage that can be set."
	MinTemp                            *float64 `json:"min_temp,omitempty"`              // "Minimum set point available. The default value depends on the temperature unit, and will be 7°C or 44.6°F."
	ModeCommandTemplate                *string  `json:"mode_command_template,omitempty"` // "A template to render the value sent to the `mode_command_topic` with."
	ModeCommandTopic                   *string  `json:"mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the HVAC operation mode."
	modeCommandFunc                    func(mqtt.Message, mqtt.Client)
	ModeStateTemplate                  *string `json:"mode_state_template,omitempty"` // "A template to render the value received on the `mode_state_topic` with."
	ModeStateTopic                     *string `json:"mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC operation mode. If this is not set, the operation mode works in optimistic mode (see below). A \"None\" payload resets to an `unknown` state. An empty payload is ignored."
	modeStateFunc                      func() string
	Modes                              *([]string) `json:"modes,omitempty"`                  // "A list of supported modes. Needs to be a subset of the default values."
	Name                               *string     `json:"name,omitempty"`                   // "The name of the HVAC. Can be set to `null` if only the device name is relevant."
	ObjectId                           *string     `json:"object_id,omitempty"`              // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	Optimistic                         *bool       `json:"optimistic,omitempty"`             // "Flag that defines if the climate works in optimistic mode"
	PayloadAvailable                   *string     `json:"payload_available,omitempty"`      // "The payload that represents the available state."
	PayloadNotAvailable                *string     `json:"payload_not_available,omitempty"`  // "The payload that represents the unavailable state."
	PayloadOff                         *string     `json:"payload_off,omitempty"`            // "The payload sent to turn off the device."
	PayloadOn                          *string     `json:"payload_on,omitempty"`             // "The payload sent to turn the device on."
	PowerCommandTemplate               *string     `json:"power_command_template,omitempty"` // "A template to render the value sent to the `power_command_topic` with. The `value` parameter is the payload set for `payload_on` or `payload_off`."
	PowerCommandTopic                  *string     `json:"power_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the HVAC power state. Sends the payload configured with `payload_on` if the climate is turned on via the `climate.turn_on`, or the payload configured with `payload_off` if the climate is turned off via the `climate.turn_off` action. Note that `optimistic` mode is not supported through `climate.turn_on` and `climate.turn_off` actions. When called, these actions will send a power command to the device but will not optimistically update the state of the climate entity. The climate device should report its state back via `mode_state_topic`."
	powerCommandFunc                   func(mqtt.Message, mqtt.Client)
	Precision                          *float64 `json:"precision,omitempty"`                    // "The desired precision for this device. Can be used to match your actual thermostat's precision. Supported values are `0.1`, `0.5` and `1.0`."
	PresetModeCommandTemplate          *string  `json:"preset_mode_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommandTopic             *string  `json:"preset_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the preset mode."
	presetModeCommandFunc              func(mqtt.Message, mqtt.Client)
	PresetModeStateTopic               *string `json:"preset_mode_state_topic,omitempty"` // "The MQTT topic subscribed to receive climate speed based on presets. When preset 'none' is received or `None` the `preset_mode` will be reset."
	presetModeStateFunc                func() string
	PresetModeValueTemplate            *string     `json:"preset_mode_value_template,omitempty"`             // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                        *([]string) `json:"preset_modes,omitempty"`                           // "List of preset modes this climate is supporting. Common examples include `eco`, `away`, `boost`, `comfort`, `home`, `sleep` and `activity`."
	Qos                                *int        `json:"qos,omitempty"`                                    // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                             *bool       `json:"retain,omitempty"`                                 // "Defines if published messages should have the retain flag set."
	SwingHorizontalModeCommandTemplate *string     `json:"swing_horizontal_mode_command_template,omitempty"` // "A template to render the value sent to the `swing_horizontal_mode_command_topic` with."
	SwingHorizontalModeCommandTopic    *string     `json:"swing_horizontal_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the swing horizontal mode."
	swingHorizontalModeCommandFunc     func(mqtt.Message, mqtt.Client)
	SwingHorizontalModeStateTemplate   *string `json:"swing_horizontal_mode_state_template,omitempty"` // "A template to render the value received on the `swing_horizontal_mode_state_topic` with."
	SwingHorizontalModeStateTopic      *string `json:"swing_horizontal_mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC swing horizontal mode. If this is not set, the swing horizontal mode works in optimistic mode (see below)."
	swingHorizontalModeStateFunc       func() string
	SwingHorizontalModes               *([]string) `json:"swing_horizontal_modes,omitempty"`      // "A list of supported swing horizontal modes."
	SwingModeCommandTemplate           *string     `json:"swing_mode_command_template,omitempty"` // "A template to render the value sent to the `swing_mode_command_topic` with."
	SwingModeCommandTopic              *string     `json:"swing_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the swing mode."
	swingModeCommandFunc               func(mqtt.Message, mqtt.Client)
	SwingModeStateTemplate             *string `json:"swing_mode_state_template,omitempty"` // "A template to render the value received on the `swing_mode_state_topic` with."
	SwingModeStateTopic                *string `json:"swing_mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC swing mode. If this is not set, the swing mode works in optimistic mode (see below)."
	swingModeStateFunc                 func() string
	SwingModes                         *([]string) `json:"swing_modes,omitempty"`                      // "A list of supported swing modes."
	TargetHumidityCommandTemplate      *string     `json:"target_humidity_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to generate the payload to send to `target_humidity_command_topic`."
	TargetHumidityCommandTopic         *string     `json:"target_humidity_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the target humidity."
	targetHumidityCommandFunc          func(mqtt.Message, mqtt.Client)
	TargetHumidityStateTemplate        *string `json:"target_humidity_state_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract a value for the climate `target_humidity` state."
	TargetHumidityStateTopic           *string `json:"target_humidity_state_topic,omitempty"`    // "The MQTT topic subscribed to receive the target humidity. If this is not set, the target humidity works in optimistic mode (see below). A `\"None\"` value received will reset the target humidity. Empty values (`'''`) will be ignored."
	targetHumidityStateFunc            func() string
	TempStep                           *float64 `json:"temp_step,omitempty"`                    // "Step size for temperature set point."
	TemperatureCommandTemplate         *string  `json:"temperature_command_template,omitempty"` // "A template to render the value sent to the `temperature_command_topic` with."
	TemperatureCommandTopic            *string  `json:"temperature_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the target temperature."
	temperatureCommandFunc             func(mqtt.Message, mqtt.Client)
	TemperatureHighCommandTemplate     *string `json:"temperature_high_command_template,omitempty"` // "A template to render the value sent to the `temperature_high_command_topic` with."
	TemperatureHighCommandTopic        *string `json:"temperature_high_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the upper target temperature."
	temperatureHighCommandFunc         func(mqtt.Message, mqtt.Client)
	TemperatureHighStateTemplate       *string `json:"temperature_high_state_template,omitempty"` // "A template to render the value received on the `temperature_high_state_topic` with. A `\"None\"` value received will reset the upper temperature setpoint. Empty values (`\"\"'`) will be ignored."
	TemperatureHighStateTopic          *string `json:"temperature_high_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes in the upper target temperature. If this is not set, the upper target temperature works in optimistic mode (see below)."
	temperatureHighStateFunc           func() string
	TemperatureLowCommandTemplate      *string `json:"temperature_low_command_template,omitempty"` // "A template to render the value sent to the `temperature_low_command_topic` with."
	TemperatureLowCommandTopic         *string `json:"temperature_low_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the lower target temperature."
	temperatureLowCommandFunc          func(mqtt.Message, mqtt.Client)
	TemperatureLowStateTemplate        *string `json:"temperature_low_state_template,omitempty"` // "A template to render the value received on the `temperature_low_state_topic` with. A `\"None\"` value received will reset the lower temperature setpoint. Empty values (`\"\"`) will be ignored."
	TemperatureLowStateTopic           *string `json:"temperature_low_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes in the lower target temperature. If this is not set, the lower target temperature works in optimistic mode (see below)."
	temperatureLowStateFunc            func() string
	TemperatureStateTemplate           *string `json:"temperature_state_template,omitempty"` // "A template to render the value received on the `temperature_state_topic` with."
	TemperatureStateTopic              *string `json:"temperature_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes in the target temperature. If this is not set, the target temperature works in optimistic mode (see below). A `\"None\"` value received will reset the temperature set point. Empty values (`'''`) will be ignored."
	temperatureStateFunc               func() string
	TemperatureUnit                    *string       `json:"temperature_unit,omitempty"` // "Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit."
	UniqueId                           *string       `json:"unique_id,omitempty"`        // "An ID that uniquely identifies this HVAC device. If two HVAC devices have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	ValueTemplate                      *string       `json:"value_template,omitempty"`   // "Default template to render the payloads on *all* `*_state_topic`s with."
	MQTT                               *MQTTFields   `json:"-"`                          // MQTT configuration parameters
	states                             climateState  // Internal Holder of States
	States                             *ClimateState `json:"-"` // External state update location
}

func (d *Climate) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
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
	if d.PowerCommandTopic != nil {
		t := c.Subscribe(*d.PowerCommandTopic, 0, d.MQTT.MessageHandler)
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
	if d.SwingHorizontalModeCommandTopic != nil {
		t := c.Subscribe(*d.SwingHorizontalModeCommandTopic, 0, d.MQTT.MessageHandler)
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
	if d.PowerCommandTopic != nil {
		t := c.Unsubscribe(*d.PowerCommandTopic)
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
	if d.SwingHorizontalModeCommandTopic != nil {
		t := c.Unsubscribe(*d.SwingHorizontalModeCommandTopic)
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
func NewClimate(o *ClimateOptions) (*Climate, error) {
	var c Climate

	c.States = &o.states
	if !reflect.ValueOf(o.actionTemplate).IsZero() {
		c.ActionTemplate = &o.actionTemplate
	}
	if !reflect.ValueOf(o.actionFunc).IsZero() {
		c.actionFunc = o.actionFunc
	}
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		c.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		c.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		c.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.currentHumidityTemplate).IsZero() {
		c.CurrentHumidityTemplate = &o.currentHumidityTemplate
	}
	if !reflect.ValueOf(o.currentHumidityFunc).IsZero() {
		c.currentHumidityFunc = o.currentHumidityFunc
	}
	if !reflect.ValueOf(o.currentTemperatureTemplate).IsZero() {
		c.CurrentTemperatureTemplate = &o.currentTemperatureTemplate
	}
	if !reflect.ValueOf(o.currentTemperatureFunc).IsZero() {
		c.currentTemperatureFunc = o.currentTemperatureFunc
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		c.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		c.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		c.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.entityPicture).IsZero() {
		c.EntityPicture = &o.entityPicture
	}
	if !reflect.ValueOf(o.fanModeCommandTemplate).IsZero() {
		c.FanModeCommandTemplate = &o.fanModeCommandTemplate
	}
	if !reflect.ValueOf(o.fanModeCommandFunc).IsZero() {
		c.fanModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.FanMode == string(message.Payload()) {
				return
			}
			o.states.FanMode = string(message.Payload())
			c.UpdateState()
			o.fanModeCommandFunc(message, client)
		}
	}
	if !reflect.ValueOf(o.fanModeStateTemplate).IsZero() {
		c.FanModeStateTemplate = &o.fanModeStateTemplate
	}
	if !reflect.ValueOf(o.fanModeStateFunc).IsZero() {
		c.fanModeStateFunc = o.fanModeStateFunc
	}
	if !reflect.ValueOf(o.fanModes).IsZero() {
		c.FanModes = &o.fanModes
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		c.Icon = &o.icon
	}
	if !reflect.ValueOf(o.initial).IsZero() {
		c.Initial = &o.initial
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		c.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		c.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.maxHumidity).IsZero() {
		c.MaxHumidity = &o.maxHumidity
	}
	if !reflect.ValueOf(o.maxTemp).IsZero() {
		c.MaxTemp = &o.maxTemp
	}
	if !reflect.ValueOf(o.minHumidity).IsZero() {
		c.MinHumidity = &o.minHumidity
	}
	if !reflect.ValueOf(o.minTemp).IsZero() {
		c.MinTemp = &o.minTemp
	}
	if !reflect.ValueOf(o.modeCommandTemplate).IsZero() {
		c.ModeCommandTemplate = &o.modeCommandTemplate
	}
	if !reflect.ValueOf(o.modeCommandFunc).IsZero() {
		c.modeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.Mode == string(message.Payload()) {
				return
			}
			o.states.Mode = string(message.Payload())
			c.UpdateState()
			o.modeCommandFunc(message, client)
		}
	}
	if !reflect.ValueOf(o.modeStateTemplate).IsZero() {
		c.ModeStateTemplate = &o.modeStateTemplate
	}
	if !reflect.ValueOf(o.modeStateFunc).IsZero() {
		c.modeStateFunc = o.modeStateFunc
	}
	if !reflect.ValueOf(o.modes).IsZero() {
		c.Modes = &o.modes
	}
	if !reflect.ValueOf(o.name).IsZero() {
		c.Name = &o.name
	} else {
		return nil, fmt.Errorf("name not defined")
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		c.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.optimistic).IsZero() {
		c.Optimistic = &o.optimistic
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		c.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		c.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadOff).IsZero() {
		c.PayloadOff = &o.payloadOff
	}
	if !reflect.ValueOf(o.payloadOn).IsZero() {
		c.PayloadOn = &o.payloadOn
	}
	if !reflect.ValueOf(o.powerCommandTemplate).IsZero() {
		c.PowerCommandTemplate = &o.powerCommandTemplate
	}
	if !reflect.ValueOf(o.powerCommandFunc).IsZero() {
		c.powerCommandFunc = o.powerCommandFunc
	}
	if !reflect.ValueOf(o.precision).IsZero() {
		c.Precision = &o.precision
	}
	if !reflect.ValueOf(o.presetModeCommandTemplate).IsZero() {
		c.PresetModeCommandTemplate = &o.presetModeCommandTemplate
	}
	if !reflect.ValueOf(o.presetModeCommandFunc).IsZero() {
		c.presetModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.PresetMode == string(message.Payload()) {
				return
			}
			o.states.PresetMode = string(message.Payload())
			c.UpdateState()
			o.presetModeCommandFunc(message, client)
		}
	}
	if !reflect.ValueOf(o.presetModeStateFunc).IsZero() {
		c.presetModeStateFunc = o.presetModeStateFunc
	}
	if !reflect.ValueOf(o.presetModeValueTemplate).IsZero() {
		c.PresetModeValueTemplate = &o.presetModeValueTemplate
	}
	if !reflect.ValueOf(o.presetModes).IsZero() {
		c.PresetModes = &o.presetModes
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		c.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		c.Retain = &o.retain
	}
	if !reflect.ValueOf(o.swingHorizontalModeCommandTemplate).IsZero() {
		c.SwingHorizontalModeCommandTemplate = &o.swingHorizontalModeCommandTemplate
	}
	if !reflect.ValueOf(o.swingHorizontalModeCommandFunc).IsZero() {
		c.swingHorizontalModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.SwingHorizontalMode == string(message.Payload()) {
				return
			}
			o.states.SwingHorizontalMode = string(message.Payload())
			c.UpdateState()
			o.swingHorizontalModeCommandFunc(message, client)
		}
	}
	if !reflect.ValueOf(o.swingHorizontalModeStateTemplate).IsZero() {
		c.SwingHorizontalModeStateTemplate = &o.swingHorizontalModeStateTemplate
	}
	if !reflect.ValueOf(o.swingHorizontalModeStateFunc).IsZero() {
		c.swingHorizontalModeStateFunc = o.swingHorizontalModeStateFunc
	}
	if !reflect.ValueOf(o.swingHorizontalModes).IsZero() {
		c.SwingHorizontalModes = &o.swingHorizontalModes
	}
	if !reflect.ValueOf(o.swingModeCommandTemplate).IsZero() {
		c.SwingModeCommandTemplate = &o.swingModeCommandTemplate
	}
	if !reflect.ValueOf(o.swingModeCommandFunc).IsZero() {
		c.swingModeCommandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.SwingMode == string(message.Payload()) {
				return
			}
			o.states.SwingMode = string(message.Payload())
			c.UpdateState()
			o.swingModeCommandFunc(message, client)
		}
	}
	if !reflect.ValueOf(o.swingModeStateTemplate).IsZero() {
		c.SwingModeStateTemplate = &o.swingModeStateTemplate
	}
	if !reflect.ValueOf(o.swingModeStateFunc).IsZero() {
		c.swingModeStateFunc = o.swingModeStateFunc
	}
	if !reflect.ValueOf(o.swingModes).IsZero() {
		c.SwingModes = &o.swingModes
	}
	if !reflect.ValueOf(o.targetHumidityCommandTemplate).IsZero() {
		c.TargetHumidityCommandTemplate = &o.targetHumidityCommandTemplate
	}
	if !reflect.ValueOf(o.targetHumidityCommandFunc).IsZero() {
		c.targetHumidityCommandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.TargetHumidity == string(message.Payload()) {
				return
			}
			o.states.TargetHumidity = string(message.Payload())
			c.UpdateState()
			o.targetHumidityCommandFunc(message, client)
		}
	}
	if !reflect.ValueOf(o.targetHumidityStateTemplate).IsZero() {
		c.TargetHumidityStateTemplate = &o.targetHumidityStateTemplate
	}
	if !reflect.ValueOf(o.targetHumidityStateFunc).IsZero() {
		c.targetHumidityStateFunc = o.targetHumidityStateFunc
	}
	if !reflect.ValueOf(o.tempStep).IsZero() {
		c.TempStep = &o.tempStep
	}
	if !reflect.ValueOf(o.temperatureCommandTemplate).IsZero() {
		c.TemperatureCommandTemplate = &o.temperatureCommandTemplate
	}
	if !reflect.ValueOf(o.temperatureCommandFunc).IsZero() {
		c.temperatureCommandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.Temperature == string(message.Payload()) {
				return
			}
			o.states.Temperature = string(message.Payload())
			c.UpdateState()
			o.temperatureCommandFunc(message, client)
		}
	}
	if !reflect.ValueOf(o.temperatureHighCommandTemplate).IsZero() {
		c.TemperatureHighCommandTemplate = &o.temperatureHighCommandTemplate
	}
	if !reflect.ValueOf(o.temperatureHighCommandFunc).IsZero() {
		c.temperatureHighCommandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.TemperatureHigh == string(message.Payload()) {
				return
			}
			o.states.TemperatureHigh = string(message.Payload())
			c.UpdateState()
			o.temperatureHighCommandFunc(message, client)
		}
	}
	if !reflect.ValueOf(o.temperatureHighStateTemplate).IsZero() {
		c.TemperatureHighStateTemplate = &o.temperatureHighStateTemplate
	}
	if !reflect.ValueOf(o.temperatureHighStateFunc).IsZero() {
		c.temperatureHighStateFunc = o.temperatureHighStateFunc
	}
	if !reflect.ValueOf(o.temperatureLowCommandTemplate).IsZero() {
		c.TemperatureLowCommandTemplate = &o.temperatureLowCommandTemplate
	}
	if !reflect.ValueOf(o.temperatureLowCommandFunc).IsZero() {
		c.temperatureLowCommandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.TemperatureLow == string(message.Payload()) {
				return
			}
			o.states.TemperatureLow = string(message.Payload())
			c.UpdateState()
			o.temperatureLowCommandFunc(message, client)
		}
	}
	if !reflect.ValueOf(o.temperatureLowStateTemplate).IsZero() {
		c.TemperatureLowStateTemplate = &o.temperatureLowStateTemplate
	}
	if !reflect.ValueOf(o.temperatureLowStateFunc).IsZero() {
		c.temperatureLowStateFunc = o.temperatureLowStateFunc
	}
	if !reflect.ValueOf(o.temperatureStateTemplate).IsZero() {
		c.TemperatureStateTemplate = &o.temperatureStateTemplate
	}
	if !reflect.ValueOf(o.temperatureStateFunc).IsZero() {
		c.temperatureStateFunc = o.temperatureStateFunc
	}
	if !reflect.ValueOf(o.temperatureUnit).IsZero() {
		c.TemperatureUnit = &o.temperatureUnit
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		c.UniqueId = &o.uniqueId
	} else {
		uniqueId := strcase.ToDelimited(o.name, uint8(0x2d))
		uniqueId = strings.ReplaceAll(uniqueId, "'", "_")
		c.UniqueId = &uniqueId
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		c.ValueTemplate = &o.valueTemplate
	}
	return &c, nil
}

type climateState struct {
	Action              *string
	Availability        *string
	CurrentHumidity     *string
	CurrentTemperature  *string
	FanMode             *string
	JsonAttributes      *string
	Mode                *string
	PresetMode          *string
	SwingHorizontalMode *string
	SwingMode           *string
	TargetHumidity      *string
	TemperatureHigh     *string
	TemperatureLow      *string
	Temperature         *string
}
type ClimateState struct {
	Action              string
	CurrentHumidity     string
	CurrentTemperature  string
	FanMode             string
	JsonAttributes      string
	Mode                string
	PresetMode          string
	SwingHorizontalMode string
	SwingMode           string
	TargetHumidity      string
	TemperatureHigh     string
	TemperatureLow      string
	Temperature         string
}

func (d *Climate) Action(s string) {
	d.States.Action = s
	d.UpdateState()
}
func (d *Climate) CurrentHumidity(s string) {
	d.States.CurrentHumidity = s
	d.UpdateState()
}
func (d *Climate) CurrentTemperature(s string) {
	d.States.CurrentTemperature = s
	d.UpdateState()
}
func (d *Climate) FanMode(s string) {
	d.States.FanMode = s
	d.UpdateState()
}
func (d *Climate) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Climate) Mode(s string) {
	d.States.Mode = s
	d.UpdateState()
}
func (d *Climate) PresetMode(s string) {
	d.States.PresetMode = s
	d.UpdateState()
}
func (d *Climate) SwingHorizontalMode(s string) {
	d.States.SwingHorizontalMode = s
	d.UpdateState()
}
func (d *Climate) SwingMode(s string) {
	d.States.SwingMode = s
	d.UpdateState()
}
func (d *Climate) TargetHumidity(s string) {
	d.States.TargetHumidity = s
	d.UpdateState()
}
func (d *Climate) TemperatureHigh(s string) {
	d.States.TemperatureHigh = s
	d.UpdateState()
}
func (d *Climate) TemperatureLow(s string) {
	d.States.TemperatureLow = s
	d.UpdateState()
}
func (d *Climate) Temperature(s string) {
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
func (d *Climate) GetDomainEntity() string {
	return fmt.Sprintf("climate.%s", strings.ReplaceAll(*d.UniqueId, "-", "_"))
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
	if d.ActionTopic != nil {
		state := d.actionFunc()
		if d.states.Action == nil || state != *d.states.Action || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ActionTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Action = &state
		}
	}
	if d.AvailabilityTopic != nil {
		state := d.availabilityFunc()
		if d.states.Availability == nil || state != *d.states.Availability || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
		}
	}
	if d.CurrentHumidityTopic != nil {
		state := d.currentHumidityFunc()
		if d.states.CurrentHumidity == nil || state != *d.states.CurrentHumidity || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.CurrentHumidityTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.CurrentHumidity = &state
		}
	}
	if d.CurrentTemperatureTopic != nil {
		state := d.currentTemperatureFunc()
		if d.states.CurrentTemperature == nil || state != *d.states.CurrentTemperature || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.CurrentTemperatureTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.CurrentTemperature = &state
		}
	}
	if d.FanModeStateTopic != nil {
		state := d.fanModeStateFunc()
		if d.states.FanMode == nil || state != *d.states.FanMode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.FanModeStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.FanMode = &state
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
	if d.ModeStateTopic != nil {
		state := d.modeStateFunc()
		if d.states.Mode == nil || state != *d.states.Mode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ModeStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Mode = &state
		}
	}
	if d.PresetModeStateTopic != nil {
		state := d.presetModeStateFunc()
		if d.states.PresetMode == nil || state != *d.states.PresetMode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.PresetModeStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.PresetMode = &state
		}
	}
	if d.SwingHorizontalModeStateTopic != nil {
		state := d.swingHorizontalModeStateFunc()
		if d.states.SwingHorizontalMode == nil || state != *d.states.SwingHorizontalMode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.SwingHorizontalModeStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.SwingHorizontalMode = &state
		}
	}
	if d.SwingModeStateTopic != nil {
		state := d.swingModeStateFunc()
		if d.states.SwingMode == nil || state != *d.states.SwingMode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.SwingModeStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.SwingMode = &state
		}
	}
	if d.TargetHumidityStateTopic != nil {
		state := d.targetHumidityStateFunc()
		if d.states.TargetHumidity == nil || state != *d.states.TargetHumidity || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TargetHumidityStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.TargetHumidity = &state
		}
	}
	if d.TemperatureHighStateTopic != nil {
		state := d.temperatureHighStateFunc()
		if d.states.TemperatureHigh == nil || state != *d.states.TemperatureHigh || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TemperatureHighStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.TemperatureHigh = &state
		}
	}
	if d.TemperatureLowStateTopic != nil {
		state := d.temperatureLowStateFunc()
		if d.states.TemperatureLow == nil || state != *d.states.TemperatureLow || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TemperatureLowStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.TemperatureLow = &state
		}
	}
	if d.TemperatureStateTopic != nil {
		state := d.temperatureStateFunc()
		if d.states.Temperature == nil || state != *d.states.Temperature || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TemperatureStateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Temperature = &state
		}
	}
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
	if d.powerCommandFunc != nil {
		d.PowerCommandTopic = new(string)
		*d.PowerCommandTopic = GetTopic(d, "power_command_topic")
		common.TopicStore[*d.PowerCommandTopic] = &d.powerCommandFunc
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
	if d.swingHorizontalModeCommandFunc != nil {
		d.SwingHorizontalModeCommandTopic = new(string)
		*d.SwingHorizontalModeCommandTopic = GetTopic(d, "swing_horizontal_mode_command_topic")
		common.TopicStore[*d.SwingHorizontalModeCommandTopic] = &d.swingHorizontalModeCommandFunc
	}
	if d.swingHorizontalModeStateFunc != nil {
		d.SwingHorizontalModeStateTopic = new(string)
		*d.SwingHorizontalModeStateTopic = GetTopic(d, "swing_horizontal_mode_state_topic")
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
