package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type UpdateOptions struct {
	States                 UpdateState // External state update location
	AvailabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	CommandFunc            func(mqtt.Message, mqtt.Client)
	DeviceClass            string // "The [type/class](/integrations/update/#device-classes) of the update to set the icon in the frontend."
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	EntityPicture          string // "Picture URL for the entity."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesFunc     func() string
	LatestVersionTemplate  string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the latest version value."
	LatestVersionFunc      func(mqtt.Message, mqtt.Client)
	Name                   string // "The name of the Select."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadInstall         string // "The MQTT payload to start installing process."
	Qos                    int    // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	ReleaseSummary         string // "Summary of the release notes or changelog. This is suitable a brief update description of max 255 characters."
	ReleaseUrl             string // "URL to the full release notes of the latest version available."
	Retain                 bool   // "If the published message should have the retain flag on or not."
	StateFunc              func() string
	Title                  string // "Title of the software, or firmware update. This helps to differentiate between the device or entity name versus the title of the software installed."
	UniqueId               string // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception."
	ValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `installed_version` state value or to render to a valid JSON payload on from the payload received on `state_topic`."
}

func NewUpdateOptions() *UpdateOptions {
	return &UpdateOptions{}
}
func (o *UpdateOptions) GetStates() *UpdateState {
	return &o.States
}
func (o *UpdateOptions) SetAvailabilityMode(mode string) *UpdateOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *UpdateOptions) SetAvailabilityTemplate(template string) *UpdateOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *UpdateOptions) SetAvailabilityFunc(f func() string) *UpdateOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *UpdateOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *UpdateOptions {
	o.CommandFunc = f
	return o
}
func (o *UpdateOptions) SetDeviceClass(class string) *UpdateOptions {
	o.DeviceClass = class
	return o
}
func (o *UpdateOptions) SetEnabledByDefault(d bool) *UpdateOptions {
	o.EnabledByDefault = d
	return o
}
func (o *UpdateOptions) SetEncoding(encoding string) *UpdateOptions {
	o.Encoding = encoding
	return o
}
func (o *UpdateOptions) SetEntityCategory(category string) *UpdateOptions {
	o.EntityCategory = category
	return o
}
func (o *UpdateOptions) SetEntityPicture(picture string) *UpdateOptions {
	o.EntityPicture = picture
	return o
}
func (o *UpdateOptions) SetIcon(icon string) *UpdateOptions {
	o.Icon = icon
	return o
}
func (o *UpdateOptions) SetJsonAttributesTemplate(template string) *UpdateOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *UpdateOptions) SetJsonAttributesFunc(f func() string) *UpdateOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *UpdateOptions) SetLatestVersionTemplate(template string) *UpdateOptions {
	o.LatestVersionTemplate = template
	return o
}
func (o *UpdateOptions) SetLatestVersionFunc(f func(mqtt.Message, mqtt.Client)) *UpdateOptions {
	o.LatestVersionFunc = f
	return o
}
func (o *UpdateOptions) SetName(name string) *UpdateOptions {
	o.Name = name
	return o
}
func (o *UpdateOptions) SetObjectId(id string) *UpdateOptions {
	o.ObjectId = id
	return o
}
func (o *UpdateOptions) SetPayloadInstall(install string) *UpdateOptions {
	o.PayloadInstall = install
	return o
}
func (o *UpdateOptions) SetQos(qos int) *UpdateOptions {
	o.Qos = qos
	return o
}
func (o *UpdateOptions) SetReleaseSummary(summary string) *UpdateOptions {
	o.ReleaseSummary = summary
	return o
}
func (o *UpdateOptions) SetReleaseUrl(url string) *UpdateOptions {
	o.ReleaseUrl = url
	return o
}
func (o *UpdateOptions) SetRetain(retain bool) *UpdateOptions {
	o.Retain = retain
	return o
}
func (o *UpdateOptions) SetStateFunc(f func() string) *UpdateOptions {
	o.StateFunc = f
	return o
}
func (o *UpdateOptions) SetTitle(title string) *UpdateOptions {
	o.Title = title
	return o
}
func (o *UpdateOptions) SetUniqueId(id string) *UpdateOptions {
	o.UniqueId = id
	return o
}
func (o *UpdateOptions) SetValueTemplate(template string) *UpdateOptions {
	o.ValueTemplate = template
	return o
}
