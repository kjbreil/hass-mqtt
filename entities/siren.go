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
type Siren struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	AvailableTones         *([]string) `json:"available_tones,omitempty"`      // "A list of available tones the siren supports. When configured, this enables the support for setting a `tone` and enables the `tone` state attribute."
	CommandOffTemplate     *string     `json:"command_off_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate a custom payload to send to `command_topic` when the siren turn off service is called. By default `command_template` will be used as template for service turn off. The variable `value` will be assigned with the configured `payload_off` setting."
	CommandTemplate        *string     `json:"command_template,omitempty"`     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate a custom payload to send to `command_topic`. The variable `value` will be assigned with the configured `payload_on` or `payload_off` setting. The siren turn on service parameters `tone`, `volume_level` or `duration` can be used as variables in the template. When operation in optimistic mode the corresponding state attributes will be set. Turn on parameters will be filtered if a device misses the support."
	CommandTopic           *string     `json:"command_topic,omitempty"`        // "The MQTT topic to publish commands to change the siren state. Without command templates, a default JSON payload like `{\"state\":\"ON\", \"tone\": \"bell\", \"duration\": 10, \"volume_level\": 0.5 }` is published. When the siren turn on service is called, the startup parameters will be added to the JSON payload. The `state` value of the JSON payload will be set to the the `payload_on` or `payload_off` configured payload.\n"
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                  // "The name to use when displaying this siren."
	ObjectId               *string `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool   `json:"optimistic,omitempty"`            // "Flag that defines if siren works in optimistic mode."
	PayloadAvailable       *string `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOff             *string `json:"payload_off,omitempty"`           // "The payload that represents `off` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_off` for details) and sending as `off` command to the `command_topic`."
	PayloadOn              *string `json:"payload_on,omitempty"`            // "The payload that represents `on` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_on`  for details) and sending as `on` command to the `command_topic`."
	Qos                    *int    `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 *bool   `json:"retain,omitempty"`                // "If the published message should have the retain flag on or not."
	StateOff               *string `json:"state_off,omitempty"`             // "The payload that represents the `off` state. Used when value that represents `off` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `off`."
	StateOn                *string `json:"state_on,omitempty"`              // "The payload that represents the `on` state. Used when value that represents `on` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `on`."
	StateTopic             *string `json:"state_topic,omitempty"`           // "The MQTT topic subscribed to receive state updates. The state update may be either JSON or a simple string. When a JSON payload is detected, the `state` value of the JSON payload should supply the `payload_on` or `payload_off` defined payload to turn the siren on or off. Additionally, the state attributes `duration`, `tone` and `volume_level` can be updated. Use `value_template` to transform the received state udpate to a compliant JSON payload. Attributes will only be set if the function is supported by the device and a valid value is supplied. When a non JSON payload is detected, it should be either of the `payload_on` or `payload_off` defined payloads or `None` to reset the siren's state to `unknown`. The initial state will be `unknown`. The state will be reset to `unknown` if a `None` payload or `null` JSON value is received as a state update."
	stateFunc              func() string
	StateValueTemplate     *string     `json:"state_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's state from the `state_topic`. To determine the siren's state result of this template will be compared to `state_on` and `state_off`. Alternatively `value_template` can be used to render to a valid JSON payload."
	SupportDuration        *bool       `json:"support_duration,omitempty"`     // "Set to `true` if the MQTT siren supports the `duration` service turn on parameter and enables the `duration` state attribute."
	SupportVolumeSet       *bool       `json:"support_volume_set,omitempty"`   // "Set to `true` if the MQTT siren supports the `volume_set` service turn on parameter and enables the `volume_level` state attribute."
	UniqueId               *string     `json:"unique_id,omitempty"`            // "An ID that uniquely identifies this siren device. If two sirens have the same unique ID, Home Assistant will raise an exception."
	MQTT                   *MQTTFields `json:"-"`                              // MQTT configuration parameters
	states                 sirenState  // Internal Holder of States
	States                 *SirenState `json:"-"` // External state update location
}

func NewSiren(o *SirenOptions) (*Siren, error) {
	var s Siren

	s.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		s.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		s.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		s.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.availableTones).IsZero() {
		s.AvailableTones = &o.availableTones
	}
	if !reflect.ValueOf(o.commandOffTemplate).IsZero() {
		s.CommandOffTemplate = &o.commandOffTemplate
	}
	if !reflect.ValueOf(o.commandTemplate).IsZero() {
		s.CommandTemplate = &o.commandTemplate
	}
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		s.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.State == string(message.Payload()) {
				return
			}
			o.states.State = string(message.Payload())
			s.UpdateState()
			o.commandFunc(message, client)
		}
	} else {
		s.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		s.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		s.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		s.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		s.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		s.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		s.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.name).IsZero() {
		s.Name = &o.name
	} else {
		return nil, fmt.Errorf("name not defined")
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		s.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.optimistic).IsZero() {
		s.Optimistic = &o.optimistic
	}
	if !reflect.ValueOf(o.payloadAvailable).IsZero() {
		s.PayloadAvailable = &o.payloadAvailable
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		s.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadOff).IsZero() {
		s.PayloadOff = &o.payloadOff
	}
	if !reflect.ValueOf(o.payloadOn).IsZero() {
		s.PayloadOn = &o.payloadOn
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		s.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		s.Retain = &o.retain
	}
	if !reflect.ValueOf(o.stateOff).IsZero() {
		s.StateOff = &o.stateOff
	}
	if !reflect.ValueOf(o.stateOn).IsZero() {
		s.StateOn = &o.stateOn
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		s.stateFunc = o.stateFunc
	} else {
		s.stateFunc = func() string {
			return s.States.State
		}
	}
	if !reflect.ValueOf(o.stateValueTemplate).IsZero() {
		s.StateValueTemplate = &o.stateValueTemplate
	}
	if !reflect.ValueOf(o.supportDuration).IsZero() {
		s.SupportDuration = &o.supportDuration
	}
	if !reflect.ValueOf(o.supportVolumeSet).IsZero() {
		s.SupportVolumeSet = &o.supportVolumeSet
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		s.UniqueId = &o.uniqueId
	} else {
		uniqueId := strcase.ToDelimited(o.name, uint8(0x2d))
		uniqueId = strings.ReplaceAll(uniqueId, "'", "_")
		s.UniqueId = &uniqueId
	}
	return &s, nil
}

type sirenState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type SirenState struct {
	JsonAttributes string
	State          string
}

func (d *Siren) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Siren) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Siren) GetRawId() string {
	return "siren"
}
func (d *Siren) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Siren) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Siren) GetDomainEntity() string {
	return fmt.Sprintf("siren.%s", strings.ReplaceAll(*d.UniqueId, "-", "_"))
}
func (d *Siren) GetName() string {
	return *d.Name
}
func (d *Siren) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Siren) UpdateState() {
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
func (d *Siren) Subscribe() {
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
func (d *Siren) UnSubscribe() {
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
func (d *Siren) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, 2, true, "online")
	token.WaitTimeout(common.WaitTimeout)
}
func (d *Siren) Initialize() {
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
func (d *Siren) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Siren) populateTopics() {
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
func (d *Siren) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Siren) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
