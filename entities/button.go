package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Button struct {
	config ButtonConfig
}
type ButtonConfig struct {
	Availability        *string // The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`.
	AvailabilityMode    *string // When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability.
	Command             *string // The MQTT topic to publish commands to trigger the button.
	DeviceClass         *string // The [type/class](/integrations/button/#device-class) of the button to set the icon in the frontend.
	EnabledByDefault    *bool   // Flag which defines if the entity should be enabled when first added.
	Encoding            *string // The encoding of the published messages.
	EntityCategory      *string // The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity.
	Icon                *string // [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	JsonAttributes      *string // The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation.
	Name                *string // The name to use when displaying this button.
	ObjectId            *string // Used instead of `name` for automatic generation of `entity_id`
	PayloadAvailable    *string // The payload that represents the available state.
	PayloadNotAvailable *string // The payload that represents the unavailable state.
	PayloadPress        *string // The payload To send to trigger the button.
	Qos                 *int    // The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages.
	Retain              *bool   // If the published message should have the retain flag on or not.
	UniqueId            *string // An ID that uniquely identifies this button entity. If two buttons have the same unique ID, Home Assistant will raise an exception.
}
