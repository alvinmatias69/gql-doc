package matcher

func (m *Matcher) Package(input string) bool {
	return m.packageName.MatchString(input)
}
