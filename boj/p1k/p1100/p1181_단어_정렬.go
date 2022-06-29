package p1100

import (
	"fmt"
	"sort"
)

func Solve1181(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	ss := make([]string, 0, n)
	notIn := func(target string) bool {
		for _, s := range ss {
			if s == target {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		if notIn(s) {
			ss = append(ss, s)
		}
	}

	sort.Slice(ss, func(i, j int) bool {
		if len(ss[i]) == len(ss[j]) {
			return ss[i] < ss[j]
		}
		return len(ss[i]) < len(ss[j])
	})

	for _, s := range ss {
		_, _ = fmt.Fprintln(writer, s)
	}
}
