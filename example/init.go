package example

import (
	"regexp"

	"github.com/alvinmatias69/gql-doc/entity"
)

type Generator struct {
	spec        entity.Spec
	definitions map[string]entity.Definition
	prefix      string
	indent      string
	queryRe     *regexp.Regexp
}

func New(spec entity.Spec, template string) *Generator {
	var (
		definitions        = make(map[string]entity.Definition)
		prefix      string = "\n"
		indent      string = "\t"
	)
	for _, definition := range spec.Definitions {
		definitions[definition.Name] = definition
	}

	if tmpl, ok := entity.AvailableTemplate[template]; ok && tmpl == entity.Confluence {
		prefix = "\\\\"
		indent = `⠀⠀`
	}

	return &Generator{
		spec:        spec,
		definitions: definitions,
		prefix:      prefix,
		indent:      indent,
		queryRe:     regexp.MustCompile(`\"(\w+)\":`),
	}
}
