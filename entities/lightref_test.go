package entities

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_LightRefConfigJson(t *testing.T) {
	lrc := &LightRefConfig{
		Availability: []Topic{
			NewTopic("zigbee2mqtt/bridge/state", "{{ value_json.state }}"),
		},
		AvailabilityMode: "",
		CommandTopic: Topic{
			Topic:         "zigbee2mqtt/Office Door Other/set",
			ValueTemplate: nil,
		},
		Brightness:      false,
		BrightnessScale: 0,
		ColorMode:       false,
	}

	d, err := json.MarshalIndent(lrc, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(d))

}
