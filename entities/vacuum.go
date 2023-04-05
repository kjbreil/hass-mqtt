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
type Vacuum struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	CommandTopic           *string `json:"command_topic,omitempty"` // "The MQTT topic to publish commands to control the vacuum."
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device      `json:"device,omitempty"`                   // Device configuration parameters
	Encoding               *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	FanSpeedList           *([]string) `json:"fan_speed_list,omitempty"`           // "List of possible fan speeds for the vacuum."
	JsonAttributesTemplate *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string     `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                   // "The name of the vacuum."
	ObjectId               *string `json:"object_id,omitempty"`              // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       *string `json:"payload_available,omitempty"`      // "The payload that represents the available state."
	PayloadCleanSpot       *string `json:"payload_clean_spot,omitempty"`     // "The payload to send to the `command_topic` to begin a spot cleaning cycle."
	PayloadLocate          *string `json:"payload_locate,omitempty"`         // "The payload to send to the `command_topic` to locate the vacuum (typically plays a song)."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"`  // "The payload that represents the unavailable state."
	PayloadPause           *string `json:"payload_pause,omitempty"`          // "The payload to send to the `command_topic` to pause the vacuum."
	PayloadReturnToBase    *string `json:"payload_return_to_base,omitempty"` // "The payload to send to the `command_topic` to tell the vacuum to return to base."
	PayloadStart           *string `json:"payload_start,omitempty"`          // "The payload to send to the `command_topic` to begin the cleaning cycle."
	PayloadStop            *string `json:"payload_stop,omitempty"`           // "The payload to send to the `command_topic` to stop cleaning."
	Qos                    *int    `json:"qos,omitempty"`                    // "The maximum QoS level of the state topic."
	Retain                 *bool   `json:"retain,omitempty"`                 // "If the published message should have the retain flag on or not."
	Schema                 *string `json:"schema,omitempty"`                 // "The schema to use. Must be `state` to select the state schema."
	SendCommandTopic       *string `json:"send_command_topic,omitempty"`     // "The MQTT topic to publish custom commands to the vacuum."
	sendCommandFunc        func(mqtt.Message, mqtt.Client)
	SetFanSpeedTopic       *string `json:"set_fan_speed_topic,omitempty"` // "The MQTT topic to publish commands to control the vacuum's fan speed."
	setFanSpeedFunc        func(mqtt.Message, mqtt.Client)
	StateTopic             *string `json:"state_topic,omitempty"` // "The MQTT topic subscribed to receive state messages from the vacuum. Messages received on the `state_topic` must be a valid JSON dictionary, with a mandatory `state` key and optionally `battery_level` and `fan_speed` keys as shown in the [example](#state-mqtt-protocol)."
	stateFunc              func() string
	SupportedFeatures      *([]string)  `json:"supported_features,omitempty"` // "List of features that the vacuum supports (possible values are `start`, `stop`, `pause`, `return_home`, `battery`, `status`, `locate`, `clean_spot`, `fan_speed`, `send_command`)."
	UniqueId               *string      `json:"unique_id,omitempty"`          // "An ID that uniquely identifies this vacuum. If two vacuums have the same unique ID, Home Assistant will raise an exception."
	MQTT                   *MQTTFields  `json:"-"`                            // MQTT configuration parameters
	states                 vacuumState  // Internal Holder of States
	States                 *VacuumState `json:"-"` // External state update location
}

