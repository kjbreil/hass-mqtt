package entities

import (
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type LightRef struct {
	config lightRefConfig
	state  lightRefState
}

type lightRefConfig struct {
	Availability     []Topic `json:"availability"`
	AvailabilityMode string  `json:"availability_mode"`
	CommandTopic     Topic   `json:"command_topic"`
	Name             string  `json:"name"`
	Schema           string  `json:"schema"`
	StateTopic       string  `json:"state_topic"`
	UniqueId         string  `json:"unique_id"`
	Device           Device  `json:"device"`
	// Entity Specific Configurations
	Brightness      *bool    `json:"brightness,omitempty"`
	BrightnessScale *int     `json:"brightness_scale,omitempty"`
	ColorMode       *bool    `json:"color_mode,omitempty"`
	Effect          *bool    `json:"effect,omitempty"`
	EffectList      []string `json:"effect_list,omitempty"`
	MaxMireds       *int     `json:"max_mireds,omitempty"`
	MinMireds       *int     `json:"min_mireds,omitempty"`

	SupportedColorModes []string `json:"supported_color_modes"`
}

type lightRefState struct {
	State string `json:"state"`

	Brightness      int         `json:"brightness"`
	ColorMode       string      `json:"color_mode"`
	ColorTemp       int         `json:"color_temp"`
	LastSeen        time.Time   `json:"last_seen"`
	PowerOnBehavior interface{} `json:"power_on_behavior"`
}

type LightRefOptions struct {
	Brightness      bool
	BrightnessScale int
}

func NewLightRefOptions() *LightRefOptions {
	return &LightRefOptions{}
}

func (o *LightRefOptions) EnableBrightness(scale int) *LightRefOptions {
	o.Brightness = true
	o.BrightnessScale = scale
	return o
}

func NewLightRef(opts *LightRefOptions) {

}
