package matcher

func (m *Matcher) Property(input string) bool {
	return m.property.MatchString(input)
}
