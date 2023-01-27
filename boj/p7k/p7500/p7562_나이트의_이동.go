package p7500

import (
	"fmt"
	"io"
)

func Solve7562(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	type pos struct {
		x, y int
	}

	dx := []int{-2, -2, -1, -1, 1, 1, 2, 2}
	dy := []int{-1, 1, -2, 2, -2, 2, -1, 1}

	const max = 300 + 1
	queue := make([]pos, 0, max*max)
	for i := 0; i < n; i++ {
		var l int
		_, _ = fmt.Fscan(reader, &l)

		var from, to pos
		_, _ = fmt.Fscan(reader, &from.x, &from.y)
		_, _ = fmt.Fscan(reader, &to.x, &to.y)

		queue = queue[:0]
		queue = append(queue, from)

		table := make([]int, max*max)

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if p.x == to.x && p.y == to.y {
				break
			}

			for j := 0; j < 8; j++ {
				newX := p.x + dx[j]
				newY := p.y + dy[j]
				if newX < 0 || newX >= l || newY < 0 || newY >= l {
					continue
				}
				if table[newY*l+newX] != 0 {
					continue
				}
				table[newY*l+newX] = table[p.y*l+p.x] + 1
				queue = append(queue, pos{newX, newY})
			}
		}
		_, _ = fmt.Fprintln(writer, table[to.y*l+to.x])
	}
}
