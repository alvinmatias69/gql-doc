package matcher

func (m *Matcher) Comment(input string) bool {
	return len(input) > 0 && input[:1] == "#"
}
