package p20100

import (
	"fmt"
	"math"

	"1d1go/boj/p17k/p17300"
)

func Solve20149(reader Reader, writer Writer) {
	const pointCount = 4
	var points [pointCount]p17300.Point
	for i := range points {
		_, _ = fmt.Fscan(reader, &points[i].X, &points[i].Y)
	}

	equalLine, ok := p17300.IsIntersect(points[0], points[1], points[2], points[3])
	if ok {
		_, _ = fmt.Fprintln(writer, 1)

		intersectPos(points[0], points[1], points[2], points[3], equalLine, writer)
	} else {
		_, _ = fmt.Fprintln(writer, 0)
	}
}

func intersectPos(p0, p1, p2, p3 p17300.Point, equalLine bool, writer Writer) {
	x1, y1, x2, y2 := float64(p0.X), float64(p0.Y), float64(p1.X), float64(p1.Y)
	x3, y3, x4, y4 := float64(p2.X), float64(p2.Y), float64(p3.X), float64(p3.Y)

	if equalLine {
		const epsilon = 1e-9

		if math.Abs(p0.Distance(p1)+p2.Distance(p3)-p17300.MaxDistance(p0, p1, p2, p3)) <= epsilon {
			if y1 == y3 || y1 == y4 {
				_, _ = fmt.Fprint(writer, x1, y1)
			} else if y2 == y3 || y2 == y4 {
				_, _ = fmt.Fprint(writer, x2, y2)
			}
		}
		return
	}

	// NOTE - https://www.team504.com/posts/2022/11/2022-11-25-acmicpc-20149-two-lines-intersection-point/
	//        두 선분의 교차점 좌표 계산식 유도 참고
	x := ((x1*y2-x2*y1)*(x3-x4) - (x3*y4-x4*y3)*(x1-x2)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
	y := ((x1*y2-x2*y1)*(y3-y4) - (x3*y4-x4*y3)*(y1-y2)) / ((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
	_, _ = fmt.Fprint(writer, x, y)
}
