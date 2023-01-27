package p7500

import (
	"fmt"
	"io"
)

func Solve7569(reader io.Reader, writer io.Writer) {
	var m, n, h int
	_, _ = fmt.Fscan(reader, &m, &n, &h)

	tomatoes := make([][][]int, h)
	for z := 0; z < h; z++ {
		tomatoes[z] = make([][]int, n)
		for y := 0; y < n; y++ {
			tomatoes[z][y] = make([]int, m)
			for x := 0; x < m; x++ {
				_, _ = fmt.Fscan(reader, &tomatoes[z][y][x])
			}
		}
	}

	type pos struct {
		x, y, z int
	}
	const max = 100 + 1
	queue := make([]pos, 0, max)
	for z := 0; z < h; z++ {
		for y := 0; y < n; y++ {
			for x := 0; x < m; x++ {
				if tomatoes[z][y][x] == 1 {
					queue = append(queue, pos{x, y, z})
				}
			}
		}
	}

	dx := []int{-1, 1, 0, 0, 0, 0}
	dy := []int{0, 0, -1, 1, 0, 0}
	dz := []int{0, 0, 0, 0, -1, 1}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for i := 0; i < 6; i++ {
			newX := p.x + dx[i]
			newY := p.y + dy[i]
			newZ := p.z + dz[i]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || newZ < 0 || newZ >= h {
				continue
			}
			if tomatoes[newZ][newY][newX] != 0 {
				continue
			}
			tomatoes[newZ][newY][newX] = tomatoes[p.z][p.y][p.x] + 1
			queue = append(queue, pos{newX, newY, newZ})
		}
	}
	_, _ = fmt.Fprint(writer, maxVal3D(tomatoes))
}

func maxVal3D(tomatoes [][][]int) int {
	max := 0
	for _, heights := range tomatoes {
		for _, row := range heights {
			for _, val := range row {
				if val == 0 {
					return -1
				}
				if val > max {
					max = val
				}
			}
		}
	}
	return max - 1
}
