package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"reflect"
	"strings"
)

// entity represents a Home Assistant entity with its configuration and state information.
type entity struct {
	name     string
	device   bool
	config   map[string]entry
	hasState bool
	states   map[string]entry
}

// entry represents a single configuration entry for an entity.
type entry struct {
	name     string
	t        reflect.Kind
	topic    bool
	required bool
	comment  string
}

var ignoreKeys = map[string]struct{}{
	"device":       struct{}{},
	"availability": struct{}{},
}

// extractEntity extracts entity information from a slice of Devices for code generation.
func extractEntity(devices []Device) []entity {
	var entities []entity
	for _, d := range devices {
		var e entity
		e.name = d.Name
		e.device = d.JSONContainer.Exists("device")
		e.hasState = d.JSONContainer.Exists("state_topic")
		e.config = make(map[string]entry)

		for k, v := range d.JSONContainer.ChildrenMap() {
			if _, ok := ignoreKeys[k]; ok {
				continue
			}
			fmt.Println(k, v)
			split := strings.Split(k, "_")
			if split[len(split)-1] == "topic" {
				e.config[strings.TrimSuffix(k, "_"+split[len(split)-1])] = entry{
					name:     strings.TrimSuffix(k, "_"+split[len(split)-1]),
					t:        dataType(v),
					topic:    true,
					required: isRequired(v),
					comment:  comment(v),
				}
				continue
			}

			if split[len(split)-1] == "template" {
				continue
			}

			e.config[k] = entry{
				name:     k,
				t:        dataType(v),
				topic:    false,
				required: isRequired(v),
				comment:  comment(v),
			}
		}

		entities = append(entities, e)

	}

	return entities
}

// isRequired checks if a configuration field is required, as indicated in its JSON definition.
func isRequired(c *gabs.Container) bool {
	req := c.Path("required")
	if req != nil {
		switch req.Data().(type) {
		case bool:
			return req.Data().(bool)
		}
	}

	return false
}

// dataType determines the Go reflect.Kind for a configuration field based on its JSON definition.
func dataType(c *gabs.Container) reflect.Kind {
	t := c.Path("type")
	if t != nil {
		switch t.Data().(type) {
		case string:
			switch t.Data().(string) {
			case "string", "device_class", "icon":
				return reflect.String
			case "integer":
				return reflect.Int
			case "boolean":
				return reflect.Bool
			case "list":
				return reflect.Slice
			default:
				panic(fmt.Sprintf("type not defined: %s", t.Data().(string)))
			}
		case []interface{}:
			return reflect.Slice
		}
	}
	return reflect.Invalid
}

// comment extracts the description comment from a configuration field's JSON definition, if present.
func comment(c *gabs.Container) string {
	req := c.Path("description")
	if req != nil {
		switch req.Data().(type) {
		case string:
			return req.Data().(string)
		}
	}

	return ""
}
