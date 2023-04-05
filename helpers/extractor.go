package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"reflect"
	"strings"
)

type entity struct {
	name     string
	device   bool
	config   map[string]entry
	hasState bool
	states   map[string]entry
}

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
