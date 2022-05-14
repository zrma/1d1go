package p2500

import (
	"sort"
)

type Rect [4]int

func Solve2583(fmt InOut) {
	var row, col, k int
	_, _ = fmt.Scan(&row, &col, &k)

	var rects []Rect
	for i := 0; i < k; i++ {
		var x1, y1, x2, y2 int
		_, _ = fmt.Scan(&x1, &y1, &x2, &y2)
		rects = append(rects, Rect{x1, y1, x2, y2})
	}

	table := make([][]bool, row)
	for i := range table {
		table[i] = make([]bool, col)
	}
	for _, rect := range rects {
		for y := rect[1]; y < rect[3]; y++ {
			for x := rect[0]; x < rect[2]; x++ {
				table[y][x] = true
			}
		}
	}

	type pos struct {
		x, y int
	}
	var res []int
	var queue []pos
	for y := 0; y < row; y++ {
		for x := 0; x < col; x++ {
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
				if p.x < col-1 {
					queue = append(queue, pos{p.x + 1, p.y})
				}
				if p.y > 0 {
					queue = append(queue, pos{p.x, p.y - 1})
				}
				if p.y < row-1 {
					queue = append(queue, pos{p.x, p.y + 1})
				}
			}
			res = append(res, cnt)
		}
	}
	sort.Ints(res)

	_, _ = fmt.Println(len(res))
	for _, n := range res {
		_, _ = fmt.Println(n)
	}
}
