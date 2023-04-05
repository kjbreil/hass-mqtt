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
type Lock struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	CodeFormat             *string `json:"code_format,omitempty"`      // "A regular expression to validate a supplied code when it is set during the service call to `open`, `lock` or `unlock` the MQTT lock."
	CommandTemplate        *string `json:"command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`. The lock command template accepts the parameters `value` and `code`. The `value` parameter will contain the configured value for either `payload_open`, `payload_lock` or `payload_unlock`. The `code` parameter is set during the service call to `open`, `lock` or `unlock` the MQTT lock and will be set `None` if no code was passed."
	CommandTopic           *string `json:"command_topic,omitempty"`    // "The MQTT topic to publish commands to change the lock state."
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                  // "The name of the lock."
	ObjectId               *string `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool   `json:"optimistic,omitempty"`            // "Flag that defines if lock works in optimistic mode."
	PayloadAvailable       *string `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadLock            *string `json:"payload_lock,omitempty"`          // "The payload sent to the lock to lock it."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOpen            *string `json:"payload_open,omitempty"`          // "The payload sent to the lock to open it."
	PayloadUnlock          *string `json:"payload_unlock,omitempty"`        // "The payload sent to the lock to unlock it."
	Qos                    *int    `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic. It will also be used for messages published to command topic."
	Retain                 *bool   `json:"retain,omitempty"`                // "If the published message should have the retain flag on or not."
	StateJammed            *string `json:"state_jammed,omitempty"`          // "The payload sent to `state_topic` by the lock when it's jammed."
	StateLocked            *string `json:"state_locked,omitempty"`          // "The payload sent to `state_topic` by the lock when it's locked."
	StateLocking           *string `json:"state_locking,omitempty"`         // "The payload sent to `state_topic` by the lock when it's locking."
	StateTopic             *string `json:"state_topic,omitempty"`           // "The MQTT topic subscribed to receive state updates. It accepts states configured with `state_jammed`, `state_locked`, `state_unlocked`, `state_locking` or `state_unlocking`."
	stateFunc              func() string
	StateUnlocked          *string     `json:"state_unlocked,omitempty"`  // "The payload sent to `state_topic` by the lock when it's unlocked."
	StateUnlocking         *string     `json:"state_unlocking,omitempty"` // "The payload sent to `state_topic` by the lock when it's unlocking."
	UniqueId               *string     `json:"unique_id,omitempty"`       // "An ID that uniquely identifies this lock. If two locks have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a state value from the payload."
	MQTT                   *MQTTFields `json:"-"`                         // MQTT configuration parameters
	states                 lockState   // Internal Holder of States
	States                 *LockState  `json:"-"` // External state update location
}

func NewLock(o *LockOptions) *Lock {
	var l Lock

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
	if !reflect.ValueOf(o.CodeFormat).IsZero() {
		l.CodeFormat = &o.CodeFormat
	}
	if !reflect.ValueOf(o.CommandTemplate).IsZero() {
		l.CommandTemplate = &o.CommandTemplate
	}
	if !reflect.ValueOf(o.CommandFunc).IsZero() {
		l.commandFunc = o.CommandFunc
	} else {
		l.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.States.State = string(message.Payload())
		}
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
	if !reflect.ValueOf(o.Icon).IsZero() {
		l.Icon = &o.Icon
	}
	if !reflect.ValueOf(o.JsonAttributesTemplate).IsZero() {
		l.JsonAttributesTemplate = &o.JsonAttributesTemplate
	}
	if !reflect.ValueOf(o.JsonAttributesFunc).IsZero() {
		l.jsonAttributesFunc = o.JsonAttributesFunc
	}
	if !reflect.ValueOf(o.Name).IsZero() {
		l.Name = &o.Name
	}
	if !reflect.ValueOf(o.ObjectId).IsZero() {
		l.ObjectId = &o.ObjectId
	}
	if !reflect.ValueOf(o.Optimistic).IsZero() {
		l.Optimistic = &o.Optimistic
	}
	if !reflect.ValueOf(o.PayloadAvailable).IsZero() {
		l.PayloadAvailable = &o.PayloadAvailable
	}
	if !reflect.ValueOf(o.PayloadLock).IsZero() {
		l.PayloadLock = &o.PayloadLock
	}
	if !reflect.ValueOf(o.PayloadNotAvailable).IsZero() {
		l.PayloadNotAvailable = &o.PayloadNotAvailable
	}
	if !reflect.ValueOf(o.PayloadOpen).IsZero() {
		l.PayloadOpen = &o.PayloadOpen
	}
	if !reflect.ValueOf(o.PayloadUnlock).IsZero() {
		l.PayloadUnlock = &o.PayloadUnlock
	}
	if !reflect.ValueOf(o.Qos).IsZero() {
		l.Qos = &o.Qos
	}
	if !reflect.ValueOf(o.Retain).IsZero() {
		l.Retain = &o.Retain
	}
	if !reflect.ValueOf(o.StateJammed).IsZero() {
		l.StateJammed = &o.StateJammed
	}
	if !reflect.ValueOf(o.StateLocked).IsZero() {
		l.StateLocked = &o.StateLocked
	}
	if !reflect.ValueOf(o.StateLocking).IsZero() {
		l.StateLocking = &o.StateLocking
	}
	if !reflect.ValueOf(o.StateFunc).IsZero() {
		l.stateFunc = o.StateFunc
	} else {
		l.stateFunc = func() string {
			return l.States.State
		}
	}
	if !reflect.ValueOf(o.StateUnlocked).IsZero() {
		l.StateUnlocked = &o.StateUnlocked
	}
	if !reflect.ValueOf(o.StateUnlocking).IsZero() {
		l.StateUnlocking = &o.StateUnlocking
	}
	if !reflect.ValueOf(o.UniqueId).IsZero() {
		l.UniqueId = &o.UniqueId
	}
	if !reflect.ValueOf(o.ValueTemplate).IsZero() {
		l.ValueTemplate = &o.ValueTemplate
	}
	return &l
}

type lockState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type LockState struct {
	JsonAttributes string
	State          string
}

func (d *Lock) SetJsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Lock) SetState(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Lock) GetRawId() string {
	return "lock"
}
func (d *Lock) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Lock) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Lock) GetName() string {
	return *d.Name
}
func (d *Lock) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Lock) UpdateState() {
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
func (d *Lock) Subscribe() {
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
func (d *Lock) UnSubscribe() {
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
func (d *Lock) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Lock) Initialize() {
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
func (d *Lock) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Lock) populateTopics() {
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
func (d *Lock) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Lock) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
