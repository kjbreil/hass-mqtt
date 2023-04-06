package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type CameraOptions struct {
	states                 CameraState // External state update location
	availabilityMode       string      // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string      // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload. Use `image_encoding` to enable `Base64` decoding on `topic`."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	imageEncoding          string // "The encoding of the image payloads received. Set to `\"b64\"` to enable base64 decoding of image payload. If not set, the image payload must be raw binary data."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	jsonAttributesFunc     func() string
	name                   string // "The name of the camera."
	objectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	stateFunc              func() string
	uniqueId               string // "An ID that uniquely identifies this camera. If two cameras have the same unique ID Home Assistant will raise an exception."
}

func NewCameraOptions() *CameraOptions {
	return &CameraOptions{}
}
func (o *CameraOptions) States() *CameraState {
	return &o.states
}
func (o *CameraOptions) AvailabilityMode(mode string) *CameraOptions {
	o.availabilityMode = mode
	return o
}
func (o *CameraOptions) AvailabilityTemplate(template string) *CameraOptions {
	o.availabilityTemplate = template
	return o
}
func (o *CameraOptions) AvailabilityFunc(f func() string) *CameraOptions {
	o.availabilityFunc = f
	return o
}
func (o *CameraOptions) EnabledByDefault(d bool) *CameraOptions {
	o.enabledByDefault = d
	return o
}
func (o *CameraOptions) Encoding(encoding string) *CameraOptions {
	o.encoding = encoding
	return o
}
func (o *CameraOptions) EntityCategory(category string) *CameraOptions {
	o.entityCategory = category
	return o
}
func (o *CameraOptions) Icon(icon string) *CameraOptions {
	o.icon = icon
	return o
}
func (o *CameraOptions) ImageEncoding(encoding string) *CameraOptions {
	o.imageEncoding = encoding
	return o
}
func (o *CameraOptions) JsonAttributesTemplate(template string) *CameraOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *CameraOptions) JsonAttributesFunc(f func() string) *CameraOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *CameraOptions) Name(name string) *CameraOptions {
	o.name = name
	return o
}
func (o *CameraOptions) ObjectId(id string) *CameraOptions {
	o.objectId = id
	return o
}
func (o *CameraOptions) StateFunc(f func() string) *CameraOptions {
	o.stateFunc = f
	return o
}
func (o *CameraOptions) UniqueId(id string) *CameraOptions {
	o.uniqueId = id
	return o
}
