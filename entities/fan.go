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
type Fan struct {
	AvailabilityMode           *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate       *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic          *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc           func() string
	CommandTemplate            *string `json:"command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandTopic               *string `json:"command_topic,omitempty"`    // "The MQTT topic to publish commands to change the fan state."
	commandFunc                func(mqtt.Message, mqtt.Client)
	Device                     Device  `json:"device,omitempty"`                   // Device configuration parameters
	EnabledByDefault           *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding                   *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory             *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                       *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate     *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic        *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc         func() string
	Name                       *string `json:"name,omitempty"`                         // "The name of the fan."
	ObjectId                   *string `json:"object_id,omitempty"`                    // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                 *bool   `json:"optimistic,omitempty"`                   // "Flag that defines if fan works in optimistic mode"
	OscillationCommandTemplate *string `json:"oscillation_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `oscillation_command_topic`."
	OscillationCommandTopic    *string `json:"oscillation_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the oscillation state."
	oscillationCommandFunc     func(mqtt.Message, mqtt.Client)
	OscillationStateTopic      *string `json:"oscillation_state_topic,omitempty"` // "The MQTT topic subscribed to receive oscillation state updates."
	oscillationStateFunc       func() string
	OscillationValueTemplate   *string `json:"oscillation_value_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the oscillation."
	PayloadAvailable           *string `json:"payload_available,omitempty"`           // "The payload that represents the available state."
	PayloadNotAvailable        *string `json:"payload_not_available,omitempty"`       // "The payload that represents the unavailable state."
	PayloadOff                 *string `json:"payload_off,omitempty"`                 // "The payload that represents the stop state."
	PayloadOn                  *string `json:"payload_on,omitempty"`                  // "The payload that represents the running state."
	PayloadOscillationOff      *string `json:"payload_oscillation_off,omitempty"`     // "The payload that represents the oscillation off state."
	PayloadOscillationOn       *string `json:"payload_oscillation_on,omitempty"`      // "The payload that represents the oscillation on state."
	PayloadResetPercentage     *string `json:"payload_reset_percentage,omitempty"`    // "A special payload that resets the `percentage` state attribute to `None` when received at the `percentage_state_topic`."
	PayloadResetPresetMode     *string `json:"payload_reset_preset_mode,omitempty"`   // "A special payload that resets the `preset_mode` state attribute to `None` when received at the `preset_mode_state_topic`."
	PercentageCommandTemplate  *string `json:"percentage_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `percentage_command_topic`."
	PercentageCommandTopic     *string `json:"percentage_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the fan speed state based on a percentage."
	percentageCommandFunc      func(mqtt.Message, mqtt.Client)
	PercentageStateTopic       *string `json:"percentage_state_topic,omitempty"` // "The MQTT topic subscribed to receive fan speed based on percentage."
	percentageStateFunc        func() string
	PercentageValueTemplate    *string `json:"percentage_value_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `percentage` value from the payload received on `percentage_state_topic`."
	PresetModeCommandTemplate  *string `json:"preset_mode_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommandTopic     *string `json:"preset_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the preset mode."
	presetModeCommandFunc      func(mqtt.Message, mqtt.Client)
	PresetModeStateTopic       *string `json:"preset_mode_state_topic,omitempty"` // "The MQTT topic subscribed to receive fan speed based on presets."
	presetModeStateFunc        func() string
	PresetModeValueTemplate    *string     `json:"preset_mode_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                *([]string) `json:"preset_modes,omitempty"`               // "List of preset modes this fan is capable of running at. Common examples include `auto`, `smart`, `whoosh`, `eco` and `breeze`."
	Qos                        *int        `json:"qos,omitempty"`                        // "The maximum QoS level of the state topic."
	Retain                     *bool       `json:"retain,omitempty"`                     // "If the published message should have the retain flag on or not."
	SpeedRangeMax              *int        `json:"speed_range_max,omitempty"`            // "The maximum of numeric output range (representing 100 %). The number of speeds within the `speed_range` / `100` will determine the `percentage_step`."
	SpeedRangeMin              *int        `json:"speed_range_min,omitempty"`            // "The minimum of numeric output range (`off` not included, so `speed_range_min` - `1` represents 0 %). The number of speeds within the speed_range / 100 will determine the `percentage_step`."
	StateTopic                 *string     `json:"state_topic,omitempty"`                // "The MQTT topic subscribed to receive state updates."
	stateFunc                  func() string
	StateValueTemplate         *string     `json:"state_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state."
	UniqueId                   *string     `json:"unique_id,omitempty"`            // "An ID that uniquely identifies this fan. If two fans have the same unique ID, Home Assistant will raise an exception."
	MQTT                       *MQTTFields `json:"-"`                              // MQTT configuration parameters
	states                     fanState    // Internal Holder of States
	States                     *FanState   `json:"-"` // External state update location
}

