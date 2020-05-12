package templating

import (
	"bytes"
	"encoding/json"
	"os"
	"path"
	"runtime"
	"text/template"

	"github.com/alvinmatias69/gql-doc/entity"
)

const markdownPath = "/template/md.tmpl"
const htmlPath = "/template/html.tmpl"

func ToTemplate(data entity.GQLDoc, tmplType entity.Template) (string, error) {
	_, b, _, _ := runtime.Caller(0)
	basePath := path.Join(path.Dir(b))
	tmplPath := string(tmplType)

	switch tmplType {
	case entity.JSON:
		byteData, err := json.Marshal(data)
		return string(byteData), err

	case entity.Markdown:
		tmplPath = path.Join(basePath, markdownPath)

	case entity.HTML:
		tmplPath = path.Join(basePath, htmlPath)

	default:
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		tmplPath = path.Join(dir, "/", tmplPath)
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
