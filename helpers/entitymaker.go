package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"sort"
	"strings"
)

func generateEntities(devices []Device, external map[string]*jen.File) {

	for _, d := range devices {

		// /
		commandParams := make(map[string]struct{})
		camelName := strcase.ToCamel(d.Name)
		lowerCamelName := strcase.ToLowerCamel(d.Name)
		firstLetter := string(lowerCamelName[0])
		optionsFileName := fmt.Sprintf("%s_options", d.Name)
		optionsCamelName := fmt.Sprintf("%sOptions", strcase.ToCamel(d.Name))
		internalStateTypeName := fmt.Sprintf("%sState", lowerCamelName)
		externalStateTypeName := fmt.Sprintf("%sState", camelName)
		// /

		// Add standalone base level fields
		st := make(map[string]item)
		for _, key := range keyNames {
			if d.JSONContainer.Exists(key) {
				st[key] = item{
					main:   d.FieldAdder(key),
					setter: d.FunctionAdder(key),
				}
			}
		}

		if d.JSONContainer.Exists("device") {
			st["device"] = item{
				main: &statement{
					camelName: "Device",
					topic:     false,
					required:  true,
					name:      jen.Id(strcase.ToCamel("device")),
					t:         jen.Id("Device"),
					tag:       jen.Tag(map[string]string{"json": "device,omitempty"}),
					comment:   jen.Comment("Device configuration parameters"),
				},
			}

		}

		sortedKeys := []string{}
		for key := range st {
			sortedKeys = append(sortedKeys, key)
		}
		sort.Strings(sortedKeys)

		// Device MQTT Struct
		external[d.Name].Type().Id(strcase.ToCamel(d.Name)).StructFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					v := st[key]
					if v.main != nil {
						g.Add(v.main.combinePointer())
					}
					if v.setter != nil {
						g.Add(v.setter.combineLower())
					}
				}
				g.Id("MQTT").Op("*").Id("MQTTFields").Tag(map[string]string{"json": "-"}).Comment("MQTT configuration parameters")
				g.Id("states").Id(internalStateTypeName).Comment("Internal Holder of States")
				g.Id("States").Op("*").Id(externalStateTypeName).Tag(map[string]string{"json": "-"}).Comment("External state update location")
			},
		)

		stateObjects := make(map[string]struct{})

		for _, key := range sortedKeys {
			if key == "device" {
				continue
			}
			v := st[key]
			if v.main != nil && v.main.topic && !v.main.command {
				stateName := strings.TrimSuffix(v.main.stripTopic, "State")
				if stateName == "" {
					stateName = v.main.stripTopic
				}
				stateObjects[stateName] = struct{}{}
			}

		}

		external[d.Name].Func().
			Id(fmt.Sprintf("New%s", camelName)).
			Params(jen.Id("o").Op("*").Id(optionsCamelName)).
			Params(jen.Op("*").Id(camelName), jen.Error()).
			BlockFunc(func(g *jen.Group) {
				g.Add(jen.Var().Id(firstLetter).Id(camelName))
				g.Add(jen.Line())

				g.Add(jen.Id(firstLetter).Dot("States").Op("=").Op("&").Id("o").Dot("states"))

				for _, key := range sortedKeys {
					if key == "device" {
						continue
					}
					v := st[key]
					if v.main != nil && !v.main.topic {
						var el *jen.Statement
						switch v.main.lowerCamelName {
						case "name":
							el = jen.Else().Block(
								jen.Return(jen.Id("nil"), jen.Qual("fmt", "Errorf").Params(jen.Lit("name not defined"))),
							)
						case "uniqueId":
							el = jen.Else().Block(
								jen.Id("uniqueId").Op(":=").Qual("github.com/iancoleman/strcase", "ToDelimited").Params(jen.Id("o").Dot("name"), jen.Id("uint8(0x2d)")),
								jen.Id("uniqueId").Op("=").Qual("strings", "ReplaceAll").Params(jen.Id("uniqueId"), jen.Lit("'"), jen.Lit("_")),
								jen.Id(firstLetter).Dot(v.main.camelName).Op("=").Op("&").Id("uniqueId"),
							)
						}

						g.Add(
							jen.If(jen.Op("!").Qual("reflect", "ValueOf").Params(jen.Id("o").Dot(v.main.lowerCamelName)).Dot("IsZero").Params()).Block(
								jen.Id(firstLetter).Dot(v.main.camelName).Op("=").Op("&").Id("o").Dot(v.main.lowerCamelName),
							).Add(el),
						)
					}
					if v.setter != nil {
						var el *jen.Statement
						if v.setter.lowerCamelName == "stateFunc" {
							el = jen.Else().Block(
								jen.Id(firstLetter).Dot(v.setter.lowerCamelName).Op("=").Func().Params().String().Block(
									jen.Return(jen.Id(firstLetter).Dot("States").Dot("State")),
									// jen.Id("l").Dot("UpdateState").Params(),
								),
							)
						}

						if v.setter.lowerCamelName == "positionFunc" {
							el = jen.Else().Block(
								jen.Id(firstLetter).Dot(v.setter.lowerCamelName).Op("=").Func().Params().String().Block(
									jen.Return(jen.Id(firstLetter).Dot("States").Dot("Position")),
									// jen.Id("l").Dot("UpdateState").Params(),
								),
							)
						}

						if v.setter.lowerCamelName == "commandFunc" {
							if _, ok := st["state_topic"]; ok {
								el = jen.Else().Block(
									jen.Id(firstLetter).Dot(v.setter.lowerCamelName).Op("=").Func().Params(
										jen.Id("message").Qual("github.com/eclipse/paho.mqtt.golang", "Message"),
										jen.Id("client").Qual("github.com/eclipse/paho.mqtt.golang", "Client"),
									).Block(

										jen.Id("o").Dot("states").Dot("State").Op("=").String().Params(jen.Id("message").Dot("Payload").Params()),

										// jen.Id("l").Dot("UpdateState").Params(),
									),
								)
							}
						}
						// && v.main != nil && v.main.topic && !v.main.command
						var hasState bool
						var stateName string
						if v.main != nil {
							stateName = strings.TrimSuffix(v.main.stripTopic, "Command")
							if stateName == "" {
								stateName = "State"
							}
							_, hasState = stateObjects[stateName]
						}
						if v.setter.command && hasState {
							g.Add(
								jen.If(jen.Op("!").Qual("reflect", "ValueOf").Params(jen.Id("o").Dot(v.setter.lowerCamelName)).Dot("IsZero").Params()).Block(
									jen.Id(firstLetter).Dot(v.setter.lowerCamelName).Op("=").
										Func().Params(
										jen.Id("message").Qual("github.com/eclipse/paho.mqtt.golang", "Message"),
										jen.Id("client").Qual("github.com/eclipse/paho.mqtt.golang", "Client"),
									).Block(
										// jen.If(jen.Id("o").Dot("states").Dot(stateName).Op("==").String().Id("message").Dot("Payload").Params()),
										jen.If(jen.Id("o").Dot("states").Dot(stateName).Op("==").String().Params(jen.Id("message").Dot("Payload").Params())).Block(
											jen.Return(),
										),
										jen.Id("o").Dot("states").Dot(stateName).Op("=").String().Params(jen.Id("message").Dot("Payload").Params()),
										jen.Id(firstLetter).Dot("UpdateState").Params(),
										jen.Id("o").Dot(v.setter.lowerCamelName).Params(jen.Id("message"), jen.Id("client")),
									),
								).Add(el),
							)
						} else {
							g.Add(
								jen.If(jen.Op("!").Qual("reflect", "ValueOf").Params(jen.Id("o").Dot(v.setter.lowerCamelName)).Dot("IsZero").Params()).Block(
									jen.Id(firstLetter).Dot(v.setter.lowerCamelName).Op("=").Id("o").Dot(v.setter.lowerCamelName),
								).Add(el),
							)

						}

					}

				}
				// setup the state

				g.Add(jen.Return(jen.Op("&").Id(firstLetter), jen.Id("nil")))
			})

		external[d.Name].Type().Id(internalStateTypeName).StructFunc(func(g *jen.Group) {
			for _, key := range sortedKeys {
				if key == "device" {
					continue
				}
				v := st[key]
				if v.main != nil && v.main.topic && !v.main.command {
					stateName := strings.TrimSuffix(v.main.stripTopic, "State")
					if stateName == "" {
						stateName = v.main.stripTopic
					}
					g.Add(jen.Id(stateName).Op("*").Add(v.main.t))
				}

			}
		})
		external[d.Name].Type().Id(externalStateTypeName).StructFunc(func(g *jen.Group) {
			for _, key := range sortedKeys {
				if key == "device" || key == "availability_topic" {
					continue
				}
				v := st[key]
				if v.main != nil && v.main.topic && !v.main.command {
					stateName := strings.TrimSuffix(v.main.stripTopic, "State")
					if stateName == "" {
						stateName = v.main.stripTopic
					}
					g.Add(jen.Id(stateName).Add(v.main.t))
				}

			}
		})

		for _, key := range sortedKeys {
			if key == "device" || key == "availability_topic" {
				continue
			}
			v := st[key]
			if v.main != nil && v.main.topic && !v.main.command {
				stateName := strings.TrimSuffix(v.main.stripTopic, "State")
				if stateName == "" {
					stateName = v.main.stripTopic
				}
				external[d.Name].Func().
					Params(jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name))).
					Id(fmt.Sprintf("%s", stateName)).Params(jen.Id("s").String()).
					Block(
						jen.Id("d").Dot("States").Dot(stateName).Op("=").Id("s"),
						jen.Id("d").Dot("UpdateState").Params(),
					)
			}

		}

		// d.GetRawID()
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("GetRawId").Params().String().Block(
			jen.Return(jen.Lit(d.Name)),
		)

		// d.AddMessageHandler()
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("AddMessageHandler").Params().Block(
			jen.Id("d").Dot("MQTT").Dot("MessageHandler").Op("=").Id("MakeMessageHandler").Params(jen.Id("d")),
		)
		// d.GetUniqueID()
		if d.JSONContainer.Exists("unique_id") {
			external[d.Name].Func().Params(
				jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
			).Id("GetUniqueId").Params().String().Block(
				jen.Return(jen.Op("*").Id("d").Dot("UniqueId")),
			)
		} else {
			external[d.Name].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("GetUniqueId").Params().String().Block(
				jen.Return(jen.Lit("")),
			)
		}
		// d.GetDomainEntity()
		if d.JSONContainer.Exists("unique_id") {
			external[d.Name].Func().Params(
				jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
			).Id("GetDomainEntity").Params().String().Block(

				jen.Return(jen.Qual("fmt", "Sprintf").
					Params(jen.Lit(
						fmt.Sprintf("%s.%s", d.Name, "%s")),
						jen.Qual("strings", "ReplaceAll").Params(jen.Op("*").Id("d").Dot("UniqueId"), jen.Lit("-"), jen.Lit("_")),
					)),
			)
		} else {
			external[d.Name].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("GetDomainEntity").Params().String().Block(
				jen.Return(jen.Lit("")),
			)
		}

		// d.GetName()
		if d.JSONContainer.Exists("unique_id") {
			external[d.Name].Func().Params(
				jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
			).Id("GetName").Params().String().Block(
				jen.Return(jen.Op("*").Id("d").Dot("Name")),
			)
		} else {
			external[d.Name].Func().Params(
				jen.Id("d").Id(strcase.ToCamel(d.Name)),
			).Id("GetName").Params().String().Block(
				jen.Return(jen.Lit("")),
			)
		}

		// d.PopulateDevice()
		if d.JSONContainer.Exists("device") {
			external[d.Name].Func().Params(
				jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params(
				jen.Id("Manufacturer").String(),
				jen.Id("SoftwareName").String(),
				jen.Id("InstanceName").String(),
				jen.Id("SWVersion").String(),
				jen.Id("Identifier").String()).
				Block(
					jen.Id("d.Device.Manufacturer").Op("=&").Id("Manufacturer"),
					jen.Id("d.Device.Model").Op("=&").Id("SoftwareName"),
					jen.Id("d.Device.Name").Op("=&").Id("InstanceName"),
					jen.Id("d.Device.SwVersion").Op("=&").Id("SWVersion"),
					jen.Id("d.Device.Identifiers").Op("=&").Id("Identifier"),
				)
		} else {
			external[d.Name].Func().Params(
				jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
			).Id("PopulateDevice").Params(
				jen.Id("_").String(),
				jen.Id("_").String(),
				jen.Id("_").String(),
				jen.Id("_").String(),
				jen.Id("_").String()).
				Block()
		}

		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("UpdateState").Params(
		// jen.Id("state").Id("string"),
		).BlockFunc(
			func(g *jen.Group) {
				for _, key := range sortedKeys {
					if strings.HasSuffix(key, "topic") {
						if !IsCommand(key) {
							if key == "topic" {
								key = "state_topic"
							}
							trimmed := strings.TrimSuffix(strings.TrimSuffix(key, "topic"), "_")
							cam := strcase.ToCamel(key)

							camTrimmed := strcase.ToCamel(trimmed)
							if camTrimmed != "State" {
								camTrimmed = strings.TrimSuffix(camTrimmed, "State")
							}
							g.Add(
								jen.If(
									jen.Id("d").Dot(cam).Op("!=").Nil(),
								).Block(
									jen.Id("state").Op(":=").Id("d").Dot(strcase.ToLowerCamel(trimmed+"_func")).Params(),

									jen.If(
										jen.Id("d").Dot("states").Dot(camTrimmed).Op("==").Id("nil").
											Op("||").
											Id("state").Op("!=").Op("*").Id("d").Dot("states").Dot(camTrimmed).
											Op("||").
											Params(jen.Id("d").Dot("MQTT").Dot("ForceUpdate").Op("!=").Nil().Op("&&").Op("*").Id("d").Dot("MQTT").Dot("ForceUpdate")),
									).Block(
										jen.Id("token").Op(":=").Params(jen.Op("*").Id("d").Dot("MQTT").Dot("Client")).Dot("Publish").ParamsFunc(
											func(g *jen.Group) {
												g.Add(jen.Op("*").Id("d").Dot(cam))

												if d.JSONContainer.Exists("qos") {
													g.Add(jen.Byte().Params(jen.Op("*").Id("d").Dot("Qos")))
												} else {
													g.Add(jen.Lit(2))
												}

												// if d.JSONContainer.Exists("retain") && cam != "AvailabilityTopic" {
												//	g.Add(jen.Op("*").Id("d").Dot("Retain"))
												// } else {
												g.Add(jen.Lit(true))
												// }

												g.Add(jen.Id("state"))
											},
										),
										jen.Id("token").Dot("WaitTimeout").Params(jen.Qual("github.com/kjbreil/hass-mqtt/common", "WaitTimeout")),
										jen.Id("d").Dot("states").Dot(camTrimmed).Op("=").Op("&").Id("state"),
									),
								),
							)
						}
					}
				}
			},
		)

		// d.Subscribe()
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("Subscribe").Params().BlockFunc(
			func(g *jen.Group) {

				g.Add(
					jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
				)

				g.Add(
					jen.List(jen.Id("message"), jen.Err()).Op(":=").Qual("encoding/json", "Marshal").Params(jen.Id("d")),
				)
				g.Add(
					jen.If(
						jen.Id("err").Op("!=").Id("nil"),
					).Block(
						jen.Qual("log", "Fatal").Params(jen.Err()),
					),
				)

				for _, key := range sortedKeys {
					if strings.HasSuffix(key, "topic") {
						if IsCommand(key) {
							cam := strcase.ToCamel(key)

							g.Add(
								jen.If(
									jen.Id("d").Dot(cam).Op("!=").Nil(),
								).Block(
									jen.Id("t").Op(":=").Id("c").Dot("Subscribe").Params(
										jen.Op("*").Id("d").Dot(cam),
										jen.Lit(0),
										jen.Id("d").Dot("MQTT").Dot("MessageHandler"),
									),
									jen.Id("t").Dot("WaitTimeout").Params(jen.Qual("github.com/kjbreil/hass-mqtt/common", "WaitTimeout")),
									jen.If(
										jen.Id("t").Dot("Error").Params().Op("!=").Nil(),
									).Block(
										jen.Qual("log", "Fatal").Params(jen.Id("t").Dot("Error").Params()),
									),
								),
							)
						}
					}
				}

				g.Add(
					jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
						jen.Id("GetDiscoveryTopic").Params(jen.Id("d")),
						jen.Lit(2),
						jen.Lit(true),
						jen.Id("message"),
					),
				)

				g.Add(
					jen.Id("token").Dot("WaitTimeout").Params(jen.Qual("github.com/kjbreil/hass-mqtt/common", "WaitTimeout")),
				)

				for _, key := range keyNames {
					if d.JSONContainer.Exists(key) {
						if key == "availability_topic" {
							g.Add(
								jen.Id("d").Dot("availabilityFunc").Params(),
							)
						}
					}
				}

				g.Add(
					jen.Id("d").Dot("UpdateState").Params(),
				)

			},
		)

		// d.UnSubscribe()
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("UnSubscribe").Params().BlockFunc(
			func(g *jen.Group) {
				if d.JSONContainer.Exists("availability_topic") {
					g.Add(
						jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
					)

					g.Add(
						jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
							jen.Op("*").Id("d").Dot("AvailabilityTopic"),
							jen.Lit(2),
							jen.Lit(false),
							jen.Lit("offline"),
						),
					)

					g.Add(
						jen.Id("token").Dot("WaitTimeout").Params(jen.Qual("github.com/kjbreil/hass-mqtt/common", "WaitTimeout")),
					)

					for _, key := range sortedKeys {
						if strings.HasSuffix(key, "topic") {
							if IsCommand(key) {
								cam := strcase.ToCamel(key)

								g.Add(
									jen.If(
										jen.Id("d").Dot(cam).Op("!=").Nil(),
									).Block(
										jen.Id("t").Op(":=").Id("c").Dot("Unsubscribe").Params(
											jen.Op("*").Id("d").Dot(cam),
										),
										jen.Id("t").Dot("WaitTimeout").Params(jen.Qual("github.com/kjbreil/hass-mqtt/common", "WaitTimeout")),
										jen.If(
											jen.Id("t").Dot("Error").Params().Op("!=").Nil(),
										).Block(
											jen.Qual("log", "Fatal").Params(jen.Id("t").Dot("Error").Params()),
										),
									),
								)
							}
						}
					}
				}
			},
		)

		// d.AnnounceAvailable()
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("AnnounceAvailable").Params().BlockFunc(
			func(g *jen.Group) {
				if d.JSONContainer.Exists("availability_topic") {
					g.Add(
						jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
					)
					g.Add(
						jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
							jen.Op("*").Id("d").Dot("AvailabilityTopic"),
							jen.Lit(2),
							jen.Lit(true),
							jen.Lit("online"),
						),
					)
					g.Add(
						jen.Id("token").Dot("WaitTimeout").Params(jen.Qual("github.com/kjbreil/hass-mqtt/common", "WaitTimeout")),
					)
				}
			},
		)

		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("Initialize").Params().BlockFunc(func(g *jen.Group) {
			if d.JSONContainer.Exists("qos") {
				g.Add(
					jen.If(
						jen.Id("d").Dot("Qos").Op("==").Nil(),
					).Block(
						jen.Id("d").Dot("Qos").Op("=").New(jen.Int()),
						jen.Op("*").Id("d").Dot("Qos").Op("=").Int().Params(jen.Lit(0)),
					),
				)
			}
			if d.JSONContainer.Exists("retain") {
				g.Add(
					jen.If(
						jen.Id("d").Dot("Retain").Op("==").Nil(),
					).Block(
						jen.Id("d").Dot("Retain").Op("=").New(jen.Bool()),
						jen.Op("*").Id("d").Dot("Retain").Op("=").Lit(false),
					),
				)
			}
			if d.JSONContainer.Exists("unique_id") {
				g.Add(
					jen.If(jen.Id("d").Dot("UniqueId").Op("==").Nil()).BlockFunc(
						func(g *jen.Group) {
							g.Add(
								jen.Id("d").Dot("UniqueId").Op("=").New(jen.String()))
							g.Add(
								jen.Op("*").Id("d").Dot("UniqueId").Op("=").Qual("github.com/iancoleman/strcase", "ToDelimited").Params(
									jen.Op("*").Id("d").Dot("Name"),
									jen.Lit(uint8('-')),
								))

						},
					),
				)
			}
			if d.JSONContainer.Exists("availability_topic") {
				g.Add(
					jen.If(
						jen.Id("d").Dot("availabilityFunc").Op("==").Nil(),
					).Block(
						jen.Id("d").Dot("availabilityFunc").Op("=").Id("AvailabilityFunc"),
					),
				)
			}
			g.Add(
				jen.If(
					jen.Id("d").Dot("MQTT").Op("==").Nil(),
				).Block(
					jen.Id("d").Dot("MQTT").Op("=").New(jen.Id("MQTTFields")),
				),
			)
			g.Add(jen.Id("d").Dot("AddMessageHandler").Params())
			g.Add(jen.Id("d").Dot("populateTopics").Params())
		})
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("DeleteTopic").Params().BlockFunc(func(g *jen.Group) {
			g.Add(
				jen.Id("c").Op(":=").Op("*").Id("d").Dot("MQTT").Dot("Client"),
			)
			g.Add(
				jen.Id("token").Op(":=").Id("c").Dot("Publish").Params(
					jen.Id("GetDiscoveryTopic").Params(jen.Id("d")),
					jen.Lit(0),
					jen.Lit(true),
					jen.Lit(""),
				),
			)

			g.Add(
				jen.Id("token").Dot("WaitTimeout").Params(jen.Qual("github.com/kjbreil/hass-mqtt/common", "WaitTimeout")),
			)

			g.Add(
				jen.Qual("time", "Sleep").Params(jen.Qual("github.com/kjbreil/hass-mqtt/common", "HADiscoveryDelay")),
			)
		})

		// d.populateTopics()
		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("populateTopics").Params().BlockFunc(func(g *jen.Group) {
			for _, key := range sortedKeys {
				v := st[key]
				if v.setter != nil && v.main != nil {
					g.Add(
						jen.If(
							jen.Id("d").Dot(v.setter.lowerCamelName).Op("!=").Nil(),
						).BlockFunc(
							func(g *jen.Group) {
								g.Add(jen.Id("d").Dot(v.main.camelName).Op("=").New(jen.String()))
								g.Add(jen.Op("*").Id("d").Dot(v.main.camelName).Op("=").Id("GetTopic").Params(jen.Id("d"), jen.Lit(v.main.snakeName)))
								if IsCommand(v.main.snakeName) {
									g.Add(jen.Qual("github.com/kjbreil/hass-mqtt/common", "TopicStore").Index(
										jen.Op("*").Id("d").Dot(strcase.ToCamel(v.main.snakeName)),
									).Op("=").Id("&d").Dot(v.setter.lowerCamelName))
								}
							},
						),
					)
				}
			}

			// for _, name := range keyNames {
			//	if strings.HasSuffix(name, "topic") && d.JSONContainer.Exists(name) {
			//		if name == "topic" {
			//			name = "state_topic"
			//		}
			//		lName := strcase.ToCamel(strings.TrimSuffix(strings.TrimSuffix(name, "topic"), "_"))
			//		commandParams[lName] = struct{}{}
			//
			//		g.Add(
			//			jen.If(
			//				jen.Id("d").Dot(lName + "Func").Op("!=").Nil(),
			//			).BlockFunc(
			//				func(g *jen.Group) {
			//					g.Add(jen.Id("d").Dot(strcase.ToCamel(
			//						func(key string) string {
			//							var s string
			//							if key == "topic" {
			//								s = "state_topic"
			//							} else {
			//								s = key
			//							}
			//							return s
			//						}(name),
			//					)).Op("=").New(jen.String()))
			//					g.Add(jen.Op("*").Id("d").Dot(strcase.ToCamel(
			//						func(key string) string {
			//							var s string
			//							if key == "topic" {
			//								s = "state_topic"
			//							} else {
			//								s = key
			//							}
			//							return s
			//						}(name),
			//					)).Op("=").Id("GetTopic").Params(jen.Id("d"), jen.Lit(name)))
			//					if IsCommand(name, d) {
			//						g.Add(jen.Qual("github.com/kjbreil/hass-mqtt/common", "TopicStore").Index(
			//							jen.Op("*").Id("d").Dot(strcase.ToCamel(name)),
			//						).Op("=").Id("&d").Dot(lName + "Func"))
			//					}
			//				},
			//			),
			//		)
			//	}
			// }
		})

		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("SetMQTTFields").Params(
			jen.Id("fields").Id("MQTTFields"),
		).BlockFunc(
			func(g *jen.Group) {
				g.Add(
					jen.Op("*").Id("d").Dot("MQTT").Op("=").Id("fields"),
				)
			},
		)

		external[d.Name].Func().Params(
			jen.Id("d").Op("*").Id(strcase.ToCamel(d.Name)),
		).Id("GetMQTTFields").Params().Params(
			jen.Id("fields").Id("MQTTFields"),
		).BlockFunc(
			func(g *jen.Group) {
				g.Add(
					jen.Return(jen.Op("*").Id("d").Dot("MQTT")),
				)
			},
		)

		// Options File

		external[optionsFileName] = jen.NewFilePathName("./entities/"+optionsFileName+".go", "entities")
		external[optionsFileName].ImportAlias("github.com/eclipse/paho.mqtt.golang", "mqtt")
		external[optionsFileName].Comment("////////////////////////////////////////////////////////////////////////////////")
		external[optionsFileName].Comment("Do not modify this file, it is automatically generated")
		external[optionsFileName].Comment("////////////////////////////////////////////////////////////////////////////////")

		commandKeys := make([]string, len(commandParams))
		i := 0
		for k := range commandParams {
			commandKeys[i] = k
			i++
		}
		sort.Strings(commandKeys)

		opst := make(map[string]item)
		for _, key := range keyNames {
			if d.JSONContainer.Exists(key) {
				opst[key] = item{
					camelName: camelName,
					main:      d.FieldAdder(key),
					setter:    d.FunctionAdder(key),
				}
			}
		}

		optsSortedKeys := []string{}
		for key := range opst {
			optsSortedKeys = append(optsSortedKeys, key)
		}
		sort.Strings(optsSortedKeys)

		external[optionsFileName].Type().Id(optionsCamelName).StructFunc(func(g *jen.Group) {
			g.Id("states").Id(externalStateTypeName).Comment("External state update location")
			for _, key := range optsSortedKeys {
				v := opst[key]
				if v.main != nil && !v.main.topic {
					g.Add(v.main.combineLowerNoTag())
				}
				if v.setter != nil {
					g.Add(v.setter.combineLowerNoTag())
				}
			}
		},
		)

		external[optionsFileName].Func().Id(fmt.Sprintf("New%s", optionsCamelName)).
			Params().
			Op("*").Id(optionsCamelName).
			Block(
				jen.Return(jen.Op("&").Id(optionsCamelName).Values()),
			)

		external[optionsFileName].Func().
			Params(jen.Id("o").Op("*").Id(optionsCamelName)).
			Id("States").
			Params().
			Params(jen.Op("*").Id(externalStateTypeName)).
			Block(

				jen.Return(jen.Op("&").Id("o").Dot("states")),
			)

		for _, key := range optsSortedKeys {

			v := opst[key]
			split := strings.Split(key, "_")
			name := split[len(split)-1]
			switch name {
			case "default":
				name = "d"
			case "type":
				name = "t"
			}

			if v.main != nil && !v.main.topic {

				external[optionsFileName].Func().
					Params(jen.Id("o").Op("*").Id(optionsCamelName)).
					Id(fmt.Sprintf("%s", v.main.camelName)).
					Params(jen.Id(name).Add(v.main.t)).
					Params(jen.Op("*").Id(optionsCamelName)).
					Block(
						jen.Id("o").Dot(v.main.lowerCamelName).Op("=").Id(name),

						jen.Return(jen.Id("o")),
					)

			}
			if v.setter != nil {

				external[optionsFileName].Func().
					Params(jen.Id("o").Op("*").Id(optionsCamelName)).
					Id(fmt.Sprintf("%s", v.setter.camelName)).
					Params(jen.Id("f").Add(v.setter.t)).
					Params(jen.Op("*").Id(optionsCamelName)).
					Block(
						jen.Id("o").Dot(v.setter.lowerCamelName).Op("=").Id("f"),

						jen.Return(jen.Id("o")),
					)

				if strings.Contains(v.setter.camelName, "StateFunc") {
					stateName := strings.TrimSuffix(v.setter.camelName, "StateFunc")
					if stateName == "" {
						continue
					}
					commandKeyName := strings.ReplaceAll(key, "state", "command")
					if _, ok := opst[commandKeyName]; !ok {
						continue
					}

					commandFuncName := strings.ReplaceAll(v.setter.lowerCamelName, "State", "Command")

					external[optionsFileName].Func().
						Params(jen.Id("o").Op("*").Id(optionsCamelName)).
						Id(fmt.Sprintf("Enable%s", stateName)).
						Params().
						Params(jen.Op("*").Id(optionsCamelName)).
						Block(
							jen.Id("o").Dot(v.setter.lowerCamelName).Op("=").Func().Params().String().Block(
								jen.Return(jen.Id("o").Dot("states").Dot(stateName)),
							),
							jen.Id("o").Dot(commandFuncName).Op("=").Func().Params(
								jen.Id("message").Qual("github.com/eclipse/paho.mqtt.golang", "Message"),
								jen.Id("client").Qual("github.com/eclipse/paho.mqtt.golang", "Client"),
							).Block(
								jen.Id("o").Dot("states").Dot(stateName).Op("=").String().Params(jen.Id("message").Dot("Payload").Params()),
							),

							jen.Return(jen.Id("o")),
						)
				}
			}
		}

	}
}

type item struct {
	// statements []statement
	main      *statement
	setter    *statement
	camelName string
}

type statement struct {
	snakeName      string
	camelName      string
	lowerCamelName string
	stripTopic     string
	topic          bool
	command        bool
	required       bool
	name           *jen.Statement
	lowerName      *jen.Statement
	t              *jen.Statement
	tag            *jen.Statement
	comment        *jen.Statement
	firstLetter    string
}

func (s *statement) combine() *jen.Statement {
	return jen.Add(s.name, s.t, s.tag, s.comment)
}

func (s *statement) combineLower() *jen.Statement {
	return jen.Add(s.lowerName, s.t, s.comment)
}
func (s *statement) combinePointer() *jen.Statement {
	if s.required {
		return s.combine()
	}
	return jen.Add(s.name, jen.Op("*").Add(s.t), s.tag, s.comment)
}

func (s *statement) combineNoTag() *jen.Statement {
	return jen.Add(s.name, s.t, s.comment)
}
func (s *statement) combineLowerNoTag() *jen.Statement {
	return jen.Add(s.lowerName, s.t, s.comment)
}
