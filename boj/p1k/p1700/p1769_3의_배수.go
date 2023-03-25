package p1700

import (
	"fmt"
	"io"
)

func Solve1769(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	cnt := 0
	var sum int
	for n := len(s); n > 1; n = len(s) {
		sum = 0
		for i := 0; i < n; i++ {
			sum += int(s[i] - '0')
		}
		s = fmt.Sprint(sum)
		cnt++
	}

	if cnt == 0 {
		sum = int(s[0] - '0')
	}

	_, _ = fmt.Fprintln(writer, cnt)
	if sum%3 == 0 {
		_, _ = fmt.Fprint(writer, "YES")
	} else {
		_, _ = fmt.Fprint(writer, "NO")
	}
}
