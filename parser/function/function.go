package function

import (
	"regexp"
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
	"github.com/alvinmatias69/gql-doc/parser/parameter"
)

var (
	funcNameRgx = regexp.MustCompile(`^[a-zA-Z]+`)
	funcRgx     = regexp.MustCompile(`[a-zA-Z]+( )*(\([a-zA-Z:\[\]\! \,]+\))*( )*:( )*[a-zA-Z]+\!`)
	paramRgx    = regexp.MustCompile(`\([a-zA-Z:, !\[\]]+\)`)
	returnRgx   = regexp.MustCompile(`[a-zA-Z]+(!)*$`)
)

func Match(strFn string) bool {
	return funcRgx.MatchString(strFn)
}

func Extract(strFn, comment string) entity.Function {
	name := funcNameRgx.FindString(strFn)
	params := parameter.Extract(paramRgx.FindString(strFn))
	returnType := returnRgx.FindString(strFn)

	function := entity.Function{
		Name:       name,
		Parameters: params,
		ReturnType: strings.Trim(returnType, "!"),
		Comment:    comment,
	}

	return function
}
