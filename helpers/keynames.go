package main

import (
	"sort"
)

// keyNames holds all unique configuration keys found in devices.
var keyNames = []string{}

// blacklistKeys are keys to be excluded from code generation.
var blacklistKeys = []string{
	"device",
	"availability",
}

// loadKeyNames populates keyNames with all unique configuration keys from DeviceNames, excluding blacklistKeys.
func loadKeyNames() {
	for _, devicename := range DeviceNames {
		g := JsonExtractor(devicename)

		for k := range g.ChildrenMap() {
			doAdd := true
			for _, b := range blacklistKeys {
				if k == b {
					doAdd = false
				}
			}
			if doAdd {
				keyNames = append(keyNames, k)
			}
		}

	}

	singles := make(map[string]bool)

	for _, n := range keyNames {
		singles[n] = true
	}

	keyNames = []string{}

	for n := range singles {
		keyNames = append(keyNames, n)
	}

	sort.Strings(keyNames)

}
