package main

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/dave/jennifer/jen"
	"github.com/ghodss/yaml"
	"github.com/iancoleman/strcase"
)

// var DeviceNames = []string{
//	//"button",
//	"light",
// }

var DeviceNames = []string{
	"alarm_control_panel",
	"binary_sensor",
	"button",
	"camera",
	"climate",
	"cover",
	"device_tracker",
	"device_trigger",
	"fan",
	"humidifier",
	"light",
	"lock",
	"number",
	"scene",
	"select",
	"sensor",
	"siren",
	"switch",
	"tag",
	"text",
	"update",
	"vacuum",
}

// Device represents a Home Assistant MQTT device with its name and JSON definition.
type Device struct {
	Name          string
	JSONContainer *gabs.Container
}

var PullNew = true

// DevicesInit initializes all devices listed in DeviceNames and returns them as a slice.
func DevicesInit() (retval []Device) {
	for _, name := range DeviceNames {
		d := Device{
			Name: name,
		}
		d.init()
		retval = append(retval, d)
	}
	return retval
}

// init populates the JSONContainer of the Device using its name.
func (dev *Device) init() {
	dev.JSONContainer = JsonExtractor(dev.Name)
}

// JsonExtractor retrieves the JSON definition for a given device name.
func JsonExtractor(deviceName string) *gabs.Container {
	yamlString, err := splitDocument(deviceName)
	if err != nil {
		log.Fatal(err)
	}

	jsonDoc, err := yaml.YAMLToJSON([]byte(yamlString))
	if err != nil {
		log.Fatal(err)
	}

	jsonParsed, err := gabs.ParseJSON(jsonDoc)
	if err != nil {
		log.Fatal(err)
	}

	return jsonParsed
}

// between returns the substring found between two specified substrings a and b in value.
func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		log.Println("Failed to find first match")
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		log.Println("Failed to find last match")
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		log.Println("First is after last")
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

// splitDocument fetches and extracts the configuration YAML section for a device from its documentation.
func splitDocument(devicename string) (string, error) {

	data, err := fetchDocument(devicename)
	if err != nil {
		return "", err
	}

	dat := string(data)

	if devicename == "vacuum" {
		idx := strings.Index(dat, "## State Configuration")
		if idx != -1 {
			dat = dat[idx:]
		}
	} else if devicename == "light" {
		dat = between(dat, "## Default schema - Configuration", "## Default schema - Examples")
	}

	match := between(dat, "{% configuration %}", "{% endconfiguration %}")

	targetFile := "./helpers/cache/" + devicename + "_split.yaml"

	err = os.WriteFile(targetFile, []byte(match), fs.FileMode(0644))
	if err != nil {
		return "nil", err
	}

	return match, nil

}

// exists checks if a file exists at the given path.
func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// fetchDocument retrieves the raw markdown documentation for a device, caching it locally if not already present.
func fetchDocument(devicename string) ([]byte, error) {

	url := "https://raw.githubusercontent.com/home-assistant/home-assistant.io/next/source/_integrations/" + devicename + ".mqtt.markdown"

	targetFile := "./helpers/cache/" + devicename + ".md"

	if exists(targetFile) && PullNew == false {

		data, err := os.ReadFile(targetFile)
		if err == nil {
			return data, nil
		}

		log.Println("Loading " + targetFile + " failed")

	}

	log.Println("Fetching " + devicename)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(targetFile, bodyBytes, fs.FileMode(0644))
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil

}

// Unquote removes surrounding quotes from a string, if present.
func Unquote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

// TypeTranslator maps a string type name to the corresponding jen.Statement for code generation.
func TypeTranslator(t string, s *jen.Statement) *jen.Statement {

	v := Unquote(t)
	switch v {
	case "string", "template", "icon", "device_class":
		return s.String()
	case "integer":
		return s.Int()
	case "boolean":
		return s.Bool()
	case "list", "[\"list\"]", "[\"string\",\"list\"]":
		return s.Params(jen.Index().String())
	case "map":
		return s.Params(jen.Map(jen.String()).String())
	case "float":
		return s.Float64()
	default:
		log.Fatalln("Unknown type " + t)
		return nil
	}
}

// FieldAdder creates a statement representing a struct field for a given key in the device's JSON definition.
func (dev *Device) FieldAdder(key string) *statement {
	snakeName := func(key string) string {
		var s string
		if key == "topic" {
			s = "state_topic"
		} else {
			s = key
		}
		return s
	}(key)

	camelName := strcase.ToCamel(snakeName)
	lowerCamelName := strcase.ToLowerCamel(camelName)
	stripTopic := strings.TrimSuffix(camelName, "Topic")
	dat := dev.JSONContainer.ChildrenMap()

	t := Unquote(dat[key].ChildrenMap()["type"].String())
	if lowerCamelName == "type" {
		lowerCamelName = "entityType"
	}

	return &statement{
		name:           jen.Id(camelName),
		snakeName:      snakeName,
		camelName:      camelName,
		lowerCamelName: lowerCamelName,
		stripTopic:     stripTopic,
		lowerName:      jen.Id(lowerCamelName),
		topic:          strings.HasSuffix(key, "topic"),
		command:        IsCommand(key),
		t:              TypeTranslator(t, jen.Add()),
		comment:        jen.Comment(dev.JSONContainer.Path(key + ".description").String()),
		tag:            jen.Tag(map[string]string{"json": key + ",omitempty"}),
	}
}

// FunctionAdder creates a statement representing a function for a given key in the device's JSON definition.
func (dev *Device) FunctionAdder(key string) *statement {

	if strings.HasSuffix(key, "topic") && (key != "availability") {
		retval := &statement{
			name:    nil,
			t:       nil,
			comment: nil,
		}
		if key == "topic" {
			key = "state_topic"
		}

		nk := strings.TrimSuffix(key, "topic")
		nk = strings.TrimSuffix(nk, "_")

		retval.camelName = strcase.ToCamel(nk) + "Func"
		retval.lowerCamelName = strcase.ToLowerCamel(nk) + "Func"
		retval.snakeName = nk
		retval.firstLetter = string(retval.lowerCamelName[0])
		retval.command = IsCommand(key)
		retval.name = jen.Id(retval.camelName)
		retval.lowerName = jen.Id(retval.lowerCamelName)
		retval.t = jen.Func()

		if IsCommand(key) {
			retval.t.Params(
				jen.Qual("github.com/eclipse/paho.mqtt.golang", "Message"),
				jen.Qual("github.com/eclipse/paho.mqtt.golang", "Client"),
			)
		} else {
			retval.t.Params().String()
		}

		retval.tag = jen.Tag(map[string]string{"json": "-"})
		return retval
	}

	return nil

}
