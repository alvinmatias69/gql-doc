package extractor

import (
	"errors"
	"strings"

	"github.com/alvinmatias69/gql-doc/entity"
)

// TODO: Parse Definition with implement
func (e *Extractor) Definition(input string) (entity.Definition, error) {
	idx := strings.Index(input, " ")
	if idx == -1 {
		return entity.Definition{}, errors.New("Invalid definition")
	}

	return entity.Definition{
		Name:    strings.TrimSpace(input[idx : len(input)-1]),
		Variant: entity.Types[input[0:idx]],
	}, nil
}
