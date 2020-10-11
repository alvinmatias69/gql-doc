package matcher

func (m *Matcher) Union(input string) bool {
	return m.union.MatchString(input)
}
