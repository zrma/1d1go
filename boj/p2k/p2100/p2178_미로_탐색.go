package p2100

import (
	"fmt"
)

func Solve2178(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	const invalid = -1
	table := make([][]int, n)
	for y := 0; y < n; y++ {
		table[y] = make([]int, m)
		var s string
		_, _ = fmt.Fscan(reader, &s)
		for x, c := range s {
			if c == '0' {
				table[y][x] = invalid
			}
		}
	}

	type pos struct {
		x, y int
	}
	queue := make([]pos, 0, n*m)
	queue = append(queue, pos{0, 0})
	table[0][0] = 1

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.x == m-1 && p.y == n-1 {
			break
		}

		for i := 0; i < 4; i++ {
			newX := p.x + dx[i]
			newY := p.y + dy[i]

			if newX < 0 || newX >= m || newY < 0 || newY >= n {
				continue
			}
			if table[newY][newX] != 0 {
				continue
			}

			table[newY][newX] = table[p.y][p.x] + 1
			queue = append(queue, pos{newX, newY})
		}
	}
	_, _ = fmt.Fprint(writer, table[n-1][m-1])
}
