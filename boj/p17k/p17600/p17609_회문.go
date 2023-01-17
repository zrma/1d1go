package p17600

import (
	"fmt"
)

func Solve17609(reader Reader, writer Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		_, _ = fmt.Fprintln(writer, solve17609(s, 0, len(s)-1, false))
	}
}

func solve17609(s string, begin int, end int, skipped bool) int {
	for begin < end {
		if s[begin] == s[end] {
			begin++
			end--
		} else {
			if skipped {
				return 2
			}

			skipped = true
			res1 := solve17609(s, begin+1, end, skipped)
			res2 := solve17609(s, begin, end-1, skipped)
			if res1 == 1 || res2 == 1 {
				return 1
			} else {
				return 2
			}
		}
	}

	if skipped {
		return 1
	} else {
		return 0
	}
}
