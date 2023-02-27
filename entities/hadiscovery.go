package entities

// DiscoveryPrefix is the prefix that HA uses to discover on.
// Does not usually need changing.
const DiscoveryPrefix = "homeassistant"

// SWVersion is the software version.
var SWVersion = "0.0.1"

var SoftwareName = "Better MQTT Room"

// InstanceName is the instance name, helpful for identifying a given client
var InstanceName = "Better MQTT Room"

// NodeID is the Node ID, that is, what that node connects under.
var NodeID = "better-mqtt-room"

var Manufacturer = "Kjell Breiland"

///////////////////

var stateStore = initStore()
