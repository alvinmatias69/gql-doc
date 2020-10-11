package matcher

import (
	"bytes"
	"text/template"
)

func formatNamed(format string, entry map[string]string) (string, error) {
	tpl, err := template.New("format").Parse(format)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, entry)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
