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
type AlarmControlPanel struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	Code                   *string `json:"code,omitempty"`                  // "If defined, specifies a code to enable or disable the alarm in the frontend. Note that the code is validated locally and blocks sending MQTT messages to the remote device. For remote code validation, the code can be configured to either of the special values `REMOTE_CODE` (numeric code) or `REMOTE_CODE_TEXT` (text code). In this case, local code validation is bypassed but the frontend will still show a numeric or text code dialog. Use `command_template` to send the code to the remote device. Example configurations for remote code validation [can be found here](#configurations-with-remote-code-validation)."
	CodeArmRequired        *bool   `json:"code_arm_required,omitempty"`     // "If true the code is required to arm the alarm. If false the code is not validated."
	CodeDisarmRequired     *bool   `json:"code_disarm_required,omitempty"`  // "If true the code is required to disarm the alarm. If false the code is not validated."
	CodeTriggerRequired    *bool   `json:"code_trigger_required,omitempty"` // "If true the code is required to trigger the alarm. If false the code is not validated."
	CommandTemplate        *string `json:"command_template,omitempty"`      // "The [template](/docs/configuration/templating/#using-command-templates-with-mqtt) used for the command payload. Available variables: `action` and `code`."
	CommandTopic           *string `json:"command_topic,omitempty"`         // "The MQTT topic to publish commands to change the alarm state."
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	EntityPicture          *string `json:"entity_picture,omitempty"`           // "Picture URL for the entity."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                      // "The name of the alarm. Can be set to `null` if only the device name is relevant."
	ObjectId               *string `json:"object_id,omitempty"`                 // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	PayloadArmAway         *string `json:"payload_arm_away,omitempty"`          // "The payload to set armed-away mode on your Alarm Panel."
	PayloadArmCustomBypass *string `json:"payload_arm_custom_bypass,omitempty"` // "The payload to set armed-custom-bypass mode on your Alarm Panel."
	PayloadArmHome         *string `json:"payload_arm_home,omitempty"`          // "The payload to set armed-home mode on your Alarm Panel."
	PayloadArmNight        *string `json:"payload_arm_night,omitempty"`         // "The payload to set armed-night mode on your Alarm Panel."
	PayloadArmVacation     *string `json:"payload_arm_vacation,omitempty"`      // "The payload to set armed-vacation mode on your Alarm Panel."
	PayloadAvailable       *string `json:"payload_available,omitempty"`         // "The payload that represents the available state."
	PayloadDisarm          *string `json:"payload_disarm,omitempty"`            // "The payload to disarm your Alarm Panel."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"`     // "The payload that represents the unavailable state."
	PayloadTrigger         *string `json:"payload_trigger,omitempty"`           // "The payload to trigger the alarm on your Alarm Panel."
	Platform               *string `json:"platform,omitempty"`                  // "Must be `alarm_control_panel`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	Qos                    *int    `json:"qos,omitempty"`                       // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                 *bool   `json:"retain,omitempty"`                    // "If the published message should have the retain flag on or not."
	StateTopic             *string `json:"state_topic,omitempty"`               // "The MQTT topic subscribed to receive state updates. A \"None\" payload resets to an `unknown` state. An empty payload is ignored. Valid state payloads are: `armed_away`, `armed_custom_bypass`, `armed_home`, `armed_night`, `armed_vacation`, `arming`, `disarmed`, `disarming` `pending` and `triggered`."
	stateFunc              func() string
	SupportedFeatures      *([]string)             `json:"supported_features,omitempty"` // "A list of features that the alarm control panel supports. The available list options are `arm_home`, `arm_away`, `arm_night`, `arm_vacation`, `arm_custom_bypass`, and `trigger`."
	UniqueId               *string                 `json:"unique_id,omitempty"`          // "An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	ValueTemplate          *string                 `json:"value_template,omitempty"`     // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the value."
	MQTT                   *MQTTFields             `json:"-"`                            // MQTT configuration parameters
	states                 alarmControlPanelState  // Internal Holder of States
	States                 *AlarmControlPanelState `json:"-"` // External state update location
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
func NewAlarmControlPanel(o *AlarmControlPanelOptions) (*AlarmControlPanel, error) {
	var a AlarmControlPanel

	a.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		a.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		a.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		a.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.code).IsZero() {
		a.Code = &o.code
	}
	if !reflect.ValueOf(o.codeArmRequired).IsZero() {
		a.CodeArmRequired = &o.codeArmRequired
	}
	if !reflect.ValueOf(o.codeDisarmRequired).IsZero() {
		a.CodeDisarmRequired = &o.codeDisarmRequired
	}
	if !reflect.ValueOf(o.codeTriggerRequired).IsZero() {
		a.CodeTriggerRequired = &o.codeTriggerRequired
	}
	if !reflect.ValueOf(o.commandTemplate).IsZero() {
		a.CommandTemplate = &o.commandTemplate
	}
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		a.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.State == string(message.Payload()) {
				return
			}
			o.states.State = string(message.Payload())
			a.UpdateState()
			o.commandFunc(message, client)
		}
	} else {
		a.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		a.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		a.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		a.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.entityPicture).IsZero() {
		a.EntityPicture = &o.entityPicture
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		a.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		a.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		a.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.name).IsZero() {
		a.Name = &o.name
	} else {
		return nil, fmt.Errorf("name not defined")
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		a.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.payloadArmAway).IsZero() {
		a.PayloadArmAway = &o.payloadArmAway
	}
	if !reflect.ValueOf(o.payloadArmCustomBypass).IsZero() {
		a.PayloadArmCustomBypass = &o.payloadArmCustomBypass
	}
	if !reflect.ValueOf(o.payloadArmHome).IsZero() {
		a.PayloadArmHome = &o.payloadArmHome
	}
	if !reflect.ValueOf(o.payloadArmNight).IsZero() {
		a.PayloadArmNight = &o.payloadArmNight
	}
	if !reflect.ValueOf(o.payloadArmVacation).IsZero() {
		a.PayloadArmVacation = &o.payloadArmVacation
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		a.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadDisarm).IsZero() {
		a.PayloadDisarm = &o.payloadDisarm
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		a.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadTrigger).IsZero() {
		a.PayloadTrigger = &o.payloadTrigger
	}
	if !reflect.ValueOf(o.platform).IsZero() {
		a.Platform = &o.platform
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		a.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		a.Retain = &o.retain
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		a.stateFunc = o.stateFunc
	} else {
		a.stateFunc = func() string {
			return a.States.State
		}
	}
	if !reflect.ValueOf(o.supportedFeatures).IsZero() {
		a.SupportedFeatures = &o.supportedFeatures
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		a.UniqueId = &o.uniqueId
	} else {
		uniqueId := strcase.ToDelimited(o.name, uint8(0x2d))
		uniqueId = strings.ReplaceAll(uniqueId, "'", "_")
		a.UniqueId = &uniqueId
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		a.ValueTemplate = &o.valueTemplate
	}
	return &a, nil
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

func (d *AlarmControlPanel) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *AlarmControlPanel) State(s string) {
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
func (d *AlarmControlPanel) GetDomainEntity() string {
	return fmt.Sprintf("alarm_control_panel.%s", strings.ReplaceAll(*d.UniqueId, "-", "_"))
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
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Availability = &state
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
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
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
