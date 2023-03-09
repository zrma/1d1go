package p1900

import (
	"fmt"
	"io"
)

func Solve1919(reader io.Reader, writer io.Writer) {
	var s1, s2 string
	_, _ = fmt.Fscanln(reader, &s1)
	_, _ = fmt.Fscanln(reader, &s2)

	cnt := make([]int, 26)
	for _, c := range s1 {
		cnt[c-'a']++
	}
	for _, c := range s2 {
		cnt[c-'a']--
	}

	var res int
	for _, c := range cnt {
		if c > 0 {
			res += c
		} else {
			res -= c
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
