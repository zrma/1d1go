package p7500

import (
	"fmt"
)

func Solve7576(reader Reader, writer Writer) {
	var m, n int
	_, _ = fmt.Fscan(reader, &m, &n)

	tomatoes := make([][]int, n)
	for y := 0; y < n; y++ {
		tomatoes[y] = make([]int, m)
		for x := 0; x < m; x++ {
			_, _ = fmt.Fscan(reader, &tomatoes[y][x])
		}
	}

	type pos struct {
		x, y int
	}
	const max = 1000 + 1
	queue := make([]pos, 0, max)
	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if tomatoes[y][x] == 1 {
				queue = append(queue, pos{x, y})
			}
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			newX := p.x + dx[i]
			newY := p.y + dy[i]
			if newX < 0 || newX >= m || newY < 0 || newY >= n {
				continue
			}
			if tomatoes[newY][newX] != 0 {
				continue
			}
			tomatoes[newY][newX] = tomatoes[p.y][p.x] + 1
			queue = append(queue, pos{newX, newY})
		}
	}
	_, _ = fmt.Fprint(writer, maxVal2D(tomatoes))
}

func maxVal2D(tomatoes [][]int) int {
	max := 1
	for _, row := range tomatoes {
		for _, val := range row {
			if val == 0 {
				return -1
			}
			if val > max {
				max = val
			}
		}
	}
	return max - 1
}
