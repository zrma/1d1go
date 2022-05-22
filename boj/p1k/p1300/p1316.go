package p1300

import (
	"fmt"
	"strconv"
)

func Solve1316(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	cnt := 0
	for i := 0; i < n; i++ {
		if scanner.Scan() && isGroupWord(scanner.Text()) {
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
