package main

import (
	"fmt"

	"1d1go/boj/p2k/p2500"
)

func main() {
	var m, n, k int
	_, _ = fmt.Scanf("%d %d %d", &m, &n, &k)

	var rects []p2500.Rect
	for i := 0; i < k; i++ {
		var x1, y1, x2, y2 int
		_, _ = fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
		rects = append(rects, p2500.Rect{x1, y1, x2, y2})
	}
	res := p2500.Solve2583(m, n, rects)
	fmt.Println(len(res))
	for _, n := range res {
		fmt.Printf("%d ", n)
	}
}
