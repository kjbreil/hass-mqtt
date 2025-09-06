package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type UpdateOptions struct {
	states                 UpdateState // External state update location
	availabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	commandFunc            func(mqtt.Message, mqtt.Client)
	deviceClass            string // "The [type/class](/integrations/update/#device-classes) of the update to set the icon in the frontend. The `device_class` can be `null`."
	displayPrecision       int    // "Number of decimal digits for display of update progress."
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	entityPicture          string // "Picture URL for the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	jsonAttributesFunc     func() string
	latestVersionTemplate  string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the latest version value. Use `state_topic` with a `value_template` if all update state values can be extracted from a single JSON payload."
	latestVersionFunc      func(mqtt.Message, mqtt.Client)
	name                   string // "The name of the Update. Can be set to `null` if only the device name is relevant."
	objectId               string // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	payloadInstall         string // "The MQTT payload to start installing process."
	platform               string // "Must be `update`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                    int    // "The maximum QoS level to be used when receiving and publishing messages."
	releaseSummary         string // "Summary of the release notes or changelog. This is suitable a brief update description of max 255 characters."
	releaseUrl             string // "URL to the full release notes of the latest version available."
	retain                 bool   // "If the published message should have the retain flag on or not."
	stateFunc              func() string
	title                  string // "Title of the software, or firmware update. This helps to differentiate between the device or entity name versus the title of the software installed."
	uniqueId               string // "An ID that uniquely identifies this Update. If two Updates have the same unique ID Home Assistant will raise an exception."
	valueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the `installed_version` state value or to render to a valid JSON payload on from the payload received on `state_topic`."
}

func NewUpdateOptions() *UpdateOptions {
	return &UpdateOptions{}
}
func (o *UpdateOptions) States() *UpdateState {
	return &o.states
}
func (o *UpdateOptions) AvailabilityMode(mode string) *UpdateOptions {
	o.availabilityMode = mode
	return o
}
func (o *UpdateOptions) AvailabilityTemplate(template string) *UpdateOptions {
	o.availabilityTemplate = template
	return o
}
func (o *UpdateOptions) AvailabilityFunc(f func() string) *UpdateOptions {
	o.availabilityFunc = f
	return o
}
func (o *UpdateOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *UpdateOptions {
	o.commandFunc = f
	return o
}
func (o *UpdateOptions) DeviceClass(class string) *UpdateOptions {
	o.deviceClass = class
	return o
}
func (o *UpdateOptions) DisplayPrecision(precision int) *UpdateOptions {
	o.displayPrecision = precision
	return o
}
func (o *UpdateOptions) EnabledByDefault(d bool) *UpdateOptions {
	o.enabledByDefault = d
	return o
}
func (o *UpdateOptions) Encoding(encoding string) *UpdateOptions {
	o.encoding = encoding
	return o
}
func (o *UpdateOptions) EntityCategory(category string) *UpdateOptions {
	o.entityCategory = category
	return o
}
func (o *UpdateOptions) EntityPicture(picture string) *UpdateOptions {
	o.entityPicture = picture
	return o
}
func (o *UpdateOptions) Icon(icon string) *UpdateOptions {
	o.icon = icon
	return o
}
func (o *UpdateOptions) JsonAttributesTemplate(template string) *UpdateOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *UpdateOptions) JsonAttributesFunc(f func() string) *UpdateOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *UpdateOptions) LatestVersionTemplate(template string) *UpdateOptions {
	o.latestVersionTemplate = template
	return o
}
func (o *UpdateOptions) LatestVersionFunc(f func(mqtt.Message, mqtt.Client)) *UpdateOptions {
	o.latestVersionFunc = f
	return o
}
func (o *UpdateOptions) Name(name string) *UpdateOptions {
	o.name = name
	return o
}
func (o *UpdateOptions) ObjectId(id string) *UpdateOptions {
	o.objectId = id
	return o
}
func (o *UpdateOptions) PayloadInstall(install string) *UpdateOptions {
	o.payloadInstall = install
	return o
}
func (o *UpdateOptions) Platform(platform string) *UpdateOptions {
	o.platform = platform
	return o
}
func (o *UpdateOptions) Qos(qos int) *UpdateOptions {
	o.qos = qos
	return o
}
func (o *UpdateOptions) ReleaseSummary(summary string) *UpdateOptions {
	o.releaseSummary = summary
	return o
}
func (o *UpdateOptions) ReleaseUrl(url string) *UpdateOptions {
	o.releaseUrl = url
	return o
}
func (o *UpdateOptions) Retain(retain bool) *UpdateOptions {
	o.retain = retain
	return o
}
func (o *UpdateOptions) StateFunc(f func() string) *UpdateOptions {
	o.stateFunc = f
	return o
}
func (o *UpdateOptions) Title(title string) *UpdateOptions {
	o.title = title
	return o
}
func (o *UpdateOptions) UniqueId(id string) *UpdateOptions {
	o.uniqueId = id
	return o
}
func (o *UpdateOptions) ValueTemplate(template string) *UpdateOptions {
	o.valueTemplate = template
	return o
}
