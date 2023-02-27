package common

import (
	"time"
)

var (
	Retain                = true
	RetainClient          = false
	QoS              byte = 2
	HADiscoveryDelay      = 500 * time.Millisecond
	WaitTimeout           = 500 * time.Millisecond
)
