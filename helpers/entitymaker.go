package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"reflect"
	"sort"
)

func generateEntities(entities []entity, external map[string]*jen.File) {

	for _, e := range entities {
		file := external[e.name]

		camelName := strcase.ToCamel(e.name)
		configType := fmt.Sprintf("%sConfig", camelName)
		stateType := fmt.Sprintf("%sState", camelName)
		optionsType := fmt.Sprintf("%sOptions", camelName)

		file.Type().Id(camelName).StructFunc(func(g *jen.Group) {
			g.Add(jen.Id("config").Id(configType))
			if e.hasState {
				g.Add(jen.Id("state").Id(stateType))
			}
		},
		)

		configKeys := make([]string, len(e.config))
		i := 0
		for k := range e.config {
			configKeys[i] = k
			i++
		}
		sort.Strings(configKeys)

		file.Type().Id(configType).StructFunc(func(g *jen.Group) {
			for _, k := range configKeys {
				c, ok := e.config[k]
				if !ok {
					continue
				}
				entryCamel := strcase.ToCamel(c.name)

				jenType := jen.Add()
				if !c.required {
					jenType = jen.Op("*")
				}

				if c.topic {
					jenType = jenType.Add(jen.Id("Topic"))
				} else {
					switch c.t {
					case reflect.String:
						jenType = jenType.Add(jen.String())
					case reflect.Bool:
						jenType = jenType.Add(jen.Bool())
					case reflect.Int:
						jenType = jenType.Add(jen.Int())
					case reflect.Slice:
						jenType = jen.Index().String()
					default:
						panic(fmt.Sprintf("type not setup: %s", c.t))
					}
				}

				g.Add(jen.Id(entryCamel).Add(jenType)).Comment(c.comment)
			}
		})

		if e.hasState {
			file.Type().Id(stateType).StructFunc(func(g *jen.Group) {

			})
		}

		file.Type().Id(optionsType).StructFunc(func(g *jen.Group) {

		})
	}
}
