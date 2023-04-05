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
type AlarmControlPanel struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	Code                   *string `json:"code,omitempty"`                  // "If defined, specifies a code to enable or disable the alarm in the frontend. Note that the code is validated locally and blocks sending MQTT messages to the remote device. For remote code validation, the code can be configured to either of the special values `REMOTE_CODE` (numeric code) or `REMOTE_CODE_TEXT` (text code). In this case, local code validation is bypassed but the frontend will still show a numeric or text code dialog. Use `command_template` to send the code to the remote device. Example configurations for remote code validation [can be found here](./#configurations-with-remote-code-validation)."
	CodeArmRequired        *bool   `json:"code_arm_required,omitempty"`     // "If true the code is required to arm the alarm. If false the code is not validated."
	CodeDisarmRequired     *bool   `json:"code_disarm_required,omitempty"`  // "If true the code is required to disarm the alarm. If false the code is not validated."
	CodeTriggerRequired    *bool   `json:"code_trigger_required,omitempty"` // "If true the code is required to trigger the alarm. If false the code is not validated."
	CommandTemplate        *string `json:"command_template,omitempty"`      // "The [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) used for the command payload. Available variables: `action` and `code`."
	CommandTopic           *string `json:"command_topic,omitempty"`         // "The MQTT topic to publish commands to change the alarm state."
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                      // "The name of the alarm."
	ObjectId               *string `json:"object_id,omitempty"`                 // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadArmAway         *string `json:"payload_arm_away,omitempty"`          // "The payload to set armed-away mode on your Alarm Panel."
	PayloadArmCustomBypass *string `json:"payload_arm_custom_bypass,omitempty"` // "The payload to set armed-custom-bypass mode on your Alarm Panel."
	PayloadArmHome         *string `json:"payload_arm_home,omitempty"`          // "The payload to set armed-home mode on your Alarm Panel."
	PayloadArmNight        *string `json:"payload_arm_night,omitempty"`         // "The payload to set armed-night mode on your Alarm Panel."
	PayloadArmVacation     *string `json:"payload_arm_vacation,omitempty"`      // "The payload to set armed-vacation mode on your Alarm Panel."
	PayloadAvailable       *string `json:"payload_available,omitempty"`         // "The payload that represents the available state."
	PayloadDisarm          *string `json:"payload_disarm,omitempty"`            // "The payload to disarm your Alarm Panel."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"`     // "The payload that represents the unavailable state."
	PayloadTrigger         *string `json:"payload_trigger,omitempty"`           // "The payload to trigger the alarm on your Alarm Panel."
	Qos                    *int    `json:"qos,omitempty"`                       // "The maximum QoS level of the state topic."
	Retain                 *bool   `json:"retain,omitempty"`                    // "If the published message should have the retain flag on or not."
	StateTopic             *string `json:"state_topic,omitempty"`               // "The MQTT topic subscribed to receive state updates."
	stateFunc              func() string
	UniqueId               *string                 `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string                 `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value."
	MQTT                   *MQTTFields             `json:"-"`                        // MQTT configuration parameters
	states                 alarmControlPanelState  // Internal Holder of States
	States                 *AlarmControlPanelState `json:"-"` // External state update location
}

