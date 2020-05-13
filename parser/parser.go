package parser

import (
	"bufio"
	"os"
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
	cmt "github.com/alvinmatias69/gql-doc/parser/comment"
	"github.com/alvinmatias69/gql-doc/parser/function"
	"github.com/alvinmatias69/gql-doc/parser/variable"
)

func Parse(path string) (entity.Doc, error) {
	file, err := os.Open(path)
	if err != nil {
		return entity.Doc{}, err
	}
	defer file.Close()

	var (
		functions []entity.Function
		variables []entity.VarType
		name      string
		comment   string
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		current = strings.TrimSpace(current)

		switch {
		case len(current) == 0:
			continue

		case cmt.Match(current):
			comment += cmt.Extract(current)

		case variable.Match(current):
			variables = append(variables, iterVariable(current, comment, scanner))
			comment = ""

		case function.Match(current):
			functions = append(functions, function.Extract(current, strings.TrimSpace(comment)))
			comment = ""

		case packageRgx.MatchString(current):
			name = strings.Split(current, " ")[1]
		}
	}

	doc := entity.Doc{
		Name:      name,
		Functions: functions,
		Types:     variables,
	}

	return doc, nil
}

func iterVariable(line, comment string, scanner *bufio.Scanner) entity.VarType {
	var (
		result = variable.New(line, comment)
		state  = entity.VariableProcess
	)
	comment = ""

	for scanner.Scan() {
		current := strings.TrimSpace(scanner.Text())

		switch {
		case len(current) == 0:
			continue

		case cmt.Match(current):
			comment += cmt.Extract(current)

		case strings.Contains(current, "}"):
			state = entity.VariableFinish

		case variable.MatchEnum(current):
			result.AddEnum(current, comment)
			comment = ""

		default:
			result.AddParameter(current, comment)
			comment = ""
		}

		if state == entity.VariableFinish {
			break
		}
	}

	return result.ToEntity()
}
