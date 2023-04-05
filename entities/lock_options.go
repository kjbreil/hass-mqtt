package entities

import mqtt "github.com/eclipse/paho.mqtt.golang"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type LockOptions struct {
	States                 LockState // External state update location
	AvailabilityMode       string    // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityFunc       func() string
	CodeFormat             string // "A regular expression to validate a supplied code when it is set during the service call to `open`, `lock` or `unlock` the MQTT lock."
	CommandTemplate        string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`. The lock command template accepts the parameters `value` and `code`. The `value` parameter will contain the configured value for either `payload_open`, `payload_lock` or `payload_unlock`. The `code` parameter is set during the service call to `open`, `lock` or `unlock` the MQTT lock and will be set `None` if no code was passed."
	CommandFunc            func(mqtt.Message, mqtt.Client)
	EnabledByDefault       bool   // "Flag which defines if the entity should be enabled when first added."
	Encoding               string // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   string // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesFunc     func() string
	Name                   string // "The name of the lock."
	ObjectId               string // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             bool   // "Flag that defines if lock works in optimistic mode."
	PayloadAvailable       string // "The payload that represents the available state."
	PayloadLock            string // "The payload sent to the lock to lock it."
	PayloadNotAvailable    string // "The payload that represents the unavailable state."
	PayloadOpen            string // "The payload sent to the lock to open it."
	PayloadUnlock          string // "The payload sent to the lock to unlock it."
	Qos                    int    // "The maximum QoS level of the state topic. It will also be used for messages published to command topic."
	Retain                 bool   // "If the published message should have the retain flag on or not."
	StateJammed            string // "The payload sent to `state_topic` by the lock when it's jammed."
	StateLocked            string // "The payload sent to `state_topic` by the lock when it's locked."
	StateLocking           string // "The payload sent to `state_topic` by the lock when it's locking."
	StateFunc              func() string
	StateUnlocked          string // "The payload sent to `state_topic` by the lock when it's unlocked."
	StateUnlocking         string // "The payload sent to `state_topic` by the lock when it's unlocking."
	UniqueId               string // "An ID that uniquely identifies this lock. If two locks have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a state value from the payload."
}

func NewLockOptions() *LockOptions {
	return &LockOptions{}
}
func (o *LockOptions) GetStates() *LockState {
	return &o.States
}
func (o *LockOptions) SetAvailabilityMode(mode string) *LockOptions {
	o.AvailabilityMode = mode
	return o
}
func (o *LockOptions) SetAvailabilityTemplate(template string) *LockOptions {
	o.AvailabilityTemplate = template
	return o
}
func (o *LockOptions) SetAvailabilityFunc(f func() string) *LockOptions {
	o.AvailabilityFunc = f
	return o
}
func (o *LockOptions) SetCodeFormat(format string) *LockOptions {
	o.CodeFormat = format
	return o
}
func (o *LockOptions) SetCommandTemplate(template string) *LockOptions {
	o.CommandTemplate = template
	return o
}
func (o *LockOptions) SetCommandFunc(f func(mqtt.Message, mqtt.Client)) *LockOptions {
	o.CommandFunc = f
	return o
}
func (o *LockOptions) SetEnabledByDefault(d bool) *LockOptions {
	o.EnabledByDefault = d
	return o
}
func (o *LockOptions) SetEncoding(encoding string) *LockOptions {
	o.Encoding = encoding
	return o
}
func (o *LockOptions) SetEntityCategory(category string) *LockOptions {
	o.EntityCategory = category
	return o
}
func (o *LockOptions) SetIcon(icon string) *LockOptions {
	o.Icon = icon
	return o
}
func (o *LockOptions) SetJsonAttributesTemplate(template string) *LockOptions {
	o.JsonAttributesTemplate = template
	return o
}
func (o *LockOptions) SetJsonAttributesFunc(f func() string) *LockOptions {
	o.JsonAttributesFunc = f
	return o
}
func (o *LockOptions) SetName(name string) *LockOptions {
	o.Name = name
	return o
}
func (o *LockOptions) SetObjectId(id string) *LockOptions {
	o.ObjectId = id
	return o
}
func (o *LockOptions) SetOptimistic(optimistic bool) *LockOptions {
	o.Optimistic = optimistic
	return o
}
func (o *LockOptions) SetPayloadAvailable(available string) *LockOptions {
	o.PayloadAvailable = available
	return o
}
func (o *LockOptions) SetPayloadLock(lock string) *LockOptions {
	o.PayloadLock = lock
	return o
}
func (o *LockOptions) SetPayloadNotAvailable(available string) *LockOptions {
	o.PayloadNotAvailable = available
	return o
}
func (o *LockOptions) SetPayloadOpen(open string) *LockOptions {
	o.PayloadOpen = open
	return o
}
func (o *LockOptions) SetPayloadUnlock(unlock string) *LockOptions {
	o.PayloadUnlock = unlock
	return o
}
func (o *LockOptions) SetQos(qos int) *LockOptions {
	o.Qos = qos
	return o
}
func (o *LockOptions) SetRetain(retain bool) *LockOptions {
	o.Retain = retain
	return o
}
func (o *LockOptions) SetStateJammed(jammed string) *LockOptions {
	o.StateJammed = jammed
	return o
}
func (o *LockOptions) SetStateLocked(locked string) *LockOptions {
	o.StateLocked = locked
	return o
}
func (o *LockOptions) SetStateLocking(locking string) *LockOptions {
	o.StateLocking = locking
	return o
}
func (o *LockOptions) SetStateFunc(f func() string) *LockOptions {
	o.StateFunc = f
	return o
}
func (o *LockOptions) SetStateUnlocked(unlocked string) *LockOptions {
	o.StateUnlocked = unlocked
	return o
}
func (o *LockOptions) SetStateUnlocking(unlocking string) *LockOptions {
	o.StateUnlocking = unlocking
	return o
}
func (o *LockOptions) SetUniqueId(id string) *LockOptions {
	o.UniqueId = id
	return o
}
func (o *LockOptions) SetValueTemplate(template string) *LockOptions {
	o.ValueTemplate = template
	return o
}
