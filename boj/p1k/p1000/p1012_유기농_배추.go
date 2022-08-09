package p1000

import (
	"fmt"
)

func Solve1012(reader Reader, writer Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		countGroups(reader, writer)
	}
}

func countGroups(reader Reader, writer Writer) {
	var m, n, k int
	_, _ = fmt.Fscan(reader, &m, &n, &k)

	field := make([][]bool, n)
	for i := 0; i < n; i++ {
		field[i] = make([]bool, m)
	}

	for i := 1; i <= k; i++ {
		var x, y int
		_, _ = fmt.Fscan(reader, &x, &y)
		field[y][x] = true
	}

	var count int
	var queue []pos
	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if !field[y][x] {
				continue
			}
			queue = append(queue, pos{x, y})
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				if !field[p.y][p.x] {
					continue
				}
				field[p.y][p.x] = false
				if p.x > 0 {
					queue = append(queue, pos{p.x - 1, p.y})
				}
				if p.x < m-1 {
					queue = append(queue, pos{p.x + 1, p.y})
				}
				if p.y > 0 {
					queue = append(queue, pos{p.x, p.y - 1})
				}
				if p.y < n-1 {
					queue = append(queue, pos{p.x, p.y + 1})
				}
			}
			count++
		}
	}
	_, _ = fmt.Fprintln(writer, count)
}

type pos struct {
	x, y int
}
