package union

import (
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
)

func Match(line string) bool {
	return len(line) > 5 && line[:5] == "union"
}

func Extract(line, comment string) entity.VarType {
	data := strings.Split(line[6:], "=")
	varName := strings.TrimSpace(data[0])
	params := strings.Split(data[1], "|")
	parameters := make([]entity.Parameter, 0, len(params))

	for _, param := range params {
		parameters = append(parameters, entity.Parameter{Name: strings.TrimSpace(param)})
	}

	return entity.VarType{
		Name:       varName,
		Comment:    comment,
		Parameters: parameters,
	}
}
