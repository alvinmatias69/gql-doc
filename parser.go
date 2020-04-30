package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var packageRgx = regexp.MustCompile(`package [a-zA-Z]+`)
var funcRgx = regexp.MustCompile(`[a-zA-Z]+( )*(\([a-zA-Z:\[\]\! \,]+\))*( )*:( )*[a-zA-Z]+\!`)
var funcNameRgx = regexp.MustCompile(`^[a-zA-Z]+`)
var paramRgx = regexp.MustCompile(`\([a-zA-Z:, !\[\]]+\)`)
var returnRgx = regexp.MustCompile(`[a-zA-Z]+(!)*$`)
var typeRgx = regexp.MustCompile(`[a-zA-Z]+`)
var varTypeRgx = regexp.MustCompile(`(type|input) [a-zA-Z]+( )*{`)
var varTypeNameRgx = regexp.MustCompile(`[a-zA-Z]+( )*{`)
var commentRgx = regexp.MustCompile(`(#)+(.)*`)

var builtInType = map[string]interface{}{
	"Int":     nil,
	"String":  nil,
	"Boolean": nil,
	"Float":   nil,
}

func parse(path string) (Docs, error) {
	var (
		functions []Function
		variables []VarType
		name      string
	)

	file, err := os.Open(path)
	if err != nil {
		return Docs{}, err
	}
	defer file.Close()

	var variable *VarType
	var comment string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		current = strings.TrimSpace(current)

		switch {
		case len(current) == 0:
			continue

		case commentRgx.MatchString(current):
			comment += extractComment(commentRgx.FindString(current))

		case strings.Contains(current, "}"):
			variables = append(variables, *variable)
			variable = nil

		case variable != nil:
			if idx := strings.Index(current, "#"); idx > -1 {
				comment += extractComment(current[idx:])
				current = current[:idx]
			}
			params := parseParams(current)
			params.Comment = strings.TrimSpace(comment)
			variable.Parameters = append(variable.Parameters, params)
			comment = ""

		case varTypeRgx.MatchString(current):
			name := varTypeNameRgx.FindString(current)
			name = name[:len(name)-1]
			name = strings.TrimSpace(name)
			variable = &VarType{
				Name:    name,
				Comment: strings.TrimSpace(comment),
			}
			comment = ""

		case funcRgx.MatchString(current):
			function := extractFunction(funcRgx.FindString(current))
			function.Comment = strings.TrimSpace(comment)
			functions = append(functions, function)
			comment = ""

		case packageRgx.MatchString(current):
			name = strings.Split(current, " ")[1]
		}
	}

	doc := Docs{
		Name:      name,
		Functions: functions,
		Types:     variables,
	}

	return doc, nil
}

func extractFunction(strFn string) Function {
	name := funcNameRgx.FindString(strFn)
	params := extractParams(paramRgx.FindString(strFn))
	returnType := returnRgx.FindString(strFn)

	function := Function{
		Name:       name,
		Parameters: params,
		ReturnType: strings.Trim(returnType, "!"),
	}

	return function
}

func extractParams(strParams string) []Parameter {
	if len(strParams) == 0 {
		return nil
	}

	sanitized := strParams[1 : len(strParams)-1]
	params := strings.Split(sanitized, ",")

	result := make([]Parameter, 0, len(params))
	for _, param := range params {
		result = append(result, parseParams(param))
	}

	return result
}

func parseParams(strParams string) Parameter {
	split := strings.Split(strParams, ":")
	name := strings.TrimSpace(split[0])
	fields := strings.TrimSpace(split[1])

	paramType := typeRgx.FindString(fields)
	_, isBuiltInType := builtInType[paramType]
	isMandatory := strings.Contains(fields, "!")
	isList := strings.Contains(fields, "[")
	isList = isList && strings.Contains(fields, "]")

	return Parameter{
		Name:          name,
		ParamType:     paramType,
		IsBuiltInType: isBuiltInType,
		IsMandatory:   isMandatory,
		IsList:        isList,
	}
}

func extractComment(strComment string) string {
	comment := strings.Trim(strComment, "#")
	comment = strings.TrimSpace(comment)
	comment = comment + "\n"
	return comment
}
