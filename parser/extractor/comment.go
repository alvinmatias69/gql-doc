package extractor

import "strings"

func (e *Extractor) Comment(input string) string {
	start := strings.Index(input, "#")
	comment := strings.TrimSpace(input[start+1:])
	return comment
}
