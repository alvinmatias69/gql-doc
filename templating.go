package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path"
	"runtime"
	"text/template"
)

const JSONType = "json"
const MarkdownType = "markdown"
const MarkdownPath = "/template/md.tmpl"
const HTMLType = "html"
const HTMLPath = "/template/html.tmpl"

func execToTemplate(data GQLDoc, tmplType string) (string, error) {
	_, b, _, _ := runtime.Caller(0)
	basePath := path.Join(path.Dir(b))
	tmplPath := tmplType

	switch tmplType {
	case JSONType:
		byteData, err := json.Marshal(data)
		return string(byteData), err

	case MarkdownType:
		tmplPath = path.Join(basePath, MarkdownPath)

	case HTMLType:
		tmplPath = path.Join(basePath, HTMLPath)

	default:
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		tmplPath = path.Join(dir, "/", tmplType)
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
