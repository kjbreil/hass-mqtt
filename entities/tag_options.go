package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type TagOptions struct {
	states        TagState // External state update location
	stateFunc     func() string
	valueTemplate string // "Defines a [template](/docs/configuration/templating/#using-value-templates-with-mqtt) that returns a tag ID."
}

func NewTagOptions() *TagOptions {
	return &TagOptions{}
}
func (o *TagOptions) States() *TagState {
	return &o.states
}
func (o *TagOptions) StateFunc(f func() string) *TagOptions {
	o.stateFunc = f
	return o
}
func (o *TagOptions) ValueTemplate(template string) *TagOptions {
	o.valueTemplate = template
	return o
}
