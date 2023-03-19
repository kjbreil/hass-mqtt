package entities

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
