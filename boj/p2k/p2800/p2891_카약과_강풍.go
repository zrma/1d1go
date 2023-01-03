package p2800

import (
	"fmt"
)

func Solve2891(reader Reader, writer Writer) {
	var n, s, r int
	_, _ = fmt.Fscan(reader, &n, &s, &r)

	lost := make([]bool, n)
	for i := 0; i < s; i++ {
		var l int
		_, _ = fmt.Fscan(reader, &l)
		lost[l-1] = true
	}

	reserve := make([]bool, n)
	for i := 0; i < r; i++ {
		var l int
		_, _ = fmt.Fscan(reader, &l)
		reserve[l-1] = true
	}

	for i := 0; i < n; i++ {
		if lost[i] && reserve[i] {
			lost[i] = false
			reserve[i] = false
		}
	}

	for i := 0; i < n; i++ {
		if lost[i] {
			if i > 0 && reserve[i-1] {
				reserve[i-1] = false
				lost[i] = false
			} else if i < n-1 && reserve[i+1] {
				reserve[i+1] = false
				lost[i] = false
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if lost[i] {
			res++
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
