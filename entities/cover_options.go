package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type CoverOptions struct {
	states                 CoverState // External state update location
	availabilityMode       string     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	commandFunc            func(mqtt.Message, mqtt.Client)
	deviceClass            string // "Sets the [class of the device](/integrations/cover/), changing the device state and icon that is displayed on the frontend."
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name of the cover."
	objectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	optimistic             bool   // "Flag that defines if switch works in optimistic mode."
	payloadAvailable       string // "The payload that represents the online state."
	payloadClose           string // "The command payload that closes the cover."
	payloadNotAvailable    string // "The payload that represents the offline state."
	payloadOpen            string // "The command payload that opens the cover."
	payloadStop            string // "The command payload that stops the cover."
	positionClosed         int    // "Number which represents closed position."
	positionOpen           int    // "Number which represents open position."
	positionTemplate       string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that can be used to extract the payload for the `position_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [States](/docs/configuration/templating/#States) template function;"
	positionFunc           func() string
	qos                    int    // "The maximum QoS level to be used when receiving and publishing messages."
	retain                 bool   // "Defines if published messages should have the retain flag set."
	setPositionTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to define the position to be sent to the `set_position_topic` topic. Incoming position value is available for use in the template `{% raw %}{{ position }}{% endraw %}`. Within the template the following variables are available: `entity_id`, `position`, the target position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [States](/docs/configuration/templating/#States) template function;"
	setPositionFunc        func(mqtt.Message, mqtt.Client)
	stateClosed            string // "The payload that represents the closed state."
	stateClosing           string // "The payload that represents the closing state."
	stateOpen              string // "The payload that represents the open state."
	stateOpening           string // "The payload that represents the opening state."
	stateStopped           string // "The payload that represents the stopped state (for covers that do not report `open`/`closed` state)."
	stateFunc              func() string
	tiltClosedValue        int    // "The value that will be sent on a `close_cover_tilt` command."
	tiltCommandTemplate    string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that can be used to extract the payload for the `tilt_command_topic` topic. Within the template the following variables are available: `entity_id`, `tilt_position`, the target tilt position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [States](/docs/configuration/templating/#States) template function;"
	tiltCommandFunc        func(mqtt.Message, mqtt.Client)
	tiltMax                int    // "The maximum tilt value."
	tiltMin                int    // "The minimum tilt value."
	tiltOpenedValue        int    // "The value that will be sent on an `open_cover_tilt` command."
	tiltOptimistic         bool   // "Flag that determines if tilt works in optimistic mode."
	tiltStatusTemplate     string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that can be used to extract the payload for the `tilt_status_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [States](/docs/configuration/templating/#States) template function;"
	tiltStatusFunc         func() string
	uniqueId               string // "An ID that uniquely identifies this cover. If two covers have the same unique ID, Home Assistant will raise an exception."
	valueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that can be used to extract the payload for the `state_topic` topic."
}

func NewCoverOptions() *CoverOptions {
	return &CoverOptions{}
}
func (o *CoverOptions) States() *CoverState {
	return &o.states
}
func (o *CoverOptions) AvailabilityMode(mode string) *CoverOptions {
	o.availabilityMode = mode
	return o
}
func (o *CoverOptions) AvailabilityTemplate(template string) *CoverOptions {
	o.availabilityTemplate = template
	return o
}
func (o *CoverOptions) AvailabilityFunc(f func() string) *CoverOptions {
	o.availabilityFunc = f
	return o
}
func (o *CoverOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *CoverOptions {
	o.commandFunc = f
	return o
}
func (o *CoverOptions) DeviceClass(class string) *CoverOptions {
	o.deviceClass = class
	return o
}
func (o *CoverOptions) EnabledByDefault(d bool) *CoverOptions {
	o.enabledByDefault = d
	return o
}
func (o *CoverOptions) Encoding(encoding string) *CoverOptions {
	o.encoding = encoding
	return o
}
func (o *CoverOptions) EntityCategory(category string) *CoverOptions {
	o.entityCategory = category
	return o
}
func (o *CoverOptions) Icon(icon string) *CoverOptions {
	o.icon = icon
	return o
}
func (o *CoverOptions) JsonAttributesTemplate(template string) *CoverOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *CoverOptions) JsonAttributesFunc(f func() string) *CoverOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *CoverOptions) Name(name string) *CoverOptions {
	o.name = name
	return o
}
func (o *CoverOptions) ObjectId(id string) *CoverOptions {
	o.objectId = id
	return o
}
func (o *CoverOptions) Optimistic(optimistic bool) *CoverOptions {
	o.optimistic = optimistic
	return o
}
func (o *CoverOptions) PayloadAvailable(available string) *CoverOptions {
	o.payloadAvailable = available
	return o
}
func (o *CoverOptions) PayloadClose(close string) *CoverOptions {
	o.payloadClose = close
	return o
}
func (o *CoverOptions) PayloadNotAvailable(available string) *CoverOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *CoverOptions) PayloadOpen(open string) *CoverOptions {
	o.payloadOpen = open
	return o
}
func (o *CoverOptions) PayloadStop(stop string) *CoverOptions {
	o.payloadStop = stop
	return o
}
func (o *CoverOptions) PositionClosed(closed int) *CoverOptions {
	o.positionClosed = closed
	return o
}
func (o *CoverOptions) PositionOpen(open int) *CoverOptions {
	o.positionOpen = open
	return o
}
func (o *CoverOptions) PositionTemplate(template string) *CoverOptions {
	o.positionTemplate = template
	return o
}
func (o *CoverOptions) PositionFunc(f func() string) *CoverOptions {
	o.positionFunc = f
	return o
}
func (o *CoverOptions) Qos(qos int) *CoverOptions {
	o.qos = qos
	return o
}
func (o *CoverOptions) Retain(retain bool) *CoverOptions {
	o.retain = retain
	return o
}
func (o *CoverOptions) SetPositionTemplate(template string) *CoverOptions {
	o.setPositionTemplate = template
	return o
}
func (o *CoverOptions) SetPositionFunc(f func(mqtt.Message, mqtt.Client)) *CoverOptions {
	o.setPositionFunc = f
	return o
}
func (o *CoverOptions) StateClosed(closed string) *CoverOptions {
	o.stateClosed = closed
	return o
}
func (o *CoverOptions) StateClosing(closing string) *CoverOptions {
	o.stateClosing = closing
	return o
}
func (o *CoverOptions) StateOpen(open string) *CoverOptions {
	o.stateOpen = open
	return o
}
func (o *CoverOptions) StateOpening(opening string) *CoverOptions {
	o.stateOpening = opening
	return o
}
func (o *CoverOptions) StateStopped(stopped string) *CoverOptions {
	o.stateStopped = stopped
	return o
}
func (o *CoverOptions) StateFunc(f func() string) *CoverOptions {
	o.stateFunc = f
	return o
}
func (o *CoverOptions) TiltClosedValue(value int) *CoverOptions {
	o.tiltClosedValue = value
	return o
}
func (o *CoverOptions) TiltCommandTemplate(template string) *CoverOptions {
	o.tiltCommandTemplate = template
	return o
}
func (o *CoverOptions) TiltCommandFunc(f func(mqtt.Message, mqtt.Client)) *CoverOptions {
	o.tiltCommandFunc = f
	return o
}
func (o *CoverOptions) TiltMax(max int) *CoverOptions {
	o.tiltMax = max
	return o
}
func (o *CoverOptions) TiltMin(min int) *CoverOptions {
	o.tiltMin = min
	return o
}
func (o *CoverOptions) TiltOpenedValue(value int) *CoverOptions {
	o.tiltOpenedValue = value
	return o
}
func (o *CoverOptions) TiltOptimistic(optimistic bool) *CoverOptions {
	o.tiltOptimistic = optimistic
	return o
}
func (o *CoverOptions) TiltStatusTemplate(template string) *CoverOptions {
	o.tiltStatusTemplate = template
	return o
}
func (o *CoverOptions) TiltStatusFunc(f func() string) *CoverOptions {
	o.tiltStatusFunc = f
	return o
}
func (o *CoverOptions) UniqueId(id string) *CoverOptions {
	o.uniqueId = id
	return o
}
func (o *CoverOptions) ValueTemplate(template string) *CoverOptions {
	o.valueTemplate = template
	return o
}
