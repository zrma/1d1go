package p6900

import (
	"fmt"
	"io"

	"1d1go/boj/p11k/p11300"
)

func Solve6996(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		var s1, s2 string
		_, _ = fmt.Fscanln(reader, &s1, &s2)

		if p11300.IsAnagram(s1, s2) {
			_, _ = fmt.Fprintln(writer, s1, "&", s2, "are anagrams.")
		} else {
			_, _ = fmt.Fprintln(writer, s1, "&", s2, "are NOT anagrams.")
		}
	}
}
