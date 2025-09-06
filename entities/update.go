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
type Update struct {
	AvailabilityMode       *string `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic      *string `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	availabilityFunc       func() string
	CommandTopic           *string `json:"command_topic,omitempty"` // "The MQTT topic to publish `payload_install` to start installing process."
	commandFunc            func(mqtt.Message, mqtt.Client)
	Device                 Device  `json:"device,omitempty"`                   // Device configuration parameters
	DeviceClass            *string `json:"device_class,omitempty"`             // "The [type/class](/integrations/update/#device-classes) of the update to set the icon in the frontend. The `device_class` can be `null`."
	DisplayPrecision       *int    `json:"display_precision,omitempty"`        // "Number of decimal digits for display of update progress."
	EnabledByDefault       *bool   `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	EntityPicture          *string `json:"entity_picture,omitempty"`           // "Picture URL for the entity."
	Icon                   *string `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic    *string `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as entity attributes. Implies `force_update` of the current select state when a message is received on this topic."
	jsonAttributesFunc     func() string
	LatestVersionTemplate  *string `json:"latest_version_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the latest version value. Use `state_topic` with a `value_template` if all update state values can be extracted from a single JSON payload."
	LatestVersionTopic     *string `json:"latest_version_topic,omitempty"`    // "The MQTT topic subscribed to receive an update of the latest version. Use `state_topic` with a `value_template` if all update state values can be extracted from a single JSON payload."
	latestVersionFunc      func(mqtt.Message, mqtt.Client)
	Name                   *string `json:"name,omitempty"`            // "The name of the Update. Can be set to `null` if only the device name is relevant."
	ObjectId               *string `json:"object_id,omitempty"`       // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	PayloadInstall         *string `json:"payload_install,omitempty"` // "The MQTT payload to start installing process."
	Platform               *string `json:"platform,omitempty"`        // "Must be `update`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	Qos                    *int    `json:"qos,omitempty"`             // "The maximum QoS level to be used when receiving and publishing messages."
	ReleaseSummary         *string `json:"release_summary,omitempty"` // "Summary of the release notes or changelog. This is suitable a brief update description of max 255 characters."
	ReleaseUrl             *string `json:"release_url,omitempty"`     // "URL to the full release notes of the latest version available."
	Retain                 *bool   `json:"retain,omitempty"`          // "If the published message should have the retain flag on or not."
	StateTopic             *string `json:"state_topic,omitempty"`     // "The MQTT topic subscribed to receive state updates. The state update may be either JSON or a simple string with `installed_version` value. When a JSON payload is detected, the state value of the JSON payload should supply the `installed_version` and can optionally supply: `latest_version`, `title`, `release_summary`, `release_url`, and an `entity_picture` URL. To allow progress monitoring `in_progress` (a boolean to indicate an update is in progress), or `update_percentage` (a float value to indicate the progress percentage) may be part of the JSON message."
	stateFunc              func() string
	Title                  *string      `json:"title,omitempty"`          // "Title of the software, or firmware update. This helps to differentiate between the device or entity name versus the title of the software installed."
	UniqueId               *string      `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this Update. If two Updates have the same unique ID Home Assistant will raise an exception."
	ValueTemplate          *string      `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the `installed_version` state value or to render to a valid JSON payload on from the payload received on `state_topic`."
	MQTT                   *MQTTFields  `json:"-"`                        // MQTT configuration parameters
	states                 updateState  // Internal Holder of States
	States                 *UpdateState `json:"-"` // External state update location
}

