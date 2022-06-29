package p1300

import (
	"fmt"
)

func Solve1316(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	cnt := 0
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		if isGroupWord(s) {
			cnt++
		}
	}
	_, _ = fmt.Fprint(writer, cnt)
}

func isGroupWord(s string) bool {
	exist := [26]bool{}
	prev := s[0]
	exist[prev-'a'] = true
	for i := 1; i < len(s); i++ {
		c := s[i]
		if prev == c {
			continue
		}
		idx := c - 'a'
		if exist[idx] {
			return false
		}
		exist[idx] = true
		prev = c
	}
	return true
}
