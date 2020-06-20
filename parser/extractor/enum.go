package extractor

import "github.com/alvinmatias69/gql-doc/entity"

func (e *Extractor) Enum(input string) entity.Property {
	return entity.Property{
		Name: input,
	}
}
