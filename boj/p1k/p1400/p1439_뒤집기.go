package p1400

import (
	"fmt"
	"io"
)

func Solve1439(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	n := len(s)
	cnt0 := 0
	cnt1 := 0

	if n < 2 {
		_, _ = fmt.Fprint(writer, 0)
		return
	}

	if s[0] == '0' {
		cnt0++
	} else {
		cnt1++
	}

	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			if s[i] == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}
	}

	if cnt0 < cnt1 {
		_, _ = fmt.Fprint(writer, cnt0)
	} else {
		_, _ = fmt.Fprint(writer, cnt1)
	}
}