func (d *Update) Subscribe() {
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
	if d.LatestVersionTopic != nil {
		t := c.Subscribe(*d.LatestVersionTopic, 0, d.MQTT.MessageHandler)
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
func (d *Update) UnSubscribe() {
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
	if d.LatestVersionTopic != nil {
		t := c.Unsubscribe(*d.LatestVersionTopic)
		t.WaitTimeout(common.WaitTimeout)
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func NewUpdate(o *UpdateOptions) (*Update, error) {
	var u Update

	u.States = &o.states
	if !reflect.ValueOf(o.availabilityMode).IsZero() {
		u.AvailabilityMode = &o.availabilityMode
	}
	if !reflect.ValueOf(o.availabilityTemplate).IsZero() {
		u.AvailabilityTemplate = &o.availabilityTemplate
	}
	if !reflect.ValueOf(o.availabilityFunc).IsZero() {
		u.availabilityFunc = o.availabilityFunc
	}
	if !reflect.ValueOf(o.commandFunc).IsZero() {
		u.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			if o.states.State == string(message.Payload()) {
				return
			}
			o.states.State = string(message.Payload())
			u.UpdateState()
			o.commandFunc(message, client)
		}
	} else {
		u.commandFunc = func(message mqtt.Message, client mqtt.Client) {
			o.states.State = string(message.Payload())
		}
	}
	if !reflect.ValueOf(o.deviceClass).IsZero() {
		u.DeviceClass = &o.deviceClass
	}
	if !reflect.ValueOf(o.displayPrecision).IsZero() {
		u.DisplayPrecision = &o.displayPrecision
	}
	if !reflect.ValueOf(o.enabledByDefault).IsZero() {
		u.EnabledByDefault = &o.enabledByDefault
	}
	if !reflect.ValueOf(o.encoding).IsZero() {
		u.Encoding = &o.encoding
	}
	if !reflect.ValueOf(o.entityCategory).IsZero() {
		u.EntityCategory = &o.entityCategory
	}
	if !reflect.ValueOf(o.entityPicture).IsZero() {
		u.EntityPicture = &o.entityPicture
	}
	if !reflect.ValueOf(o.icon).IsZero() {
		u.Icon = &o.icon
	}
	if !reflect.ValueOf(o.jsonAttributesTemplate).IsZero() {
		u.JsonAttributesTemplate = &o.jsonAttributesTemplate
	}
	if !reflect.ValueOf(o.jsonAttributesFunc).IsZero() {
		u.jsonAttributesFunc = o.jsonAttributesFunc
	}
	if !reflect.ValueOf(o.latestVersionTemplate).IsZero() {
		u.LatestVersionTemplate = &o.latestVersionTemplate
	}
	if !reflect.ValueOf(o.latestVersionFunc).IsZero() {
		u.latestVersionFunc = o.latestVersionFunc
	}
	if !reflect.ValueOf(o.name).IsZero() {
		u.Name = &o.name
	} else {
		return nil, fmt.Errorf("name not defined")
	}
	if !reflect.ValueOf(o.objectId).IsZero() {
		u.ObjectId = &o.objectId
	}
	if !reflect.ValueOf(o.payloadInstall).IsZero() {
		u.PayloadInstall = &o.payloadInstall
	}
	if !reflect.ValueOf(o.platform).IsZero() {
		u.Platform = &o.platform
	}
	if !reflect.ValueOf(o.qos).IsZero() {
		u.Qos = &o.qos
	}
	if !reflect.ValueOf(o.releaseSummary).IsZero() {
		u.ReleaseSummary = &o.releaseSummary
	}
	if !reflect.ValueOf(o.releaseUrl).IsZero() {
		u.ReleaseUrl = &o.releaseUrl
	}
	if !reflect.ValueOf(o.retain).IsZero() {
		u.Retain = &o.retain
	}
	if !reflect.ValueOf(o.stateFunc).IsZero() {
		u.stateFunc = o.stateFunc
	} else {
		u.stateFunc = func() string {
			return u.States.State
		}
	}
	if !reflect.ValueOf(o.title).IsZero() {
		u.Title = &o.title
	}
	if !reflect.ValueOf(o.uniqueId).IsZero() {
		u.UniqueId = &o.uniqueId
	} else {
		uniqueId := strcase.ToDelimited(o.name, uint8(0x2d))
		uniqueId = strings.ReplaceAll(uniqueId, "'", "_")
		u.UniqueId = &uniqueId
	}
	if !reflect.ValueOf(o.valueTemplate).IsZero() {
		u.ValueTemplate = &o.valueTemplate
	}
	return &u, nil
}

type updateState struct {
	Availability   *string
	JsonAttributes *string
	State          *string
}
type UpdateState struct {
	JsonAttributes string
	State          string
}

func (d *Update) JsonAttributes(s string) {
	d.States.JsonAttributes = s
	d.UpdateState()
}
func (d *Update) State(s string) {
	d.States.State = s
	d.UpdateState()
}
func (d *Update) GetRawId() string {
	return "update"
}
func (d *Update) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Update) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Update) GetDomainEntity() string {
	return fmt.Sprintf("update.%s", strings.ReplaceAll(*d.UniqueId, "-", "_"))
}
func (d *Update) GetName() string {
	return *d.Name
}
func (d *Update) PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string) {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &Identifier
}
func (d *Update) UpdateState() {
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
func (d *Update) Initialize() {
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
func (d *Update) DeleteTopic() {
	c := *d.MQTT.Client
	token := c.Publish(GetDiscoveryTopic(d), 0, true, "")
	token.WaitTimeout(common.WaitTimeout)
	time.Sleep(common.HADiscoveryDelay)
}
func (d *Update) populateTopics() {
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
	if d.latestVersionFunc != nil {
		d.LatestVersionTopic = new(string)
		*d.LatestVersionTopic = GetTopic(d, "latest_version_topic")
		common.TopicStore[*d.LatestVersionTopic] = &d.latestVersionFunc
	}
	if d.stateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Update) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Update) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
