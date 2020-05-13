package variable

import (
	"regexp"
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
	cmt "github.com/alvinmatias69/gql-doc/parser/comment"
	"github.com/alvinmatias69/gql-doc/parser/parameter"
)

type Variable struct {
	Name       string
	Comment    string
	Parameters []entity.Parameter
}

var (
	regex   = regexp.MustCompile(`(type|input|enum) [a-zA-Z]+( )*{`)
	nameRgx = regexp.MustCompile(`[a-zA-Z]+( )*{`)
	enumRgx = regexp.MustCompile(`^([a-zA-Z]+(_)*)+$`)
)

func Match(data string) bool {
	return regex.MatchString(data)
}

func MatchEnum(data string) bool {
	return enumRgx.MatchString(data)
}

func New(line, comment string) *Variable {
	name := nameRgx.FindString(line)
	name = strings.Trim(name, "{")
	name = strings.TrimSpace(name)

	return &Variable{
		Name:    name,
		Comment: strings.TrimSpace(comment),
	}
}

func (v *Variable) AddParameter(line, comment string) {
	line, inlineCmt := getInlineComment(line)
	params := parameter.Parse(line)

	params.Comment = strings.TrimSpace(comment + inlineCmt)
	v.Parameters = append(v.Parameters, params)
}

func (v *Variable) AddEnum(line, comment string) {
	parameter := entity.Parameter{
		Name:    line,
		Comment: strings.TrimSpace(comment),
	}

	v.Parameters = append(v.Parameters, parameter)
}

func (v *Variable) ToEntity() entity.VarType {
	return entity.VarType{
		Name:       v.Name,
		Comment:    v.Comment,
		Parameters: v.Parameters,
	}
}

func getInlineComment(line string) (trimLine, comment string) {
	trimLine = line
	if idx := strings.Index(line, "#"); idx > -1 {
		comment = cmt.Extract(line[idx:])
		trimLine = line[:idx]
	}

	return
}