func NewAlarmControlPanel(o *AlarmControlPanelOptions) *AlarmControlPanel {
	var a AlarmControlPanel

	a.States = &o.States
	if !reflect.ValueOf(o.AvailabilityMode).IsZero() {
		a.AvailabilityMode = &o.AvailabilityMode
	}
	if !reflect.ValueOf(o.AvailabilityTemplate).IsZero() {
		a.AvailabilityTemplate = &o.AvailabilityTemplate
	}
	if !reflect.ValueOf(o.AvailabilityFunc).IsZero() {
		a.availabilityFunc = o.AvailabilityFunc
	}
	if !reflect.ValueOf(o.Code).IsZero() {
		a.Code = &o.Code
	}
	if !reflect.ValueOf(o.CodeArmRequired).IsZero() {
		a.CodeArmRequired = &o.CodeArmRequired
	}
	if !reflect.ValueOf(o.CodeDisarmRequired).IsZero() {
		a.CodeDisarmRequired = &o.CodeDisarmRequired
	}
	if !reflect.ValueOf(o.CodeTriggerRequired).IsZero() {
		a.CodeTriggerRequired = &o.CodeTriggerRequired
	}
	if !reflect.ValueOf(o.CommandTemplate).IsZero() {
		a.CommandTemplate = &o.CommandTemplate
	}
	if !reflect.ValueOf(o.CommandFunc).IsZero() {
		a.commandFunc = o.CommandFunc
	} else {
		a.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.States.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.EnabledByDefault).IsZero() {
		a.EnabledByDefault = &o.EnabledByDefault
	}
	if !reflect.ValueOf(o.Encoding).IsZero() {
		a.Encoding = &o.Encoding
	}
	if !reflect.ValueOf(o.EntityCategory).IsZero() {
		a.EntityCategory = &o.EntityCategory
	}
	if !reflect.ValueOf(o.Icon).IsZero() {
		a.Icon = &o.Icon
	}
	if !reflect.ValueOf(o.JsonAttributesTemplate).IsZero() {
		a.JsonAttributesTemplate = &o.JsonAttributesTemplate
	}
	if !reflect.ValueOf(o.JsonAttributesFunc).IsZero() {
		a.jsonAttributesFunc = o.JsonAttributesFunc
	}
	if !reflect.ValueOf(o.Name).IsZero() {
		a.Name = &o.Name
	}
	if !reflect.ValueOf(o.ObjectId).IsZero() {
		a.ObjectId = &o.ObjectId
	}
	if !reflect.ValueOf(o.PayloadArmAway).IsZero() {
		a.PayloadArmAway = &o.PayloadArmAway
	}
	if !reflect.ValueOf(o.PayloadArmCustomBypass).IsZero() {
		a.PayloadArmCustomBypass = &o.PayloadArmCustomBypass
	}
	if !reflect.ValueOf(o.PayloadArmHome).IsZero() {
		a.PayloadArmHome = &o.PayloadArmHome
	}
	if !reflect.ValueOf(o.PayloadArmNight).IsZero() {
		a.PayloadArmNight = &o.PayloadArmNight
	}
	if !reflect.ValueOf(o.PayloadArmVacation).IsZero() {
		a.PayloadArmVacation = &o.PayloadArmVacation
	}
	if !reflect.ValueOf(o.PayloadAvailable).IsZero() {
		a.PayloadAvailable = &o.PayloadAvailable
	}
	if !reflect.ValueOf(o.PayloadDisarm).IsZero() {
		a.PayloadDisarm = &o.PayloadDisarm
	}
	if !reflect.ValueOf(o.PayloadNotAvailable).IsZero() {
		a.PayloadNotAvailable = &o.PayloadNotAvailable
	}
	if !reflect.ValueOf(o.PayloadTrigger).IsZero() {
		a.PayloadTrigger = &o.PayloadTrigger
	}
	if !reflect.ValueOf(o.Qos).IsZero() {
		a.Qos = &o.Qos
	}
	if !reflect.ValueOf(o.Retain).IsZero() {
		a.Retain = &o.Retain
	}
	if !reflect.ValueOf(o.StateFunc).IsZero() {
		a.stateFunc = o.StateFunc
	} else {
		a.stateFunc = func() string {
			return a.States.State
		}
	}
	if !reflect.ValueOf(o.UniqueId).IsZero() {
		a.UniqueId = &o.UniqueId
	}
	if !reflect.ValueOf(o.ValueTemplate).IsZero() {
		a.ValueTemplate = &o.ValueTemplate
	}
	return &a
}

type alarmControlPanelState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type AlarmControlPanelState struct {
	JsonAttributes string
	State          string
}

func (d *AlarmControlPanel) SetJsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *AlarmControlPanel) SetState(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *AlarmControlPanel) GetRawId() string {
	return "alarm_control_panel"
}
func (d *AlarmControlPanel) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *AlarmControlPanel) GetUniqueId() string {
	return *d.UniqueId
}
func (d *AlarmControlPanel) GetName() string {
	return *d.Name
}
func (d *AlarmControlPanel) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *AlarmControlPanel) UpdateState() {
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
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
}
func (d *AlarmControlPanel) Subscribe() {
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
	token := c.Publish(GetDiscoveryTopic(d), 2, true, message)
	token.WaitTimeout(common.WaitTimeout)
	d.availabilityFunc()
	d.UpdateState()
}
func (d *AlarmControlPanel) UnSubscribe() {
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
}
func (d *AlarmControlPanel) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *AlarmControlPanel) Initialize() {
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
func (d *AlarmControlPanel) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *AlarmControlPanel) populateTopics() {
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
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *AlarmControlPanel) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *AlarmControlPanel) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
