package p2200

import (
	"fmt"
)

func Solve2206(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	board := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		board[i] = make([]int, m+1)

		var s string
		_, _ = fmt.Fscan(reader, &s)

		for j, c := range s {
			if c == '1' {
				board[i][j+1] = 1
			}
		}
	}

	type pos struct {
		x, y   int
		broken int
		step   int
	}
	visited := make([][][]bool, n+1)
	for i := 0; i <= n; i++ {
		visited[i] = make([][]bool, m+1)
		for j := 0; j <= m; j++ {
			visited[i][j] = make([]bool, 2)
		}
	}
	queue := make([]pos, 0, n*m)

	queue = append(queue, pos{1, 1, 0, 0})
	visited[1][1][0] = true

	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.x == m && p.y == n {
			_, _ = fmt.Fprint(writer, p.step+1)
			return
		}

		for i := 0; i < 4; i++ {
			newX := p.x + dx[i]
			newY := p.y + dy[i]
			alreadyBroken := p.broken

			if newX < 1 || newX > m || newY < 1 || newY > n {
				continue
			}
			if visited[newY][newX][alreadyBroken] {
				continue
			}
			if board[newY][newX] == 1 {
				if alreadyBroken == 1 {
					continue
				}
				alreadyBroken = 1
			}
			queue = append(queue, pos{newX, newY, alreadyBroken, p.step + 1})
			visited[newY][newX][alreadyBroken] = true
		}
	}
	_, _ = fmt.Fprint(writer, -1)
}
