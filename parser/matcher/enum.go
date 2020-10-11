package matcher

func (m *Matcher) Enum(input string) bool {
	return m.enum.MatchString(input)
}
