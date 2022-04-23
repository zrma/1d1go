package main

import (
	"fmt"

	"1d1go/boj"
)

func main() {
	var m, n, k int
	_, _ = fmt.Scanf("%d %d %d", &m, &n, &k)

	var rects []boj.Rect
	for i := 0; i < k; i++ {
		var x1, y1, x2, y2 int
		_, _ = fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
		rects = append(rects, boj.Rect{x1, y1, x2, y2})
	}
	res := boj.Solve2583(m, n, rects)
	fmt.Println(len(res))
	for _, n := range res {
		fmt.Printf("%d ", n)
	}
}
