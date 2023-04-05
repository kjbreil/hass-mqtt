package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Light struct {
	config LightConfig
	state  LightState
}
type LightConfig struct {
	Availability        *Topic   // The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`.
	AvailabilityMode    *string  // When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability.
	Brightness          *bool    // Flag that defines if the light supports brightness.
	BrightnessScale     *int     // Defines the maximum brightness value (i.e., 100%) of the MQTT device.
	ColorMode           *bool    // Flag that defines if the light supports color modes.
	Command             Topic    // The MQTT topic to publish commands to change the light’s state.
	Effect              *bool    // Flag that defines if the light supports effects.
	EffectList          []string // The list of effects the light supports.
	EnabledByDefault    *bool    // Flag which defines if the entity should be enabled when first added.
	Encoding            *string  // The encoding of the payloads received and published messages. Set to `""` to disable decoding of incoming payload.
	EntityCategory      *string  // The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity.
	FlashTimeLong       *int     // The duration, in seconds, of a “long” flash.
	FlashTimeShort      *int     // The duration, in seconds, of a “short” flash.
	Icon                *string  // [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	JsonAttributes      *Topic   // The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation.
	MaxMireds           *int     // The maximum color temperature in mireds.
	MinMireds           *int     // The minimum color temperature in mireds.
	Name                *string  // The name of the light.
	ObjectId            *string  // Used instead of `name` for automatic generation of `entity_id`
	Optimistic          *bool    // Flag that defines if the light works in optimistic mode.
	PayloadAvailable    *string  // The payload that represents the available state.
	PayloadNotAvailable *string  // The payload that represents the unavailable state.
	Qos                 *int     // The maximum QoS level of the state topic.
	Retain              *bool    // If the published message should have the retain flag on or not.
	Schema              *string  // The schema to use. Must be `json` to select the JSON schema.
	State               *Topic   // The MQTT topic subscribed to receive state updates.
	SupportedColorModes []string // A list of color modes supported by the list. This is required if `color_mode` is `True`. Possible color modes are `onoff`, `brightness`, `color_temp`, `hs`, `xy`, `rgb`, `rgbw`, `rgbww`, `white`.
	UniqueId            *string  // An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception.
	WhiteScale          *int     // Defines the maximum white level (i.e., 100%) of the MQTT device. This is used when setting the light to white mode.
}
type LightState struct{}
type LightOptions struct{}
