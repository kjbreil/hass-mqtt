package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type CameraOptions struct {
	States                 CameraState // External state update location
	AvailabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload. Use `image_encoding` to enable `Base64` decoding on `topic`."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	ImageEncoding          string // "The encoding of the image payloads received. Set to `\"b64\"` to enable base64 decoding of image payload. If not set, the image payload must be raw binary data."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesFunc     func() string
	Name                   string // "The name of the camera."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	StateFunc              func() string
	UniqueId               string // "An ID that uniquely identifies this camera. If two cameras have the same unique ID Home Assistant will raise an exception."
}

func NewCameraOptions() *CameraOptions {
	return &CameraOptions{}
}
func (o *CameraOptions) GetStates() *CameraState {
	return &o.States
}
func (o *CameraOptions) SetAvailabilityMode(mode string) *CameraOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *CameraOptions) SetAvailabilityTemplate(template string) *CameraOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *CameraOptions) SetAvailabilityFunc(f func() string) *CameraOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *CameraOptions) SetEnabledByDefault(d bool) *CameraOptions {
	o.EnabledByDefault = d
	return o
}
func (o *CameraOptions) SetEncoding(encoding string) *CameraOptions {
	o.Encoding = encoding
	return o
}
func (o *CameraOptions) SetEntityCategory(category string) *CameraOptions {
	o.EntityCategory = category
	return o
}
func (o *CameraOptions) SetIcon(icon string) *CameraOptions {
	o.Icon = icon
	return o
}
func (o *CameraOptions) SetImageEncoding(encoding string) *CameraOptions {
	o.ImageEncoding = encoding
	return o
}
func (o *CameraOptions) SetJsonAttributesTemplate(template string) *CameraOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *CameraOptions) SetJsonAttributesFunc(f func() string) *CameraOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *CameraOptions) SetName(name string) *CameraOptions {
	o.Name = name
	return o
}
func (o *CameraOptions) SetObjectId(id string) *CameraOptions {
	o.ObjectId = id
	return o
}
func (o *CameraOptions) SetStateFunc(f func() string) *CameraOptions {
	o.StateFunc = f
	return o
}
func (o *CameraOptions) SetUniqueId(id string) *CameraOptions {
	o.UniqueId = id
	return o
}
