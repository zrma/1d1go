package p11300

import (
	"fmt"
	"io"
)

func Solve11328(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var s1, s2 string
		_, _ = fmt.Fscanln(reader, &s1, &s2)

		if IsAnagram(s1, s2) {
			_, _ = fmt.Fprintln(writer, "Possible")
		} else {
			_, _ = fmt.Fprintln(writer, "Impossible")
		}
	}
}

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var counts [26]int
	for i := 0; i < len(s1); i++ {
		counts[s1[i]-'a']++
		counts[s2[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if counts[i] != 0 {
			return false
		}
	}

	return true
}
