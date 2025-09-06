package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type LockOptions struct {
	states                 LockState // External state update location
	availabilityMode       string    // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	availabilityTemplate   string    // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	availabilityFunc       func() string
	codeFormat             string // "A regular expression to validate a supplied code when it is set during the action to `open`, `lock` or `unlock` the MQTT lock."
	commandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-command-templates-with-mqtt) to generate the payload to send to `command_topic`. The lock command template accepts the parameters `value` and `code`. The `value` parameter will contain the configured value for either `payload_open`, `payload_lock` or `payload_unlock`. The `code` parameter is set during the action to `open`, `lock` or `unlock` the MQTT lock and will be set `None` if no code was passed."
	commandFunc            func(mqtt.Message, mqtt.Client)
	enabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	entityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	entityPicture          string // "Picture URL for the entity."
	icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	jsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	jsonAttributesFunc     func() string
	name                   string // "The name of the lock. Can be set to `null` if only the device name is relevant."
	objectId               string // "Used `object_id` instead of `name` for automatic generation of `entity_id`. This only works when the entity is added for the first time. When set, this overrides a user-customized Entity ID in case the entity was deleted and added again."
	optimistic             bool   // "Flag that defines if lock works in optimistic mode."
	payloadAvailable       string // "The payload that represents the available state."
	payloadLock            string // "The payload sent to the lock to lock it."
	payloadNotAvailable    string // "The payload that represents the unavailable state."
	payloadOpen            string // "The payload sent to the lock to open it."
	payloadReset           string // "A special payload that resets the state to `unknown` when received on the `state_topic`."
	payloadUnlock          string // "The payload sent to the lock to unlock it."
	platform               string // "Must be `lock`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos                    int    // "The maximum QoS level to be used when receiving and publishing messages."
	retain                 bool   // "If the published message should have the retain flag on or not."
	stateJammed            string // "The payload sent to `state_topic` by the lock when it's jammed."
	stateLocked            string // "The payload sent to `state_topic` by the lock when it's locked."
	stateLocking           string // "The payload sent to `state_topic` by the lock when it's locking."
	stateFunc              func() string
	stateUnlocked          string // "The payload sent to `state_topic` by the lock when it's unlocked."
	stateUnlocking         string // "The payload sent to `state_topic` by the lock when it's unlocking."
	uniqueId               string // "An ID that uniquely identifies this lock. If two locks have the same unique ID, Home Assistant will raise an exception. Required when used with device-based discovery."
	valueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract a state value from the payload."
}

func NewLockOptions() *LockOptions {
	return &LockOptions{}
}
func (o *LockOptions) States() *LockState {
	return &o.states
}
func (o *LockOptions) AvailabilityMode(mode string) *LockOptions {
	o.availabilityMode = mode
	return o
}
func (o *LockOptions) AvailabilityTemplate(template string) *LockOptions {
	o.availabilityTemplate = template
	return o
}
func (o *LockOptions) AvailabilityFunc(f func() string) *LockOptions {
	o.availabilityFunc = f
	return o
}
func (o *LockOptions) CodeFormat(format string) *LockOptions {
	o.codeFormat = format
	return o
}
func (o *LockOptions) CommandTemplate(template string) *LockOptions {
	o.commandTemplate = template
	return o
}
func (o *LockOptions) CommandFunc(f func(mqtt.Message, mqtt.Client)) *LockOptions {
	o.commandFunc = f
	return o
}
func (o *LockOptions) EnabledByDefault(d bool) *LockOptions {
	o.enabledByDefault = d
	return o
}
func (o *LockOptions) Encoding(encoding string) *LockOptions {
	o.encoding = encoding
	return o
}
func (o *LockOptions) EntityCategory(category string) *LockOptions {
	o.entityCategory = category
	return o
}
func (o *LockOptions) EntityPicture(picture string) *LockOptions {
	o.entityPicture = picture
	return o
}
func (o *LockOptions) Icon(icon string) *LockOptions {
	o.icon = icon
	return o
}
func (o *LockOptions) JsonAttributesTemplate(template string) *LockOptions {
	o.jsonAttributesTemplate = template
	return o
}
func (o *LockOptions) JsonAttributesFunc(f func() string) *LockOptions {
	o.jsonAttributesFunc = f
	return o
}
func (o *LockOptions) Name(name string) *LockOptions {
	o.name = name
	return o
}
func (o *LockOptions) ObjectId(id string) *LockOptions {
	o.objectId = id
	return o
}
func (o *LockOptions) Optimistic(optimistic bool) *LockOptions {
	o.optimistic = optimistic
	return o
}
func (o *LockOptions) PayloadAvailable(available string) *LockOptions {
	o.payloadAvailable = available
	return o
}
func (o *LockOptions) PayloadLock(lock string) *LockOptions {
	o.payloadLock = lock
	return o
}
func (o *LockOptions) PayloadNotAvailable(available string) *LockOptions {
	o.payloadNotAvailable = available
	return o
}
func (o *LockOptions) PayloadOpen(open string) *LockOptions {
	o.payloadOpen = open
	return o
}
func (o *LockOptions) PayloadReset(reset string) *LockOptions {
	o.payloadReset = reset
	return o
}
func (o *LockOptions) PayloadUnlock(unlock string) *LockOptions {
	o.payloadUnlock = unlock
	return o
}
func (o *LockOptions) Platform(platform string) *LockOptions {
	o.platform = platform
	return o
}
func (o *LockOptions) Qos(qos int) *LockOptions {
	o.qos = qos
	return o
}
func (o *LockOptions) Retain(retain bool) *LockOptions {
	o.retain = retain
	return o
}
func (o *LockOptions) StateJammed(jammed string) *LockOptions {
	o.stateJammed = jammed
	return o
}
func (o *LockOptions) StateLocked(locked string) *LockOptions {
	o.stateLocked = locked
	return o
}
func (o *LockOptions) StateLocking(locking string) *LockOptions {
	o.stateLocking = locking
	return o
}
func (o *LockOptions) StateFunc(f func() string) *LockOptions {
	o.stateFunc = f
	return o
}
func (o *LockOptions) StateUnlocked(unlocked string) *LockOptions {
	o.stateUnlocked = unlocked
	return o
}
func (o *LockOptions) StateUnlocking(unlocking string) *LockOptions {
	o.stateUnlocking = unlocking
	return o
}
func (o *LockOptions) UniqueId(id string) *LockOptions {
	o.uniqueId = id
	return o
}
func (o *LockOptions) ValueTemplate(template string) *LockOptions {
	o.valueTemplate = template
	return o
}
