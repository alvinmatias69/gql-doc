package matcher

func (m *Matcher) Definition(input string) bool {
	return m.definition.MatchString(input)
}
