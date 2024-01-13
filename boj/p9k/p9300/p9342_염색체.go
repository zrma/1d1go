package p9300

import (
	"fmt"
	"io"
	"regexp"
)

func Solve9342(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)

		if isMatched(s) {
			_, _ = fmt.Fprintln(writer, "Infected!")
		} else {
			_, _ = fmt.Fprintln(writer, "Good")
		}
	}
}

var regexpMatchString = regexp.MatchString

func isMatched(s string) bool {
	//goland:noinspection SpellCheckingInspection
	const (
		regexPattern = `^[ABCDEF]?A+F+C+[ABCDEF]?$`
	)

	matched, err := regexpMatchString(regexPattern, s)
	if err != nil {
		return false
	}

	return matched
}
