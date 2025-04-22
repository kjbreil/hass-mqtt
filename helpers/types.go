package main

import (
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

// generateTypes generates Go interface and struct types for Home Assistant MQTT entities and devices.
func generateTypes(external map[string]*jen.File) {
	external["types"].Type().Id("Entity").Interface(
		// jen.UnionFunc(
		// 	func(g *jen.Group) {
		// 		for _, d := range devices {
		// 			g.Add(jen.Id(strcase.ToCamel(d.Name)))
		// 		}
		// 	},
		// ),
		jen.Id("GetRawId").Params().String(),
		jen.Id("GetUniqueId").Params().String(),
		jen.Id("GetDomainEntity").Params().String(),
		jen.Id("GetName").Params().String(),
		jen.Id("PopulateDevice").Params(jen.Id("Manufacturer").String(), jen.Id("SoftwareName").String(), jen.Id("InstanceName").String(), jen.Id("SWVersion").String(), jen.Id("Identifier").String()),
		jen.Id("UpdateState").Params(),
		jen.Id("Subscribe").Params(),
		jen.Id("UnSubscribe").Params(),
		jen.Id("AddMessageHandler").Params(),
		jen.Id("DeleteTopic").Params(),
		jen.Id("Initialize").Params(),
		jen.Id("SetMQTTFields").Params(
			jen.Id("MQTTFields"),
		),
		jen.Id("GetMQTTFields").Params().Id("MQTTFields"),
	)

	// Device type
	// TOOD: Determine required fields
	external["types"].Type().Id("Device").StructFunc(
		func(g *jen.Group) {
			for _, v := range []string{
				"configuration_url",
				"connections",
				"identifiers",
				"manufacturer",
				"model",
				"name",
				"suggested_area",
				"sw_version",
				"via_device",
			} {
				g.Add(
					jen.Id(strcase.ToCamel(v)).Op("*").String().Tag(map[string]string{"json": v + ",omitempty"}),
					//.Comment(d.JSONContainer.Path("device.keys." + v + ".description").String()),
				)
			}
		},
	)

	external["types"].Type().Id("Topic").Struct(
		jen.Id("Topic").String().Tag(map[string]string{"json": "topic"}),
		jen.Id("ValueTemplate").Op("*").String().Tag(map[string]string{"json": "value_template"}),
	)

	external["types"].Func().
		Params(jen.Id("t").Id("Topic")).
		Id("MarshalJSON").Params().
		Params(jen.Index().Byte(), jen.Error()).
		Block(
			jen.If(
				jen.Id("t").Dot("ValueTemplate").Op("==").Id("nil")).Block(
				jen.Return(jen.Qual("encoding/json", "Marshal").Params(jen.Id("t").Dot("Topic"))),
			),
			jen.Return(
				jen.Qual("encoding/json", "Marshal").Params(jen.Struct(
					jen.Id("Topic").String().Tag(map[string]string{"json": "topic"}),
					jen.Id("ValueTemplate").Op("*").String().Tag(map[string]string{"json": "value_template"}),
				).Values(
					jen.Dict{
						jen.Id("Topic"):         jen.Id("t").Dot("Topic"),
						jen.Id("ValueTemplate"): jen.Id("t").Dot("ValueTemplate"),
					},
				)),
			),
		)

	external["types"].Func().
		Id("NewTopic").
		Params(jen.Id("topic").String(), jen.Id("valueTemplate").String()).
		Id("Topic").
		Block(
			jen.If(jen.Id("valueTemplate").Op("==").Lit("")).Block(
				jen.Return(jen.Id("Topic").
					Values(
						jen.Dict{
							jen.Id("Topic"):         jen.Id("topic"),
							jen.Id("ValueTemplate"): jen.Id("nil"),
						},
					),
				),
			),
			jen.Return(jen.Id("Topic").
				Values(
					jen.Dict{
						jen.Id("Topic"):         jen.Id("topic"),
						jen.Id("ValueTemplate"): jen.Op("&").Id("valueTemplate"),
					},
				),
			),
		)
}

//type Topic struct {
//	Topic         string  `json:"topic"`
//	ValueTemplate *string `json:"value_template"`
//}
//
//func NewTopic(topic string, valueTemplate string) Topic {
//	if valueTemplate == "" {
//		return Topic{
//			Topic:         topic,
//			ValueTemplate: nil,
//		}
//	}
//	return Topic{
//		Topic:         topic,
//		ValueTemplate: &valueTemplate,
//	}
//}

//func (t Topic) MarshalJSON() ([]byte, error) {
//	if t.ValueTemplate == nil {
//		return json.Marshal(t.Topic)
//	}
//	return json.Marshal(struct {
//		Topic         string  `json:"topic"`
//		ValueTemplate *string `json:"value_template"`
//	}{
//		Topic:         t.Topic,
//		ValueTemplate: t.ValueTemplate,
//	})
//}