func NewVacuum(o *VacuumOptions) *Vacuum {
	var v Vacuum

	v.States = &o.States
	if !reflect.ValueOf(o.AvailabilityMode).IsZero() {
		v.AvailabilityMode = &o.AvailabilityMode
	}
	if !reflect.ValueOf(o.AvailabilityTemplate).IsZero() {
		v.AvailabilityTemplate = &o.AvailabilityTemplate
	}
	if !reflect.ValueOf(o.AvailabilityFunc).IsZero() {
		v.availabilityFunc = o.AvailabilityFunc
	}
	if !reflect.ValueOf(o.CommandFunc).IsZero() {
		v.commandFunc = o.CommandFunc
	} else {
		v.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.States.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.Encoding).IsZero() {
		v.Encoding = &o.Encoding
	}
	if !reflect.ValueOf(o.FanSpeedList).IsZero() {
		v.FanSpeedList = &o.FanSpeedList
	}
	if !reflect.ValueOf(o.JsonAttributesTemplate).IsZero() {
		v.JsonAttributesTemplate = &o.JsonAttributesTemplate
	}
	if !reflect.ValueOf(o.JsonAttributesFunc).IsZero() {
		v.jsonAttributesFunc = o.JsonAttributesFunc
	}
	if !reflect.ValueOf(o.Name).IsZero() {
		v.Name = &o.Name
	}
	if !reflect.ValueOf(o.ObjectId).IsZero() {
		v.ObjectId = &o.ObjectId
	}
	if !reflect.ValueOf(o.PayloadAvailable).IsZero() {
		v.PayloadAvailable = &o.PayloadAvailable
	}
	if !reflect.ValueOf(o.PayloadCleanSpot).IsZero() {
		v.PayloadCleanSpot = &o.PayloadCleanSpot
	}
	if !reflect.ValueOf(o.PayloadLocate).IsZero() {
		v.PayloadLocate = &o.PayloadLocate
	}
	if !reflect.ValueOf(o.PayloadNotAvailable).IsZero() {
		v.PayloadNotAvailable = &o.PayloadNotAvailable
	}
	if !reflect.ValueOf(o.PayloadPause).IsZero() {
		v.PayloadPause = &o.PayloadPause
	}
	if !reflect.ValueOf(o.PayloadReturnToBase).IsZero() {
		v.PayloadReturnToBase = &o.PayloadReturnToBase
	}
	if !reflect.ValueOf(o.PayloadStart).IsZero() {
		v.PayloadStart = &o.PayloadStart
	}
	if !reflect.ValueOf(o.PayloadStop).IsZero() {
		v.PayloadStop = &o.PayloadStop
	}
	if !reflect.ValueOf(o.Qos).IsZero() {
		v.Qos = &o.Qos
	}
	if !reflect.ValueOf(o.Retain).IsZero() {
		v.Retain = &o.Retain
	}
	if !reflect.ValueOf(o.Schema).IsZero() {
		v.Schema = &o.Schema
	}
	if !reflect.ValueOf(o.SendCommandFunc).IsZero() {
		v.sendCommandFunc = o.SendCommandFunc
	}
	if !reflect.ValueOf(o.SetFanSpeedFunc).IsZero() {
		v.setFanSpeedFunc = o.SetFanSpeedFunc
	}
	if !reflect.ValueOf(o.StateFunc).IsZero() {
		v.stateFunc = o.StateFunc
	} else {
		v.stateFunc = func() string {
			return v.States.State
		}
	}
	if !reflect.ValueOf(o.SupportedFeatures).IsZero() {
		v.SupportedFeatures = &o.SupportedFeatures
	}
	if !reflect.ValueOf(o.UniqueId).IsZero() {
		v.UniqueId = &o.UniqueId
	}
	return &v
}

type vacuumState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type VacuumState struct {
	JsonAttributes string
	State          string
}

func (d *Vacuum) SetJsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Vacuum) SetState(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Vacuum) GetRawId() string {
	return "vacuum"
}
func (d *Vacuum) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Vacuum) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Vacuum) GetName() string {
	return *d.Name
}
func (d *Vacuum) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Vacuum) UpdateState() {
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
func (d *Vacuum) Subscribe() {
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
	if d.SendCommandTopic != nil {
		t := c.Subscribe(*d.SendCommandTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SetFanSpeedTopic != nil {
		t := c.Subscribe(*d.SetFanSpeedTopic, 0, d.MQTT.MessageHandler)
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
func (d *Vacuum) UnSubscribe() {
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
	if d.SendCommandTopic != nil {
		t := c.Unsubscribe(*d.SendCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SetFanSpeedTopic != nil {
		t := c.Unsubscribe(*d.SetFanSpeedTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Vacuum) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Vacuum) Initialize() {
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
func (d *Vacuum) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Vacuum) populateTopics() {
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
	if d.sendCommandFunc != nil {
		d.SendCommandTopic = new(string)
		*d.SendCommandTopic = GetTopic(d, "send_command_topic")
		common.TopicStore[*d.SendCommandTopic] = &d.sendCommandFunc
	}
	if d.setFanSpeedFunc != nil {
		d.SetFanSpeedTopic = new(string)
		*d.SetFanSpeedTopic = GetTopic(d, "set_fan_speed_topic")
		common.TopicStore[*d.SetFanSpeedTopic] = &d.setFanSpeedFunc
	}
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Vacuum) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Vacuum) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
