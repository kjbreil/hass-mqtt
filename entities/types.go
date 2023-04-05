package entities

import "encoding/json"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Entity interface {
	GetRawId() string
	GetUniqueId() string
	GetName() string
	PopulateDevice(Manufacturer string, SoftwareName string, InstanceName string, SWVersion string, Identifier string)
	PopulateTopics()
	UpdateState()
	Subscribe()
	UnSubscribe()
	AddMessageHandler()
	DeleteTopic()
	Initialize()
	SetMQTTFields(MQTTFields)
	GetMQTTFields() MQTTFields
}
type Device struct {
	ConfigurationUrl *string `json:"configuration_url,omitempty"`
	Connections      *string `json:"connections,omitempty"`
	Identifiers      *string `json:"identifiers,omitempty"`
	Manufacturer     *string `json:"manufacturer,omitempty"`
	Model            *string `json:"model,omitempty"`
	Name             *string `json:"name,omitempty"`
	SuggestedArea    *string `json:"suggested_area,omitempty"`
	SwVersion        *string `json:"sw_version,omitempty"`
	ViaDevice        *string `json:"via_device,omitempty"`
}
type Topic struct {
	Topic         string  `json:"topic"`
	ValueTemplate *string `json:"value_template"`
}

func (t Topic) MarshalJSON() ([]byte, error) {
	if t.ValueTemplate == nil {
		return json.Marshal(t.Topic)
	}
	return json.Marshal(struct {
		Topic         string  `json:"topic"`
		ValueTemplate *string `json:"value_template"`
	}{
		Topic:         t.Topic,
		ValueTemplate: t.ValueTemplate,
	})
}
func NewTopic(topic string, valueTemplate string) Topic {
	if valueTemplate == "" {
		return Topic{
			Topic:         topic,
			ValueTemplate: nil,
		}
	}
	return Topic{
		Topic:         topic,
		ValueTemplate: &valueTemplate,
	}
}
