package template

import (
	"bytes"
	"encoding/json"
	"path"
	"strings"
	"text/template"

	"github.com/alvinmatias69/gql-doc/entity"
	"github.com/rs/zerolog/log"
)

var tmplPathMap = map[entity.Template]string{
	entity.Confluence: "confluence.tmpl",
	entity.Markdown:   "markdown.tmpl",
	entity.HTML:       "html.tmpl",
}

var basepath = "/usr/local/lib/gql-doc/template/"

func (t *Template) Generate() (string, error) {
	if t.template == entity.JSON {
		resByte, err := json.Marshal(t.data)
		return string(resByte), err
	}

	var templatePath = t.templatePath
	if !t.isCustom {
		templatePath = path.Join(basepath, tmplPathMap[t.template])
	}

	log.Info().Msg("Parsing template")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	if t.template == entity.Confluence {
		log.Info().Msg("Escaping confluence special character")
		t.escape()
	}

	var result bytes.Buffer
	log.Info().Msg("Generating template")
	err = tmpl.Execute(&result, t.data)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

func (t *Template) escape() {
	for idx := range t.data.Queries {
		request := t.data.Queries[idx].Example.Request
		request = strings.ReplaceAll(request, "[", `\[`)
		request = strings.ReplaceAll(request, "{", `\{`)
		t.data.Queries[idx].Example.Request = request

		response := t.data.Queries[idx].Example.Response
		response = strings.ReplaceAll(response, "[", `\[`)
		response = strings.ReplaceAll(response, "{", `\{`)
		t.data.Queries[idx].Example.Response = response
	}

	for idx := range t.data.Mutations {
		request := t.data.Mutations[idx].Example.Request
		request = strings.ReplaceAll(request, "[", `\[`)
		request = strings.ReplaceAll(request, "{", `\{`)
		t.data.Mutations[idx].Example.Request = request

		response := t.data.Mutations[idx].Example.Response
		response = strings.ReplaceAll(response, "[", `\[`)
		response = strings.ReplaceAll(response, "{", `\{`)
		t.data.Mutations[idx].Example.Response = response
	}
}
