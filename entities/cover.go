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
type Cover struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to to receive birth and LWT messages from the MQTT cover device. If an `availability` topic is not defined, the cover availability state will always be `available`. If an `availability` topic is defined, the cover availability state will be `unavailable` by default. Must not be used together with `availability`."
	availabilityFunc       func() string
	CommandTopic           *string `json:"command_topic,omitempty"` // "The MQTT topic to publish commands to control the cover."
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	DeviceClass            *string `json:"device_class,omitempty"`             // "Sets the [class of the device](/integrations/cover/), changing the device state and icon that is displayed on the frontend."
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	jsonAttributesFunc     func() string
	Name                   *string `json:"name,omitempty"`                  // "The name of the cover."
	ObjectId               *string `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool   `json:"optimistic,omitempty"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable       *string `json:"payload_available,omitempty"`     // "The payload that represents the online state."
	PayloadClose           *string `json:"payload_close,omitempty"`         // "The command payload that closes the cover."
	PayloadNotAvailable    *string `json:"payload_not_available,omitempty"` // "The payload that represents the offline state."
	PayloadOpen            *string `json:"payload_open,omitempty"`          // "The command payload that opens the cover."
	PayloadStop            *string `json:"payload_stop,omitempty"`          // "The command payload that stops the cover."
	PositionClosed         *int    `json:"position_closed,omitempty"`       // "Number which represents closed position."
	PositionOpen           *int    `json:"position_open,omitempty"`         // "Number which represents open position."
	PositionTemplate       *string `json:"position_template,omitempty"`     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that can be used to extract the payload for the `position_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [States](/docs/configuration/templating/#States) template function;"
	PositionTopic          *string `json:"position_topic,omitempty"`        // "The MQTT topic subscribed to receive cover position messages."
	positionFunc           func() string
	Qos                    *int    `json:"qos,omitempty"`                   // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                 *bool   `json:"retain,omitempty"`                // "Defines if published messages should have the retain flag set."
	SetPositionTemplate    *string `json:"set_position_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to define the position to be sent to the `set_position_topic` topic. Incoming position value is available for use in the template `{% raw %}{{ position }}{% endraw %}`. Within the template the following variables are available: `entity_id`, `position`, the target position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [States](/docs/configuration/templating/#States) template function;"
	SetPositionTopic       *string `json:"set_position_topic,omitempty"`    // "The MQTT topic to publish position commands to. You need to set position_topic as well if you want to use position topic. Use template if position topic wants different values than within range `position_closed` - `position_open`. If template is not defined and `position_closed != 100` and `position_open != 0` then proper position value is calculated from percentage position."
	setPositionFunc        func(mqtt.Message, mqtt.Client)
	StateClosed            *string `json:"state_closed,omitempty"`  // "The payload that represents the closed state."
	StateClosing           *string `json:"state_closing,omitempty"` // "The payload that represents the closing state."
	StateOpen              *string `json:"state_open,omitempty"`    // "The payload that represents the open state."
	StateOpening           *string `json:"state_opening,omitempty"` // "The payload that represents the opening state."
	StateStopped           *string `json:"state_stopped,omitempty"` // "The payload that represents the stopped state (for covers that do not report `open`/`closed` state)."
	StateTopic             *string `json:"state_topic,omitempty"`   // "The MQTT topic subscribed to receive cover state messages. State topic can only read (`open`, `opening`, `closed`, `closing` or `stopped`) state."
	stateFunc              func() string
	TiltClosedValue        *int    `json:"tilt_closed_value,omitempty"`     // "The value that will be sent on a `close_cover_tilt` command."
	TiltCommandTemplate    *string `json:"tilt_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that can be used to extract the payload for the `tilt_command_topic` topic. Within the template the following variables are available: `entity_id`, `tilt_position`, the target tilt position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [States](/docs/configuration/templating/#States) template function;"
	TiltCommandTopic       *string `json:"tilt_command_topic,omitempty"`    // "The MQTT topic to publish commands to control the cover tilt."
	tiltCommandFunc        func(mqtt.Message, mqtt.Client)
	TiltMax                *int    `json:"tilt_max,omitempty"`             // "The maximum tilt value."
	TiltMin                *int    `json:"tilt_min,omitempty"`             // "The minimum tilt value."
	TiltOpenedValue        *int    `json:"tilt_opened_value,omitempty"`    // "The value that will be sent on an `open_cover_tilt` command."
	TiltOptimistic         *bool   `json:"tilt_optimistic,omitempty"`      // "Flag that determines if tilt works in optimistic mode."
	TiltStatusTemplate     *string `json:"tilt_status_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that can be used to extract the payload for the `tilt_status_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [States](/docs/configuration/templating/#States) template function;"
	TiltStatusTopic        *string `json:"tilt_status_topic,omitempty"`    // "The MQTT topic subscribed to receive tilt status update values."
	tiltStatusFunc         func() string
	UniqueId               *string     `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this cover. If two covers have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that can be used to extract the payload for the `state_topic` topic."
	MQTT                   *MQTTFields `json:"-"`                        // MQTT configuration parameters
	states                 coverState  // Internal Holder of States
	States                 *CoverState `json:"-"` // External state update location
}

func (d *Cover) Subscribe() {
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
	if d.SetPositionTopic != nil {
		t := c.Subscribe(*d.SetPositionTopic, 0, d.MQTT.MessageHandler)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TiltCommandTopic != nil {
		t := c.Subscribe(*d.TiltCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d *Cover) UnSubscribe() {
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
	if d.SetPositionTopic != nil {
		t := c.Unsubscribe(*d.SetPositionTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TiltCommandTopic != nil {
		t := c.Unsubscribe(*d.TiltCommandTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func NewCover(o *CoverOptions) (*Cover, error) {
	var c Cover

	c.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		c.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		c.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		c.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		c.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.State == string(message.Payload()) {
				return
			}
			o.states.State = string(message.Payload())
			c.UpdateState()
			o.commandFunc(message, client)
		}
	} else {
		c.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.deviceClass).IsZero() {
		c.DeviceClass = &o.deviceClass
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
	if !reflect.ValueOf(o.icon).IsZero() {
		c.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		c.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		c.jsonAttributesFunc = o.jsonAttributesFunc
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
	if !reflect.ValueOf(o.payloadClose).IsZero() {
		c.PayloadClose = &o.payloadClose
	}
	if !reflect.ValueOf(o.payloadNotAvailable).IsZero() {
		c.PayloadNotAvailable = &o.payloadNotAvailable
	}
	if !reflect.ValueOf(o.payloadOpen).IsZero() {
		c.PayloadOpen = &o.payloadOpen
	}
	if !reflect.ValueOf(o.payloadStop).IsZero() {
		c.PayloadStop = &o.payloadStop
	}
	if !reflect.ValueOf(o.positionClosed).IsZero() {
		c.PositionClosed = &o.positionClosed
	}
	if !reflect.ValueOf(o.positionOpen).IsZero() {
		c.PositionOpen = &o.positionOpen
	}
	if !reflect.ValueOf(o.positionTemplate).IsZero() {
		c.PositionTemplate = &o.positionTemplate
	}
	if !reflect.ValueOf(o.positionFunc).IsZero() {
		c.positionFunc = o.positionFunc
	} else {
		c.positionFunc = func() string {
			return c.States.Position
		}
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		c.Qos = &o.qos
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		c.Retain = &o.retain
	}
	if !reflect.ValueOf(o.setPositionTemplate).IsZero() {
		c.SetPositionTemplate = &o.setPositionTemplate
	}
	if !reflect.ValueOf(o.setPositionFunc).IsZero() {
		c.setPositionFunc = o.setPositionFunc
	}
	if !reflect.ValueOf(o.stateClosed).IsZero() {
		c.StateClosed = &o.stateClosed
	}
	if !reflect.ValueOf(o.stateClosing).IsZero() {
		c.StateClosing = &o.stateClosing
	}
	if !reflect.ValueOf(o.stateOpen).IsZero() {
		c.StateOpen = &o.stateOpen
	}
	if !reflect.ValueOf(o.stateOpening).IsZero() {
		c.StateOpening = &o.stateOpening
	}
	if !reflect.ValueOf(o.stateStopped).IsZero() {
		c.StateStopped = &o.stateStopped
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		c.stateFunc = o.stateFunc
	} else {
		c.stateFunc = func() string {
			return c.States.State
		}
	}
	if !reflect.ValueOf(o.tiltClosedValue).IsZero() {
		c.TiltClosedValue = &o.tiltClosedValue
	}
	if !reflect.ValueOf(o.tiltCommandTemplate).IsZero() {
		c.TiltCommandTemplate = &o.tiltCommandTemplate
	}
	if !reflect.ValueOf(o.tiltCommandFunc).IsZero() {
		c.tiltCommandFunc = o.tiltCommandFunc
	}
	if !reflect.ValueOf(o.tiltMax).IsZero() {
		c.TiltMax = &o.tiltMax
	}
	if !reflect.ValueOf(o.tiltMin).IsZero() {
		c.TiltMin = &o.tiltMin
	}
	if !reflect.ValueOf(o.tiltOpenedValue).IsZero() {
		c.TiltOpenedValue = &o.tiltOpenedValue
	}
	if !reflect.ValueOf(o.tiltOptimistic).IsZero() {
		c.TiltOptimistic = &o.tiltOptimistic
	}
	if !reflect.ValueOf(o.tiltStatusTemplate).IsZero() {
		c.TiltStatusTemplate = &o.tiltStatusTemplate
	}
	if !reflect.ValueOf(o.tiltStatusFunc).IsZero() {
		c.tiltStatusFunc = o.tiltStatusFunc
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

type coverState struct {
	Availability   *string
	JsonAttributes *string
	Position       *string
	State          *string
	TiltStatus     *string
}
type CoverState struct {
	JsonAttributes string
	Position       string
	State          string
	TiltStatus     string
}

func (d *Cover) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Cover) Position(s string) {
	d.States.Position = s
	d.UpdateState()
}
func (d *Cover) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Cover) TiltStatus(s string) {
	d.States.TiltStatus = s
	d.UpdateState()
}
func (d *Cover) GetRawId() string {
	return "cover"
}
func (d *Cover) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Cover) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Cover) GetDomainEntity() string {
	return fmt.Sprintf("cover.%s", strings.ReplaceAll(*d.UniqueId, "-", "_"))
}
func (d *Cover) GetName() string {
	return *d.Name
}
func (d *Cover) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Cover) UpdateState() {
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
	if d.PositionTopic != nil {
		state := d.positionFunc()
		if d.states.Position == nil || state != *d.states.Position || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.PositionTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.Position = &state
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
	if d.TiltStatusTopic != nil {
		state := d.tiltStatusFunc()
		if d.states.TiltStatus == nil || state != *d.states.TiltStatus || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TiltStatusTopic, byte(*d.Qos), true, state)
			token.WaitTimeout(common.WaitTimeout)
			d.states.TiltStatus = &state
		}
	}
}
func (d *Cover) Initialize() {
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
func (d *Cover) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Cover) populateTopics() {
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
	if d.positionFunc != nil {
		d.PositionTopic = new(string)
		*d.PositionTopic = GetTopic(d, "position_topic")
	}
	if d.setPositionFunc != nil {
		d.SetPositionTopic = new(string)
		*d.SetPositionTopic = GetTopic(d, "set_position_topic")
		common.TopicStore[*d.SetPositionTopic] = &d.setPositionFunc
	}
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.tiltCommandFunc != nil {
		d.TiltCommandTopic = new(string)
		*d.TiltCommandTopic = GetTopic(d, "tilt_command_topic")
		common.TopicStore[*d.TiltCommandTopic] = &d.tiltCommandFunc
	}
	if d.tiltStatusFunc != nil {
		d.TiltStatusTopic = new(string)
		*d.TiltStatusTopic = GetTopic(d, "tilt_status_topic")
	}
}
func (d *Cover) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Cover) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
