package matcher

import "regexp"

type Matcher struct {
	property    *regexp.Regexp
	packageName *regexp.Regexp
	union       *regexp.Regexp
	definition  *regexp.Regexp
	enum        *regexp.Regexp
}

func New() *Matcher {
	return &Matcher{
		// regexr.com/55s05
		property: regexp.MustCompile(`\w+\s*(\((\w+\s*:\s*\w+\s*((\"\w+\")|\d+)?\s*(@\w+)?,*\s*)+\))?\s*:\s*\[?\w+\]?(\s*@\w+)*`),
		// regexr.com/55s0e
		packageName: regexp.MustCompile(`package\s+\w+`),
		// regexr.com/55s54
		union: regexp.MustCompile(`union\s*\w*\s*(@\w+)?\s*=(\s*\w\s*\|?)*`),
		// regexr.com/577jk
		definition: regexp.MustCompile(`\w+\s*\w+\s*(implements\s*(\w+\s*&?\s*)+\s*)?{`),
		enum:       regexp.MustCompile(`^\w+$`),
	}
}