func NewFan(o *FanOptions) *Fan {
	var f Fan

	f.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		f.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		f.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		f.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.commandTemplate).IsZero() {
		f.CommandTemplate = &o.commandTemplate
	}
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		f.commandFunc = o.commandFunc
	} else {
		f.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		f.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		f.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		f.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		f.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		f.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		f.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.name).IsZero() {
		f.Name = &o.name
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		f.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.optimistic).IsZero() {
		f.Optimistic = &o.optimistic
	}
	if !reflect.ValueOf(o.oscillationCommandTemplate).IsZero() {
		f.OscillationCommandTemplate = &o.oscillationCommandTemplate
	}
	if !reflect.ValueOf(o.oscillationCommandFunc).IsZero() {
		f.oscillationCommandFunc = o.oscillationCommandFunc
	}
	if !reflect.ValueOf(o.oscillationStateFunc).IsZero() {
		f.oscillationStateFunc = o.oscillationStateFunc
	}
	if !reflect.ValueOf(o.oscillationValueTemplate).IsZero() {
		f.OscillationValueTemplate = &o.oscillationValueTemplate
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		f.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		f.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadOff).IsZero() {
		f.PayloadOff = &o.payloadOff
	}
	if !reflect.ValueOf(o.payloadOn).IsZero() {
		f.PayloadOn = &o.payloadOn
	}
	if !reflect.ValueOf(o.payloadOscillationOff).IsZero() {
		f.PayloadOscillationOff = &o.payloadOscillationOff
	}
	if !reflect.ValueOf(o.payloadOscillationOn).IsZero() {
		f.PayloadOscillationOn = &o.payloadOscillationOn
	}
	if !reflect.ValueOf(o.payloadResetPercentage).IsZero() {
		f.PayloadResetPercentage = &o.payloadResetPercentage
	}
	if !reflect.ValueOf(o.payloadResetPresetMode).IsZero() {
		f.PayloadResetPresetMode = &o.payloadResetPresetMode
	}
	if !reflect.ValueOf(o.percentageCommandTemplate).IsZero() {
		f.PercentageCommandTemplate = &o.percentageCommandTemplate
	}
	if !reflect.ValueOf(o.percentageCommandFunc).IsZero() {
		f.percentageCommandFunc = o.percentageCommandFunc
	}
	if !reflect.ValueOf(o.percentageStateFunc).IsZero() {
		f.percentageStateFunc = o.percentageStateFunc
	}
	if !reflect.ValueOf(o.percentageValueTemplate).IsZero() {
		f.PercentageValueTemplate = &o.percentageValueTemplate
	}
	if !reflect.ValueOf(o.presetModeCommandTemplate).IsZero() {
		f.PresetModeCommandTemplate = &o.presetModeCommandTemplate
	}
	if !reflect.ValueOf(o.presetModeCommandFunc).IsZero() {
		f.presetModeCommandFunc = o.presetModeCommandFunc
	}
	if !reflect.ValueOf(o.presetModeStateFunc).IsZero() {
		f.presetModeStateFunc = o.presetModeStateFunc
	}
	if !reflect.ValueOf(o.presetModeValueTemplate).IsZero() {
		f.PresetModeValueTemplate = &o.presetModeValueTemplate
	}
	if !reflect.ValueOf(o.presetModes).IsZero() {
		f.PresetModes = &o.presetModes
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		f.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		f.Retain = &o.retain
	}
	if !reflect.ValueOf(o.speedRangeMax).IsZero() {
		f.SpeedRangeMax = &o.speedRangeMax
	}
	if !reflect.ValueOf(o.speedRangeMin).IsZero() {
		f.SpeedRangeMin = &o.speedRangeMin
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		f.stateFunc = o.stateFunc
	} else {
		f.stateFunc = func() string {
			return f.States.State
		}
	}
	if !reflect.ValueOf(o.stateValueTemplate).IsZero() {
		f.StateValueTemplate = &o.stateValueTemplate
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		f.UniqueId = &o.uniqueId
	}
	return &f
}

