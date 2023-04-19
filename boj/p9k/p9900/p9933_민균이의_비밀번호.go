package p9900

import (
	"fmt"
	"io"
)

func Solve9933(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	words := make([]string, n)
	for i := range words {
		_, _ = fmt.Fscanln(reader, &words[i])
	}

	for i := range words {
		for j := range words {
			if isReversed(words[i], words[j]) {
				_, _ = fmt.Fprintf(writer, "%d %c", len(words[i]), words[i][len(words[i])/2])
				return
			}
		}
	}
}

func isReversed(s1 string, s2 string) bool {
	for i := range s1 {
		if s1[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}
