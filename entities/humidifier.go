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
type Humidifier struct {
	AvailabilityMode              *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate          *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic             *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc              func() string
	CommandTemplate               *string `json:"command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandTopic                  *string `json:"command_topic,omitempty"`    // "The MQTT topic to publish commands to change the humidifier state."
	commandFunc                   func(mqtt.Message, mqtt.Client)
	Device                        Device  `json:"device,omitempty"`                   // Device configuration parameters
	DeviceClass                   *string `json:"device_class,omitempty"`             // "The device class of the MQTT device. Must be either `humidifier` or `dehumidifier`."
	EnabledByDefault              *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding                      *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                          *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate        *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic           *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc            func() string
	MaxHumidity                   *int    `json:"max_humidity,omitempty"`          // "The minimum target humidity percentage that can be set."
	MinHumidity                   *int    `json:"min_humidity,omitempty"`          // "The maximum target humidity percentage that can be set."
	ModeCommandTemplate           *string `json:"mode_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `mode_command_topic`."
	ModeCommandTopic              *string `json:"mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the `mode` on the humidifier. This attribute ust be configured together with the `modes` attribute."
	modeCommandFunc               func(mqtt.Message, mqtt.Client)
	ModeStateTemplate             *string `json:"mode_state_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `mode` state."
	ModeStateTopic                *string `json:"mode_state_topic,omitempty"`    // "The MQTT topic subscribed to receive the humidifier `mode`."
	modeStateFunc                 func() string
	Modes                         *([]string) `json:"modes,omitempty"`                  // "List of available modes this humidifier is capable of running at. Common examples include `normal`, `eco`, `away`, `boost`, `comfort`, `home`, `sleep`, `auto` and `baby`. These examples offer built-in translations but other custom modes are allowed as well.  This attribute ust be configured together with the `mode_command_topic` attribute."
	Name                          *string     `json:"name,omitempty"`                   // "The name of the humidifier."
	ObjectId                      *string     `json:"object_id,omitempty"`              // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                    *bool       `json:"optimistic,omitempty"`             // "Flag that defines if humidifier works in optimistic mode"
	PayloadAvailable              *string     `json:"payload_available,omitempty"`      // "The payload that represents the available state."
	PayloadNotAvailable           *string     `json:"payload_not_available,omitempty"`  // "The payload that represents the unavailable state."
	PayloadOff                    *string     `json:"payload_off,omitempty"`            // "The payload that represents the stop state."
	PayloadOn                     *string     `json:"payload_on,omitempty"`             // "The payload that represents the running state."
	PayloadResetHumidity          *string     `json:"payload_reset_humidity,omitempty"` // "A special payload that resets the `target_humidity` state attribute to `None` when received at the `target_humidity_state_topic`."
	PayloadResetMode              *string     `json:"payload_reset_mode,omitempty"`     // "A special payload that resets the `mode` state attribute to `None` when received at the `mode_state_topic`."
	Qos                           *int        `json:"qos,omitempty"`                    // "The maximum QoS level of the state topic."
	Retain                        *bool       `json:"retain,omitempty"`                 // "If the published message should have the retain flag on or not."
	StateTopic                    *string     `json:"state_topic,omitempty"`            // "The MQTT topic subscribed to receive state updates."
	stateFunc                     func() string
	StateValueTemplate            *string `json:"state_value_template,omitempty"`             // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state."
	TargetHumidityCommandTemplate *string `json:"target_humidity_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `target_humidity_command_topic`."
	TargetHumidityCommandTopic    *string `json:"target_humidity_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the humidifier target humidity state based on a percentage."
	targetHumidityCommandFunc     func(mqtt.Message, mqtt.Client)
	TargetHumidityStateTemplate   *string `json:"target_humidity_state_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `target_humidity` state."
	TargetHumidityStateTopic      *string `json:"target_humidity_state_topic,omitempty"`    // "The MQTT topic subscribed to receive humidifier target humidity."
	targetHumidityStateFunc       func() string
	UniqueId                      *string          `json:"unique_id,omitempty"` // "An ID that uniquely identifies this humidifier. If two humidifiers have the same unique ID, Home Assistant will raise an exception."
	MQTT                          *MQTTFields      `json:"-"`                   // MQTT configuration parameters
	states                        humidifierState  // Internal Holder of States
	States                        *HumidifierState `json:"-"` // External state update location
}

func NewHumidifier(o *HumidifierOptions) *Humidifier {
	var h Humidifier

	h.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		h.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		h.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		h.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.commandTemplate).IsZero() {
		h.CommandTemplate = &o.commandTemplate
	}
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		h.commandFunc = o.commandFunc
	} else {
		h.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.deviceClass).IsZero() {
		h.DeviceClass = &o.deviceClass
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		h.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		h.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		h.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		h.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		h.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		h.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.maxHumidity).IsZero() {
		h.MaxHumidity = &o.maxHumidity
	}
	if !reflect.ValueOf(o.minHumidity).IsZero() {
		h.MinHumidity = &o.minHumidity
	}
	if !reflect.ValueOf(o.modeCommandTemplate).IsZero() {
		h.ModeCommandTemplate = &o.modeCommandTemplate
	}
	if !reflect.ValueOf(o.modeCommandFunc).IsZero() {
		h.modeCommandFunc = o.modeCommandFunc
	}
	if !reflect.ValueOf(o.modeStateTemplate).IsZero() {
		h.ModeStateTemplate = &o.modeStateTemplate
	}
	if !reflect.ValueOf(o.modeStateFunc).IsZero() {
		h.modeStateFunc = o.modeStateFunc
	}
	if !reflect.ValueOf(o.modes).IsZero() {
		h.Modes = &o.modes
	}
	if !reflect.ValueOf(o.name).IsZero() {
		h.Name = &o.name
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		h.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.optimistic).IsZero() {
		h.Optimistic = &o.optimistic
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		h.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		h.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadOff).IsZero() {
		h.PayloadOff = &o.payloadOff
	}
	if !reflect.ValueOf(o.payloadOn).IsZero() {
		h.PayloadOn = &o.payloadOn
	}
	if !reflect.ValueOf(o.payloadResetHumidity).IsZero() {
		h.PayloadResetHumidity = &o.payloadResetHumidity
	}
	if !reflect.ValueOf(o.payloadResetMode).IsZero() {
		h.PayloadResetMode = &o.payloadResetMode
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		h.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		h.Retain = &o.retain
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		h.stateFunc = o.stateFunc
	} else {
		h.stateFunc = func() string {
			return h.States.State
		}
	}
	if !reflect.ValueOf(o.stateValueTemplate).IsZero() {
		h.StateValueTemplate = &o.stateValueTemplate
	}
	if !reflect.ValueOf(o.targetHumidityCommandTemplate).IsZero() {
		h.TargetHumidityCommandTemplate = &o.targetHumidityCommandTemplate
	}
	if !reflect.ValueOf(o.targetHumidityCommandFunc).IsZero() {
		h.targetHumidityCommandFunc = o.targetHumidityCommandFunc
	}
	if !reflect.ValueOf(o.targetHumidityStateTemplate).IsZero() {
		h.TargetHumidityStateTemplate = &o.targetHumidityStateTemplate
	}
	if !reflect.ValueOf(o.targetHumidityStateFunc).IsZero() {
		h.targetHumidityStateFunc = o.targetHumidityStateFunc
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		h.UniqueId = &o.uniqueId
	}
	return &h
}

type humidifierState struct {
	Availability   *string
	JsonAttributes *string
	Mode           *string
	State          *string
	TargetHumidity *string
}
type HumidifierState struct {
	JsonAttributes string
	Mode           string
	State          string
	TargetHumidity string
}

func (d *Humidifier) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Humidifier) Mode(s string) {
	d.States.Mode = s
	d.UpdateState()
}
func (d *Humidifier) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Humidifier) TargetHumidity(s string) {
	d.States.TargetHumidity = s
	d.UpdateState()
}
func (d *Humidifier) GetRawId() string {
	return "humidifier"
}
func (d *Humidifier) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Humidifier) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Humidifier) GetName() string {
	return *d.Name
}
func (d *Humidifier) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Humidifier) UpdateState() {
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
	if d.ModeStateTopic != nil {
		state := d.modeStateFunc()
		if d.states.Mode == nil || state != *d.states.Mode || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ModeStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Mode = &state
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
	if d.TargetHumidityStateTopic != nil {
		state := d.targetHumidityStateFunc()
		if d.states.TargetHumidity == nil || state != *d.states.TargetHumidity || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TargetHumidityStateTopic, byte(*d.Qos), *d.Retain, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.TargetHumidity = &state
		}
	}
}
func (d *Humidifier) Subscribe() {
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
	if d.ModeCommandTopic != nil {
		t := c.Subscribe(*d.ModeCommandTopic, 0, d.MQTT.MessageHandler)
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
	token := c.Publish(GetDiscoveryTopic(d), 2, true, message)
	token.WaitTimeout(common.WaitTimeout)
	d.availabilityFunc()
	d.UpdateState()
}
func (d *Humidifier) UnSubscribe() {
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
	if d.ModeCommandTopic != nil {
		t := c.Unsubscribe(*d.ModeCommandTopic)
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
}
func (d *Humidifier) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Humidifier) Initialize() {
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
func (d *Humidifier) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Humidifier) populateTopics() {
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
	if d.modeCommandFunc != nil {
		d.ModeCommandTopic = new(string)
		*d.ModeCommandTopic = GetTopic(d, "mode_command_topic")
		common.TopicStore[*d.ModeCommandTopic] = &d.modeCommandFunc
	}
	if d.modeStateFunc != nil {
		d.ModeStateTopic = new(string)
		*d.ModeStateTopic = GetTopic(d, "mode_state_topic")
	}
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
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
}
func (d *Humidifier) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Humidifier) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
