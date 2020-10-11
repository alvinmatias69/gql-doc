package extractor

import "regexp"

type Extractor struct {
	propertyParams *regexp.Regexp
}

func New() *Extractor {
	return &Extractor{
		propertyParams: regexp.MustCompile(`\(.*\)`),
	}
}
