package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SceneOptions struct {
	States               SceneState // External state update location
	AvailabilityMode     string     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc     func() string
	CommandFunc          func(mqtt.Message, mqtt.Client)
	EnabledByDefault     bool   // "Flag which defines if the entity should be enabled when first added."
	EntityCategory       string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                 string // "Icon for the scene."
	Name                 string // "The name to use when displaying this scene."
	ObjectId             string // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable     string // "The payload that represents the available state."
	PayloadNotAvailable  string // "The payload that represents the unavailable state."
	PayloadOn            string // "The payload that will be sent to `command_topic` when activating the MQTT scene."
	Qos                  int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain               bool   // "If the published message should have the retain flag on or not."
	UniqueId             string // "An ID that uniquely identifies this scene entity. If two scenes have the same unique ID, Home Assistant will raise an exception."
}

func NewSceneOptions() *SceneOptions {
	return &SceneOptions{}
}
func (o *SceneOptions) GetStates() *SceneState {
	return &o.States
}
func (o *SceneOptions) SetAvailabilityMode(mode string) *SceneOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *SceneOptions) SetAvailabilityTemplate(template string) *SceneOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *SceneOptions) SetAvailabilityFunc(f func() string) *SceneOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *SceneOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *SceneOptions {
	o.CommandFunc = f
	return o
}
func (o *SceneOptions) SetEnabledByDefault(d bool) *SceneOptions {
	o.EnabledByDefault = d
	return o
}
func (o *SceneOptions) SetEntityCategory(category string) *SceneOptions {
	o.EntityCategory = category
	return o
}
func (o *SceneOptions) SetIcon(icon string) *SceneOptions {
	o.Icon = icon
	return o
}
func (o *SceneOptions) SetName(name string) *SceneOptions {
	o.Name = name
	return o
}
func (o *SceneOptions) SetObjectId(id string) *SceneOptions {
	o.ObjectId = id
	return o
}
func (o *SceneOptions) SetPayloadAvailable(available string) *SceneOptions {
	o.PayloadAvailable = available
	return o
}
func (o *SceneOptions) SetPayloadNotAvailable(available string) *SceneOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *SceneOptions) SetPayloadOn(on string) *SceneOptions {
	o.PayloadOn = on
	return o
}
func (o *SceneOptions) SetQos(qos int) *SceneOptions {
	o.Qos = qos
	return o
}
func (o *SceneOptions) SetRetain(retain bool) *SceneOptions {
	o.Retain = retain
	return o
}
func (o *SceneOptions) SetUniqueId(id string) *SceneOptions {
	o.UniqueId = id
	return o
}
