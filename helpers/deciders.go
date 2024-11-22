package main

import (
	"strings"
)

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
