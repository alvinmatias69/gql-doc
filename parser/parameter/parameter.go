package parameter

import (
	"errors"
	"regexp"
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
)

var (
	typeRgx     = regexp.MustCompile(`[a-zA-Z]+`)
	builtInType = map[string]interface{}{
		"Int":     nil,
		"String":  nil,
		"Boolean": nil,
		"Float":   nil,
	}
)

func Extract(strParams string) []entity.Parameter {
	if len(strParams) == 0 {
		return nil
	}

	sanitized := strParams[1 : len(strParams)-1]
	params := strings.Split(sanitized, ",")

	result := make([]entity.Parameter, 0, len(params))
	for _, param := range params {
		result = append(result, Parse(param))
	}

	return result
}

func Parse(strParams string) entity.Parameter {
	split := strings.Split(strParams, ":")
	name := strings.TrimSpace(split[0])
	fields := strings.TrimSpace(split[1])

	paramType := typeRgx.FindString(fields)
	_, isBuiltInType := builtInType[paramType]
	isMandatory := strings.Contains(fields, "!")
	isList := strings.Contains(fields, "[")
	isList = isList && strings.Contains(fields, "]")

	return entity.Parameter{
		Name:          name,
		ParamType:     paramType,
		IsBuiltInType: isBuiltInType,
		IsMandatory:   isMandatory,
		IsList:        isList,
	}
}

func ParseWithError(strParams string) (entity.Parameter, error) {
	if len(strParams) < 2 {
		return entity.Parameter{}, errors.New("invalid parameter")
	}

	return Parse(strParams), nil
}
