package example

import (
	"encoding/json"
	"fmt"

	"github.com/alvinmatias69/gql-doc/entity"
	"github.com/rs/zerolog/log"
)

func (g *Generator) Generate() (entity.Spec, error) {
	var err error

	log.Info().Msg("generating query example")
	for idx := range g.spec.Queries {
		err = g.generate(&g.spec.Queries[idx])
		if err != nil {
			return g.spec, err
		}
	}

	log.Info().Msg("generating mutation example")
	for idx := range g.spec.Mutations {
		err = g.generate(&g.spec.Mutations[idx])
		if err != nil {
			return g.spec, err
		}
	}

	return g.spec, nil
}

func (g *Generator) generate(function *entity.Property) error {
	var (
		response = make(map[string]interface{})
	)

	response[function.Name] = g.mapType(function.Type, function.IsList)
	resString, err := json.MarshalIndent(response, "", g.indent)
	if err != nil {
		return fmt.Errorf("error generating response example: %v", err)
	}

	request, err := g.buildQueryRequest(function)
	if err != nil {
		return err
	}

	function.Example = entity.Example{
		Request:  request,
		Response: string(resString),
	}

	return nil
}

func (g *Generator) mapType(types string, isList bool) (final interface{}) {
	defer func() {
		if isList {
			final = []interface{}{final}
		}
	}()

	if _, ok := entity.ScalarTypes[types]; ok {
		final = entity.Default[types]
		return
	}

	var (
		definition = g.definitions[types]
		result     = make(map[string]interface{})
	)

	for _, prop := range definition.Properties {
		if prop.IsScalar {
			result[prop.Name] = entity.Default[prop.Type]
			if prop.IsList {
				result[prop.Name] = []interface{}{entity.Default[prop.Type]}
			}
			continue
		}

		result[prop.Name] = g.mapType(prop.Type, prop.IsList)
	}

	final = result
	return
}

func (g *Generator) buildQueryRequest(function *entity.Property) (string, error) {
	var (
		result  = function.Name + " "
		request = make(map[string]interface{})
		qBody   = g.mapQuery(function.Type, 0) + g.prefix
	)

	if len(function.Parameters) == 0 {
		return result + qBody, nil
	}

	for _, param := range function.Parameters {
		request[param.Name] = g.mapType(param.Type, param.IsList)
	}
	reqByte, err := json.MarshalIndent(request, "", g.indent)
	if err != nil {
		return "", err
	}

	reqString := g.queryRe.ReplaceAllString(string(reqByte), `$1:`)
	reqString = "(" + reqString[1:len(reqString)-1] + ") "

	return result + reqString + qBody, nil
}

func (g *Generator) mapQuery(types string, level int) string {
	if _, ok := entity.ScalarTypes[types]; ok {
		return ""
	}

	var (
		result      = "{" + g.prefix
		definition  = g.definitions[types]
		totalIndent = ""
	)

	for idx := 0; idx < level; idx++ {
		totalIndent += g.indent
	}

	for _, prop := range definition.Properties {
		result += g.indent + totalIndent + prop.Name + " " + g.mapQuery(prop.Type, level+1) + g.prefix
	}

	return result + totalIndent + " }"
}
