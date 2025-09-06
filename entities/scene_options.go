package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type SceneOptions struct {
	states                 SceneState // External state update location
	availabilityMode       string     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string     // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	commandFunc            func(mqtt.Message, mqtt.Client)
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the published messages."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	entityPicture          string // "Picture URL for the entity."
	icon                   string // "Icon for the scene."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name to use when displaying this scene."
	objectId               string // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	payloadAvailable       string // "The payload that represents the available state."
	payloadNotAvailable    string // "The payload that represents the unavailable state."
	payloadOn              string // "The payload that will be sent to `command_topic` when activating the MQTT scene."
	platform               string // "Must be `scene`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                    int    // "The maximum QoS level to be used when receiving and publishing messages."
	retain                 bool   // "If the published message should have the retain flag on or not."
	uniqueId               string // "An ID that uniquely identifies this scene entity. If two scenes have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
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
func (o *SceneOptions) Encoding(encoding string) *SceneOptions {
	o.encoding = encoding
	return o
}
func (o *SceneOptions) EntityCategory(category string) *SceneOptions {
	o.entityCategory = category
	return o
}
func (o *SceneOptions) EntityPicture(picture string) *SceneOptions {
	o.entityPicture = picture
	return o
}
func (o *SceneOptions) Icon(icon string) *SceneOptions {
	o.icon = icon
	return o
}
func (o *SceneOptions) JsonAttributesTemplate(template string) *SceneOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *SceneOptions) JsonAttributesFunc(f func() string) *SceneOptions {
	o.jsonAttributesFunc = f
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
func (o *SceneOptions) Platform(platform string) *SceneOptions {
	o.platform = platform
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
