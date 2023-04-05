package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type DeviceTriggerOptions struct {
	States         DeviceTriggerState // External state update location
	AutomationType string             // "The type of automation, must be 'trigger'."
	Payload        string             // "Optional payload to match the payload being sent over the topic."
	Qos            int                // "The maximum QoS level to be used when receiving messages."
	Subtype        string             // "The subtype of the trigger, e.g. `button_1`. Entries supported by the frontend: `turn_on`, `turn_off`, `button_1`, `button_2`, `button_3`, `button_4`, `button_5`, `button_6`. If set to an unsupported value, will render as `subtype type`, e.g. `left_button pressed` with `type` set to `button_short_press` and `subtype` set to `left_button`"
	StateFunc      func() string
	Type           string // "The type of the trigger, e.g. `button_short_press`. Entries supported by the frontend: `button_short_press`, `button_short_release`, `button_long_press`, `button_long_release`, `button_double_press`, `button_triple_press`, `button_quadruple_press`, `button_quintuple_press`. If set to an unsupported value, will render as `subtype type`, e.g. `button_1 spammed` with `type` set to `spammed` and `subtype` set to `button_1`"
	ValueTemplate  string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value."
}

func NewDeviceTriggerOptions() *DeviceTriggerOptions {
	return &DeviceTriggerOptions{}
}
func (o *DeviceTriggerOptions) GetStates() *DeviceTriggerState {
	return &o.States
}
func (o *DeviceTriggerOptions) SetAutomationType(t string) *DeviceTriggerOptions {
	o.AutomationType = t
	return o
}
func (o *DeviceTriggerOptions) SetPayload(payload string) *DeviceTriggerOptions {
	o.Payload = payload
	return o
}
func (o *DeviceTriggerOptions) SetQos(qos int) *DeviceTriggerOptions {
	o.Qos = qos
	return o
}
func (o *DeviceTriggerOptions) SetSubtype(subtype string) *DeviceTriggerOptions {
	o.Subtype = subtype
	return o
}
func (o *DeviceTriggerOptions) SetStateFunc(f func() string) *DeviceTriggerOptions {
	o.StateFunc = f
	return o
}
func (o *DeviceTriggerOptions) SetType(t string) *DeviceTriggerOptions {
	o.Type = t
	return o
}
func (o *DeviceTriggerOptions) SetValueTemplate(template string) *DeviceTriggerOptions {
	o.ValueTemplate = template
	return o
}
