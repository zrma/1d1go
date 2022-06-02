package p1100

import (
	"fmt"
	"sort"
	"strconv"
)

func Solve1181(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

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
		scanner.Scan()
		if s := scanner.Text(); notIn(s) {
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
