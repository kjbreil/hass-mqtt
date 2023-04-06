package entities

import (
	"encoding/json"
	common "github.com/kjbreil/hass-mqtt/common"
	"log"
	"reflect"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type DeviceTrigger struct {
	AutomationType *string `json:"automation_type,omitempty"` // "The type of automation, must be 'trigger'."
	Device         Device  `json:"device,omitempty"`          // Device configuration parameters
	Payload        *string `json:"payload,omitempty"`         // "Optional payload to match the payload being sent over the topic."
	Qos            *int    `json:"qos,omitempty"`             // "The maximum QoS level to be used when receiving messages."
	Subtype        *string `json:"subtype,omitempty"`         // "The subtype of the trigger, e.g. `button_1`. Entries supported by the frontend: `turn_on`, `turn_off`, `button_1`, `button_2`, `button_3`, `button_4`, `button_5`, `button_6`. If set to an unsupported value, will render as `subtype type`, e.g. `left_button pressed` with `type` set to `button_short_press` and `subtype` set to `left_button`"
	StateTopic     *string `json:"topic,omitempty"`           // "The MQTT topic subscribed to receive trigger events."
	stateFunc      func() string
	Type           *string             `json:"type,omitempty"`           // "The type of the trigger, e.g. `button_short_press`. Entries supported by the frontend: `button_short_press`, `button_short_release`, `button_long_press`, `button_long_release`, `button_double_press`, `button_triple_press`, `button_quadruple_press`, `button_quintuple_press`. If set to an unsupported value, will render as `subtype type`, e.g. `button_1 spammed` with `type` set to `spammed` and `subtype` set to `button_1`"
	ValueTemplate  *string             `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value."
	MQTT           *MQTTFields         `json:"-"`                        // MQTT configuration parameters
	states         deviceTriggerState  // Internal Holder of States
	States         *DeviceTriggerState `json:"-"` // External state update location
}

func NewDeviceTrigger(o *DeviceTriggerOptions) (*DeviceTrigger, error) {
	var d DeviceTrigger

	d.States = &o.states
	if !reflect.ValueOf(o.automationType).IsZero() {
		d.AutomationType = &o.automationType
	}
	if !reflect.ValueOf(o.payload).IsZero() {
		d.Payload = &o.payload
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		d.Qos = &o.qos
	}
	if !reflect.ValueOf(o.subtype).IsZero() {
		d.Subtype = &o.subtype
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		d.stateFunc = o.stateFunc
	} else {
		d.stateFunc = func() string {
			return d.States.State
		}
	}
	if !reflect.ValueOf(o.entityType).IsZero() {
		d.Type = &o.entityType
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		d.ValueTemplate = &o.valueTemplate
	}
	return &d, nil
}

type deviceTriggerState struct {
	State *string
}
type DeviceTriggerState struct {
	State string
}

func (d *DeviceTrigger) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *DeviceTrigger) GetRawId() string {
	return "device_trigger"
}
func (d *DeviceTrigger) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d DeviceTrigger) GetUniqueId() string {
	return ""
}
func (d DeviceTrigger) GetName() string {
	return ""
}
func (d *DeviceTrigger) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *DeviceTrigger) UpdateState() {
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
}
func (d *DeviceTrigger) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 2, true, message)
	token.WaitTimeout(common.WaitTimeout)
	d.UpdateState()
}
func (d *DeviceTrigger) UnSubscribe()       {}
func (d *DeviceTrigger) AnnounceAvailable() {}
func (d *DeviceTrigger) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(0)
	}
	if d.MQTT == nil {
		d.MQTT = new(MQTTFields)
	}
	d.AddMessageHandler()
	d.populateTopics()
}
func (d *DeviceTrigger) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *DeviceTrigger) populateTopics() {
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *DeviceTrigger) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *DeviceTrigger) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
