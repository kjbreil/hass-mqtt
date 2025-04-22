package main

import (
	"strings"
)

// IsCommand determines whether a given key name represents a command topic, based on naming conventions and exclusions.
func IsCommand(name string) bool {

	retval := (!strings.Contains(name, "state")) &&
		(!strings.HasPrefix(name, "position")) &&
		(!strings.Contains(name, "availability")) &&
		// (!d.JSONContainer.Exists("set_" + name)) &&
		(!strings.Contains(name, "status")) &&
		(!strings.Contains(name, "current")) &&
		(!strings.Contains(name, "attributes")) &&
		(!strings.Contains(name, "action")) &&
		(name != "topic")

	return retval
}
