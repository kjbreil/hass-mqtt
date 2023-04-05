package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type TagOptions struct {
	States        TagState // External state update location
	StateFunc     func() string
	ValueTemplate string // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that returns a tag ID."
}

func NewTagOptions() *TagOptions {
	return &TagOptions{}
}
func (o *TagOptions) GetStates() *TagState {
	return &o.States
}
func (o *TagOptions) SetStateFunc(f func() string) *TagOptions {
	o.StateFunc = f
	return o
}
func (o *TagOptions) SetValueTemplate(template string) *TagOptions {
	o.ValueTemplate = template
	return o
}
