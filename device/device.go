package device

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kjbreil/hass-mqtt/entities"
)

type Device struct {
	ConfigurationUrl *string `json:"configuration_url,omitempty"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
	Connections      *string `json:"connections,omitempty"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`."
	Identifiers      *string `json:"identifiers,omitempty"`       // "A list of IDs that uniquely identify the device. For example a serial number."
	Manufacturer     *string `json:"manufacturer,omitempty"`      // "The manufacturer of the device."
	Model            *string `json:"model,omitempty"`             // "The model of the device."
	Name             *string `json:"name,omitempty"`              // "The name of the device."
	SuggestedArea    *string `json:"suggested_area,omitempty"`    // "Suggest an area if the device isn’t in one yet."
	SwVersion        *string `json:"sw_version,omitempty"`        // "The firmware version of the device."
	ViaDevice        *string `json:"via_device,omitempty"`        // "Identifier of a device that routes messages between this device and Home Assistant. Examples of such devices are hubs, or parent devices of a sub-device. This is used to show device topology in Home Assistant."

	entities map[string]entities.Entity
}

func New(name, identifier, model, manufacturer, swVersion string) *Device {
	return &Device{
		Identifiers:  &identifier,
		Manufacturer: &manufacturer,
		Model:        &model,
		Name:         &name,
		SwVersion:    &swVersion,
		entities:     make(map[string]entities.Entity),
	}
}

func (d *Device) Add(entity entities.Entity) error {
	if _, ok := d.entities[entity.GetUniqueId()]; ok {
		return fmt.Errorf("%s entity already exists", entity.GetUniqueId())
	}
	d.entities[entity.GetUniqueId()] = entity
	return nil
}

func (d *Device) Initialize() error {
	if len(d.entities) == 0 {
		return fmt.Errorf("device %s has no entities", *d.Name)
	}
	for _, e := range d.entities {
		e.Initialize()
		e.PopulateDevice(
			*d.Manufacturer,
			*d.Name,
			*d.Name,
			*d.SwVersion,
			*d.Identifiers,
		)
	}
	return nil
}

func (d *Device) Subscribe() {
	for _, e := range d.entities {
		e.Subscribe()
	}
}

func (d *Device) SetMQTTFields(client mqtt.Client) {
	for _, e := range d.entities {
		f := e.GetMQTTFields()
		f.Client = &client
		e.SetMQTTFields(f)
	}
}

func (d *Device) GetUniqueId() string {
	return *d.Identifiers
}

func (d *Device) UnSubscribe() {
	for _, e := range d.entities {
		e.UnSubscribe()
	}
}
