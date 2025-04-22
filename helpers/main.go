package main

import (
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

// main is the entry point for the code generation tool. It initializes devices, loads key names, and generates entity and type files for Home Assistant MQTT integration.
func main() {

	devices := DevicesInit()

	loadKeyNames()

	//entities := extractEntity(devices)

	external := make(map[string]*jen.File)

	fileList := []string{
		"types",
		"store",
	}

	for _, v := range append(DeviceNames, fileList...) {
		external[v] = jen.NewFilePathName("./entities/"+v+".go", "entities")
		external[v].ImportAlias("github.com/eclipse/paho.mqtt.golang", "mqtt")
		external[v].Comment("////////////////////////////////////////////////////////////////////////////////")
		external[v].Comment("Do not modify this file, it is automatically generated")
		external[v].Comment("////////////////////////////////////////////////////////////////////////////////")
	}

	fileList = []string{"types"}

	sort.Strings(keyNames)

	generateTypes(external)

	external["store"].Type().Id("StateStore").StructFunc(
		func(g *jen.Group) {
			for _, d := range devices {
				g.Add(
					jen.Id(strcase.ToCamel(d.Name)).StructFunc(
						func(h *jen.Group) {
							h.Add(jen.Id("Mutex").Op("*").Qual("sync", "Mutex"))
							for _, key := range keyNames {
								if d.JSONContainer.Exists(key) && strings.HasSuffix(key, "topic") {
									if !IsCommand(key) {
										h.Add(
											jen.Id(strcase.ToCamel(
												func(key string) string {
													var s string
													if key == "topic" {
														s = "State"
													} else {
														s = strings.TrimSuffix(strings.TrimSuffix(key, "topic"), "_")
													}
													return s
												}(key),
											)).Map(jen.String()).String(),
										)
									}
								}
							}
						},
					),
				)
			}
		},
	)

	external["store"].Func().Id("initStore").Params().Id("StateStore").BlockFunc(
		func(g *jen.Group) {
			g.Add(
				jen.Id("s").Op(":=").Id("StateStore").Block(),
			)
			for _, d := range devices {
				for _, key := range keyNames {
					if d.JSONContainer.Exists(key) && strings.HasSuffix(key, "topic") {
						if !IsCommand(key) {
							g.Add(
								jen.Id("s").Dot(strcase.ToCamel(d.Name)).Dot(strcase.ToCamel(
									func(key string) string {
										var s string
										if key == "topic" {
											s = "State"
										} else {
											s = strings.TrimSuffix(strings.TrimSuffix(key, "topic"), "_")
										}
										return s
									}(key),
								)).
									Op("=").
									Make(jen.Map(jen.String()).String()),
							)

						}
					}
				}

				g.Add(
					jen.Id("s").Dot(strcase.ToCamel(d.Name)).Dot("Mutex").
						Op("=").Op("&").Qual("sync", "Mutex").Values(),
				)
			}
			g.Add(
				jen.Return(jen.Id("s")),
			)
		},
	)

	generateEntities(devices, external)

	////////////////////////////////////////////////////////////////////////////////

	config := jen.NewFilePathName("./config/config.go", "config")
	config.ImportAlias("github.com/eclipse/paho.mqtt.golang", "mqtt")
	config.ImportAlias("github.com/kjbreil/hass-mqtt/devices", "devices")
	config.Comment("////////////////////////////////////////////////////////////////////////////////")
	config.Comment("Do not modify this file, it is automatically generated")
	config.Comment("////////////////////////////////////////////////////////////////////////////////")

	config.Type().Id("Config").StructFunc(
		func(g *jen.Group) {
			g.Add(
				jen.Id("MQTT").StructFunc(
					func(h *jen.Group) {
						p := []string{"broker", "username", "password", "node_id", "instance_name"}
						for _, n := range p {
							h.Add(jen.Id(strcase.ToCamel(n)).String().Tag(map[string]string{"json": n + ",omitempty"}))
						}
					},
				),
			)
			g.Add(
				jen.Id("Logging").StructFunc(
					func(h *jen.Group) {
						p := []string{"critical", "debug", "error", "warn", "mqtt"}
						for _, n := range p {
							h.Add(jen.Id(strcase.ToCamel(n)).Bool().Tag(map[string]string{"json": n + ",omitempty"}))
						}
					},
				),
			)
		},
	)

	config.Func().Params(
		jen.Id("c").Id("Config"),
	).Id("Translate").Params().Params(
		jen.Id("output").Index().Qual("github.com/kjbreil/hass-mqtt/devices", "Device"),
	).BlockFunc(
		func(g *jen.Group) {
			// g.Add(
			// 	jen.Id("output").Op(":=").Index().Qual("github.com/kjbreil/hass-mqtt/devices/externaldevice", "Device").Values(),
			// )
			for _, d := range devices {
				g.Add(
					jen.For(
						jen.List(
							jen.Id("_"),
							jen.Id("d"),
						),
					).Op(":=").Range().Id("c").Dot(strcase.ToCamel(d.Name)).Block(
						jen.Id("new"+strcase.ToCamel(d.Name)).Op(":=").Id("d").Dot("Translate").Params(),
						jen.Id("newDevice").Op(":=").Qual("github.com/kjbreil/hass-mqtt/devices/externaldevice", "Device").Params(jen.Op("&").Id("new"+strcase.ToCamel(d.Name))),
						jen.Id("output").Op("=").Append(jen.Id("output"), jen.Id("newDevice")),
					),
					// jen.Id(strcase.ToCamel(d.Name)).Index().Qual("github.com/kjbreil/hass-mqtt/devices/internaldevice", strcase.ToCamel(d.Name)),
				)
			}
			g.Add(
				jen.Return(),
			)
		},
	)

	////////////////////////////////////////////////////////////////////////////////

	for k, v := range external {
		err := v.Save("./entities/" + k + ".go")
		if err != nil {
			panic(err)
		}
	}

}
