package p2100

import (
	"fmt"
	"io"
)

func Solve2166(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	type point struct {
		x, y int
	}

	points := make([]point, n+1)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &points[i].x, &points[i].y)
	}

	points[n] = points[0]

	// NOTE - https://en.wikipedia.org/wiki/Shoelace_formula
	area := 0
	for i := 0; i < n; i++ {
		area += points[i].x * points[i+1].y
		area -= points[i].y * points[i+1].x
	}
	if area < 0 {
		area = -area
	}

	_, _ = fmt.Fprintf(writer, "%.1f", float64(area)/2)
}
