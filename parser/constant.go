package parser

import "regexp"

var (
	packageRgx = regexp.MustCompile(`package [a-zA-Z]+`)
)
