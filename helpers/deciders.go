package main

import (
	"strings"
)

func IsCommand(name string) bool {

	retval := (!strings.Contains(name, "state")) &&
		(!strings.Contains(name, "availability")) &&
		//(!d.JSONContainer.Exists("set_" + name)) &&
		(!strings.Contains(name, "status")) &&
		(!strings.Contains(name, "current")) &&
		(!strings.Contains(name, "attributes")) &&
		(name != "topic")

	return retval
}
