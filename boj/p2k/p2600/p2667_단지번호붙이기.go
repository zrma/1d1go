package p2600

import (
	"fmt"
	"sort"
)

func Solve2667(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	table := make([][]bool, n)
	for y := 0; y < n; y++ {
		table[y] = make([]bool, n)
		var s string
		_, _ = fmt.Fscan(reader, &s)
		for x, c := range s {
			if c == '0' {
				table[y][x] = true
			}
		}
	}

	type pos struct {
		x, y int
	}
	var res []int
	queue := make([]pos, 0, n*n)

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if table[y][x] {
				continue
			}
			var cnt int
			queue = append(queue, pos{x, y})
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				if table[p.y][p.x] {
					continue
				}
				cnt++
				table[p.y][p.x] = true
				if p.x > 0 {
					queue = append(queue, pos{p.x - 1, p.y})
				}
				if p.x < n-1 {
					queue = append(queue, pos{p.x + 1, p.y})
				}
				if p.y > 0 {
					queue = append(queue, pos{p.x, p.y - 1})
				}
				if p.y < n-1 {
					queue = append(queue, pos{p.x, p.y + 1})
				}
			}
			res = append(res, cnt)
		}
	}
	sort.Ints(res)
	_, _ = fmt.Fprintln(writer, len(res))
	for _, v := range res {
		_, _ = fmt.Fprintln(writer, v)
	}
}
