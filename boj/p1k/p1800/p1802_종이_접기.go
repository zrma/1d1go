package p1800

import (
	"fmt"
)

func Solve1802(reader Reader, writer Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		if solve1802(s, 0, len(s)-1) {
			_, _ = fmt.Fprintln(writer, "YES")
		} else {
			_, _ = fmt.Fprintln(writer, "NO")
		}
	}
}

func solve1802(s string, start, end int) bool {
	if start >= end {
		return true
	}

	lo := start
	hi := end
	for lo < hi {
		if s[lo] == s[hi] {
			return false
		}
		lo++
		hi--
	}

	mid := start + (end-start)/2
	return solve1802(s, start, mid-1) && solve1802(s, mid+1, end)
}
