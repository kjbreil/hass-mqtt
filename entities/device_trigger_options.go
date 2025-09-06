package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type DeviceTriggerOptions struct {
	states         DeviceTriggerState // External state update location
	automationType string             // "The type of automation, must be 'trigger'."
	payload        string             // "Optional payload to match the payload being sent over the topic."
	platform       string             // "Must be `device_automation`. Only allowed and required in [MQTT auto discovery device messages](/integrations/mqtt/#device-discovery-payload)."
	qos            int                // "The maximum QoS level to be used when receiving and publishing messages."
	subtype        string             // "The subtype of the trigger, e.g. `button_1`. Entries supported by the frontend: `turn_on`, `turn_off`, `button_1`, `button_2`, `button_3`, `button_4`, `button_5`, `button_6`. If set to an unsupported value, will render as `subtype type`, e.g. `left_button pressed` with `type` set to `button_short_press` and `subtype` set to `left_button`"
	stateFunc      func() string
	entityType     string // "The type of the trigger, e.g. `button_short_press`. Entries supported by the frontend: `button_short_press`, `button_short_release`, `button_long_press`, `button_long_release`, `button_double_press`, `button_triple_press`, `button_quadruple_press`, `button_quintuple_press`. If set to an unsupported value, will render as `subtype type`, e.g. `button_1 spammed` with `type` set to `spammed` and `subtype` set to `button_1`"
	valueTemplate  string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) to extract the value."
}

func NewDeviceTriggerOptions() *DeviceTriggerOptions {
	return &DeviceTriggerOptions{}
}
func (o *DeviceTriggerOptions) States() *DeviceTriggerState {
	return &o.states
}
func (o *DeviceTriggerOptions) AutomationType(t string) *DeviceTriggerOptions {
	o.automationType = t
	return o
}
func (o *DeviceTriggerOptions) Payload(payload string) *DeviceTriggerOptions {
	o.payload = payload
	return o
}
func (o *DeviceTriggerOptions) Platform(platform string) *DeviceTriggerOptions {
	o.platform = platform
	return o
}
func (o *DeviceTriggerOptions) Qos(qos int) *DeviceTriggerOptions {
	o.qos = qos
	return o
}
func (o *DeviceTriggerOptions) Subtype(subtype string) *DeviceTriggerOptions {
	o.subtype = subtype
	return o
}
func (o *DeviceTriggerOptions) StateFunc(f func() string) *DeviceTriggerOptions {
	o.stateFunc = f
	return o
}
func (o *DeviceTriggerOptions) Type(t string) *DeviceTriggerOptions {
	o.entityType = t
	return o
}
func (o *DeviceTriggerOptions) ValueTemplate(template string) *DeviceTriggerOptions {
	o.valueTemplate = template
	return o
}