type fanState struct {
	Availability   *string
	JsonAttributes *string
	Oscillation    *string
	Percentage     *string
	PresetMode     *string
	State          *string
}
type FanState struct {
	JsonAttributes string
	Oscillation    string
	Percentage     string
	PresetMode     string
	State          string
}

func (d *Fan) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Fan) Oscillation(s string) {
	d.States.Oscillation = s
	d.UpdateState()
}
func (d *Fan) Percentage(s string) {
	d.States.Percentage = s
	d.UpdateState()
}
func (d *Fan) PresetMode(s string) {
	d.States.PresetMode = s
	d.UpdateState()
}
func (d *Fan) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Fan) GetRawId() string {
	return "fan"
}
func (d *Fan) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Fan) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Fan) GetName() string {
	return *d.Name
}
func (d *Fan) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Fan) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.availabilityFunc()
		if d.states.Availability == nil || state != *d.states.Availability || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
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
	if d.OscillationStateTopic != nil {
		state := d.oscillationStateFunc()
		if d.states.Oscillation == nil || state != *d.states.Oscillation || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.OscillationStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Oscillation = &state
		}
	}
	if d.PercentageStateTopic != nil {
		state := d.percentageStateFunc()
		if d.states.Percentage == nil || state != *d.states.Percentage || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.PercentageStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Percentage = &state
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
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
}
func (d *Fan) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.CommandTopic != nil {
		t := c.Subscribe(*d.CommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.OscillationCommandTopic != nil {
		t := c.Subscribe(*d.OscillationCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PercentageCommandTopic != nil {
		t := c.Subscribe(*d.PercentageCommandTopic, 0, d.MQTT.MessageHandler)
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
	token := c.Publish(GetDiscoveryTopic(d), 2, true, message)
	token.WaitTimeout(common.WaitTimeout)
	d.availabilityFunc()
	d.UpdateState()
}
func (d *Fan) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, false, "offline")
	token.WaitTimeout(common.WaitTimeout)
	if d.CommandTopic != nil {
		t := c.Unsubscribe(*d.CommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.OscillationCommandTopic != nil {
		t := c.Unsubscribe(*d.OscillationCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PercentageCommandTopic != nil {
		t := c.Unsubscribe(*d.PercentageCommandTopic)
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
}
func (d *Fan) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Fan) Initialize() {
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
func (d *Fan) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Fan) populateTopics() {
	if d.availabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.commandFunc != nil {
		d.CommandTopic = new(string)
		*d.CommandTopic = GetTopic(d, "command_topic")
		common.TopicStore[*d.CommandTopic] = &d.commandFunc
	}
	if d.jsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
	}
	if d.oscillationCommandFunc != nil {
		d.OscillationCommandTopic = new(string)
		*d.OscillationCommandTopic = GetTopic(d, "oscillation_command_topic")
		common.TopicStore[*d.OscillationCommandTopic] = &d.oscillationCommandFunc
	}
	if d.oscillationStateFunc != nil {
		d.OscillationStateTopic = new(string)
		*d.OscillationStateTopic = GetTopic(d, "oscillation_state_topic")
	}
	if d.percentageCommandFunc != nil {
		d.PercentageCommandTopic = new(string)
		*d.PercentageCommandTopic = GetTopic(d, "percentage_command_topic")
		common.TopicStore[*d.PercentageCommandTopic] = &d.percentageCommandFunc
	}
	if d.percentageStateFunc != nil {
		d.PercentageStateTopic = new(string)
		*d.PercentageStateTopic = GetTopic(d, "percentage_state_topic")
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
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Fan) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Fan) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
