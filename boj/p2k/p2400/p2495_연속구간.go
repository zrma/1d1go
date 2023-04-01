package p2400

import (
	"fmt"
	"io"
)

func Solve2495(reader io.Reader, writer io.Writer) {
	for i := 0; i < 3; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)

		max := 1
		cnt := 1
		for j := 1; j < len(s); j++ {
			if s[j-1] == s[j] {
				cnt++
			} else {
				if max < cnt {
					max = cnt
				}
				cnt = 1
			}
		}

		if max < cnt {
			max = cnt
		}
		_, _ = fmt.Fprintln(writer, max)
	}
}
