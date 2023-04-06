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
type Tag struct {
	Device        Device  `json:"device,omitempty"` // Device configuration parameters
	StateTopic    *string `json:"topic,omitempty"`  // "The MQTT topic subscribed to receive tag scanned events."
	stateFunc     func() string
	ValueTemplate *string     `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that returns a tag ID."
	MQTT          *MQTTFields `json:"-"`                        // MQTT configuration parameters
	states        tagState    // Internal Holder of States
	States        *TagState   `json:"-"` // External state update location
}

func NewTag(o *TagOptions) (*Tag, error) {
	var t Tag

	t.States = &o.states
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		t.stateFunc = o.stateFunc
	} else {
		t.stateFunc = func() string {
			return t.States.State
		}
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		t.ValueTemplate = &o.valueTemplate
	}
	return &t, nil
}

type tagState struct {
	State *string
}
type TagState struct {
	State string
}

func (d *Tag) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Tag) GetRawId() string {
	return "tag"
}
func (d *Tag) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Tag) GetUniqueId() string {
	return ""
}
func (d Tag) GetName() string {
	return ""
}
func (d *Tag) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Tag) UpdateState() {
	if d.StateTopic != nil {
		state := d.stateFunc()
		if d.states.State == nil || state != *d.states.State || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, 2, true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.State = &state
		}
	}
}
func (d *Tag) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 2, true, message)
	token.WaitTimeout(common.WaitTimeout)
	d.UpdateState()
}
func (d *Tag) UnSubscribe()       {}
func (d *Tag) AnnounceAvailable() {}
func (d *Tag) Initialize() {
	if d.MQTT == nil {
		d.MQTT = new(MQTTFields)
	}
	d.AddMessageHandler()
	d.populateTopics()
}
func (d *Tag) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Tag) populateTopics() {
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Tag) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Tag) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
