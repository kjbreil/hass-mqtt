package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SceneOptions struct {
	states               SceneState // External state update location
	availabilityMode     string     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc     func() string
	commandFunc          func(mqtt.Message, mqtt.Client)
	enabledByDefault     bool   // "Flag which defines if the entity should be enabled when first added."
	entityCategory       string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	icon                 string // "Icon for the scene."
	name                 string // "The name to use when displaying this scene."
	objectId             string // "Used instead of `name` for automatic generation of `entity_id`"
	payloadAvailable     string // "The payload that represents the available state."
	payloadNotAvailable  string // "The payload that represents the unavailable state."
	payloadOn            string // "The payload that will be sent to `command_topic` when activating the MQTT scene."
	qos                  int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	retain               bool   // "If the published message should have the retain flag on or not."
	uniqueId             string // "An ID that uniquely identifies this scene entity. If two scenes have the same unique ID, Home Assistant will raise an exception."
}

func NewSceneOptions() *SceneOptions {
	return &SceneOptions{}
}
func (o *SceneOptions) States() *SceneState {
	return &o.states
}
func (o *SceneOptions) AvailabilityMode(mode string) *SceneOptions {
	o.availabilityMode = mode
	return o
}
func (o *SceneOptions) AvailabilityTemplate(template string) *SceneOptions {
	o.availabilityTemplate = template
	return o
}
func (o *SceneOptions) AvailabilityFunc(f func() string) *SceneOptions {
	o.availabilityFunc = f
	return o
}
func (o *SceneOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *SceneOptions {
	o.commandFunc = f
	return o
}
func (o *SceneOptions) EnabledByDefault(d bool) *SceneOptions {
	o.enabledByDefault = d
	return o
}
func (o *SceneOptions) EntityCategory(category string) *SceneOptions {
	o.entityCategory = category
	return o
}
func (o *SceneOptions) Icon(icon string) *SceneOptions {
	o.icon = icon
	return o
}
func (o *SceneOptions) Name(name string) *SceneOptions {
	o.name = name
	return o
}
func (o *SceneOptions) ObjectId(id string) *SceneOptions {
	o.objectId = id
	return o
}
func (o *SceneOptions) PayloadAvailable(available string) *SceneOptions {
	o.payloadAvailable = available
	return o
}
func (o *SceneOptions) PayloadNotAvailable(available string) *SceneOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *SceneOptions) PayloadOn(on string) *SceneOptions {
	o.payloadOn = on
	return o
}
func (o *SceneOptions) Qos(qos int) *SceneOptions {
	o.qos = qos
	return o
}
func (o *SceneOptions) Retain(retain bool) *SceneOptions {
	o.retain = retain
	return o
}
func (o *SceneOptions) UniqueId(id string) *SceneOptions {
	o.uniqueId = id
	return o
}
