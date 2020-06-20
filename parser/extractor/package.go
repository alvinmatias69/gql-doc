package extractor

import (
	"strings"
)

func (e *Extractor) Package(input string) string {
	return input[strings.Index(input, " ")+1:]
}
