package p17400

import (
	"fmt"
	"io"
	"sort"

	"1d1go/utils/ds"
)

func Solve17472(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	board := make([][]int, n)
	islands := make([][]int, n)
	for y := range board {
		board[y] = make([]int, m)
		islands[y] = make([]int, m)
		for x := range board[y] {
			_, _ = fmt.Fscan(reader, &board[y][x])
		}
	}

	cntOfIslands := countAndMarkNumToIslands(n, m, board, islands)
	bridges := makeBridge(n, m, islands)

	res := kruskal(bridges, cntOfIslands)
	_, _ = fmt.Fprint(writer, res)
}

func kruskal(bridges []bridge, cntOfIslands int) int {
	sort.Slice(bridges, func(i, j int) bool {
		return bridges[i].cost < bridges[j].cost
	})

	uf := ds.NewUnionFind(cntOfIslands)
	cnt := 0
	res := 0
	for _, b := range bridges {
		if uf.Find(b.from) == uf.Find(b.to) {
			continue
		}
		uf.Union(b.from, b.to)
		res += b.cost
		cnt++
	}

	if cnt != cntOfIslands-1 {
		return -1
	}
	return res
}

func countAndMarkNumToIslands(n, m int, board, islands [][]int) int {
	idx := 0
	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if board[y][x] == 1 && islands[y][x] == 0 {
				idx++
				dfs(board, islands, y, x, idx)
			}
		}
	}
	return idx
}

func dfs(board [][]int, islands [][]int, y int, x int, idx int) {
	if y < 0 || y >= len(board) || x < 0 || x >= len(board[0]) {
		return
	}

	if board[y][x] == 0 || islands[y][x] != 0 {
		return
	}

	islands[y][x] = idx

	dfs(board, islands, y-1, x, idx)
	dfs(board, islands, y+1, x, idx)
	dfs(board, islands, y, x-1, idx)
	dfs(board, islands, y, x+1, idx)
}

type bridge struct {
	from, to int
	cost     int
}

func makeBridge(n, m int, islands [][]int) []bridge {
	direct := [2]int{-1, 1}

	bridges := make([]bridge, 0, n*m)
	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if islands[y][x] == 0 {
				continue
			}

			for _, d := range direct {
				{
					ny := y + d
					for ny >= 0 && ny < n {
						// 바다일 경우 다리 확장
						if islands[ny][x] == 0 {
							ny += d
							continue
						}

						// 자기 자신에 닿음
						if islands[ny][x] == islands[y][x] {
							break
						}

						// 다른 섬에 닿음
						dist := ny - y
						if dist < 0 {
							dist = -dist
						}
						if dist > 2 { // [0, 1, 2, 3] -> ( 3 - 0 ) - 1 = 2
							bridges = append(bridges, bridge{
								from: islands[y][x],
								to:   islands[ny][x],
								cost: dist - 1, // 중요
							})
						}
						break
					}
				}

				{
					nx := x + d
					for nx >= 0 && nx < m {
						// 바다일 경우 다리 확장
						if islands[y][nx] == 0 {
							nx += d
							continue
						}

						// 자기 자신에 닿음
						if islands[y][nx] == islands[y][x] {
							break
						}

						// 다른 섬에 닿음
						dist := nx - x
						if dist < 0 {
							dist = -dist
						}
						if dist > 2 { // [0, 1, 2, 3] -> ( 3 - 0 ) - 1 = 2
							bridges = append(bridges, bridge{
								from: islands[y][x],
								to:   islands[y][nx],
								cost: dist - 1, // 중요
							})
						}
						break
					}
				}
			}
		}
	}
	return bridges
}
