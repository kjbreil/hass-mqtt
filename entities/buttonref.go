package entities

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type ButtonRef struct {
	config ButtonRefConfig
}

type ButtonRefConfig struct {
	Availability     []Topic `json:"availability"`
	AvailabilityMode string  `json:"availability_mode"`
	CommandTopic     Topic   `json:"command_topic"`
	Name             string  `json:"name"`
	Schema           string  `json:"schema"`
	StateTopic       string  `json:"state_topic"`
	UniqueId         string  `json:"unique_id"`
	Device           Device  `json:"device"`
	// Entity Specific Configurations
}

type ButtonRefOptions struct {
	Brightness      bool
	BrightnessScale int
}

func NewButtonRefOptions() *ButtonRefOptions {
	return &ButtonRefOptions{}
}

func NewLight(opts *LightRefOptions) {

}
