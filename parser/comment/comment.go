package comment

import (
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`(#)+(.)*`)

func Extract(strComment string) string {
	comment := strings.Trim(strComment, "#")
	comment = strings.TrimSpace(comment)
	comment = comment + "\n"
	return comment
}

func Match(data string) bool {
	return regex.MatchString(data)
}
