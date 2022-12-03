package p16200

import (
	"fmt"
	"sort"
)

func Solve16236(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	sharkX := 0
	sharkY := 0
	sharkSize := 2

	board := make([][]int, n)
	for y := 0; y < n; y++ {
		board[y] = make([]int, n)
		for x := 0; x < n; x++ {
			_, _ = fmt.Fscan(reader, &board[y][x])
			if board[y][x] == 9 {
				sharkX = x
				sharkY = y
				board[y][x] = 0
			}
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	type pos struct {
		x, y, dist int
	}

	eatCnt := 0
	res := 0
	for {
		visited := make([][]bool, n)
		for y := 0; y < n; y++ {
			visited[y] = make([]bool, n)
		}
		visited[sharkY][sharkX] = true

		queue := make([]pos, 0, n*n)
		queue = append(queue, pos{sharkX, sharkY, 0})

		fishes := make([]pos, 0, n)

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			for i := 0; i < 4; i++ {
				newX := p.x + dx[i]
				newY := p.y + dy[i]

				if newX < 0 || newX >= n || newY < 0 || newY >= n {
					continue
				}
				if visited[newY][newX] {
					continue
				}
				if board[newY][newX] > sharkSize {
					continue
				}

				visited[newY][newX] = true
				if board[newY][newX] != 0 && board[newY][newX] < sharkSize {
					fishes = append(fishes, pos{newX, newY, p.dist + 1})
				} else {
					queue = append(queue, pos{newX, newY, p.dist + 1})
				}
			}
		}

		if len(fishes) == 0 {
			break
		}

		sort.Slice(fishes, func(i, j int) bool {
			if fishes[i].dist == fishes[j].dist {
				if fishes[i].y == fishes[j].y {
					return fishes[i].x < fishes[j].x
				}
				return fishes[i].y < fishes[j].y
			}
			return fishes[i].dist < fishes[j].dist
		})

		sharkX = fishes[0].x
		sharkY = fishes[0].y
		res += fishes[0].dist
		board[sharkY][sharkX] = 0
		eatCnt++

		if eatCnt == sharkSize {
			sharkSize++
			eatCnt = 0
		}
	}

	_, _ = fmt.Fprint(writer, res)
}
