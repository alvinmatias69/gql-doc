package extractor

import (
	"errors"
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
)

func (e *Extractor) Property(input string) (entity.Property, error) {
	index := strings.LastIndex(input, ":")
	if index == -1 {
		return entity.Property{}, errors.New("invalid property types")
	}

	var parameters []entity.Property
	name := input[:index]
	if !strings.Contains(name, "()") && strings.Contains(name, "(") {
		paramString := e.propertyParams.FindString(name)
		paramString = strings.Trim(paramString, `()`)
		params := strings.Split(paramString, ",")
		result := make([]entity.Property, 0, len(params))
		for _, param := range params {
			idx := strings.Index(param, ":")
			if idx == -1 {
				return entity.Property{}, errors.New("invalid property params")
			}
			prop := extractPropTypes(param[idx+1:])
			result = append(result, entity.Property{
				Name:       strings.TrimSpace(param[:idx]),
				Type:       prop.Name,
				IsList:     prop.IsList,
				IsNullable: prop.IsNullable,
				IsScalar:   prop.IsScalar,
			})
		}
		parameters = result
		name = strings.TrimSpace(name[:strings.Index(name, "(")])
	}

	prop := extractPropTypes(input[index+1:])

	return entity.Property{
		Name:       strings.ReplaceAll(strings.TrimSpace(name), "()", ""),
		Type:       prop.Name,
		IsScalar:   prop.IsScalar,
		IsNullable: prop.IsNullable,
		IsList:     prop.IsList,
		Parameters: parameters,
	}, nil
}

func extractPropTypes(input string) entity.Property {
	input = strings.TrimSpace(input)

	isNullable := input[len(input)-1] != '!'
	if !isNullable {
		input = input[:len(input)-1]
	}

	isList := input[0] == '[' && input[len(input)-1] == ']'
	if isList {
		input = input[1 : len(input)-1]
	}

	_, isScalar := entity.ScalarTypes[input]

	return entity.Property{
		Name:       input,
		IsNullable: isNullable,
		IsScalar:   isScalar,
		IsList:     isList,
	}
}
