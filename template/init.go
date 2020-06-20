package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alvinmatias69/gql-doc/entity"
	"github.com/rs/zerolog/log"
)

type Template struct {
	isCustom     bool
	templatePath string
	template     entity.Template
	data         entity.Spec
}

func New(data entity.Spec, template string) (*Template, error) {
	tmplType, ok := entity.AvailableTemplate[template]
	if ok {
		return &Template{
			data:     data,
			template: tmplType,
			isCustom: false,
		}, nil
	}

	log.Info().Msg("Checking template file")
	templatePath := filepath.Clean(os.ExpandEnv(template))
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Unable to locate %s", templatePath)
	}

	return &Template{
		data:         data,
		template:     tmplType,
		isCustom:     true,
		templatePath: templatePath,
	}, nil
}
