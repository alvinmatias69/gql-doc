package extractor

import (
	"errors"
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
)

func (e *Extractor) Union(input string) (entity.Definition, error) {
	splitted := strings.Split(input, "=")
	if len(splitted) < 2 {
		return entity.Definition{}, errors.New("invalid union types")
	}

	name := splitted[0]
	name = strings.TrimSpace(name[strings.Index(name, " ")+1:])

	propString := strings.Split(splitted[1], "|")
	properties := make([]entity.Property, 0, len(propString))

	for _, prop := range propString {
		properties = append(properties, entity.Property{
			Name: strings.TrimSpace(prop),
		})
	}

	return entity.Definition{
		Name:       name,
		Variant:    entity.Union,
		Properties: properties,
	}, nil
}
