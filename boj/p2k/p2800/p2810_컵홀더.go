package p2800

import (
	"fmt"
	"io"
)

func Solve2810(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	var s string
	_, _ = fmt.Fscanln(reader, &s)

	cnt := 1
	for i := 0; i < n; i++ {
		if s[i] == 'S' {
			cnt++
		} else {
			cnt++
			i++
		}
	}

	if cnt > n {
		cnt = n
	}

	_, _ = fmt.Fprint(writer, cnt)
}
