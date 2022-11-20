package p17300

import (
	"fmt"
	"math"

	"1d1go/utils/integer"
)

func Solve17386(reader Reader, writer Writer) {
	const pointCount = 4
	var points [pointCount]Point
	for i := range points {
		_, _ = fmt.Fscan(reader, &points[i].X, &points[i].Y)
	}

	_, res := IsIntersect(points[0], points[1], points[2], points[3])
	if res {
		_, _ = fmt.Fprint(writer, 1)
	} else {
		_, _ = fmt.Fprint(writer, 0)

	}
}

type Point struct {
	X, Y int
}

func (p Point) LessThan(other Point) bool {
	if p.X < other.X {
		return true
	} else if p.X == other.X {
		return p.Y < other.Y
	}
	return false
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X-other.X), 2) + math.Pow(float64(p.Y-other.Y), 2))
}

func MaxDistance(p0, p1, p2, p3 Point) float64 {
	dist0 := p0.Distance(p2)
	dist1 := p0.Distance(p3)
	dist2 := p1.Distance(p2)
	dist3 := p1.Distance(p3)

	return math.Max(math.Max(dist0, dist1), math.Max(dist2, dist3))
}

func IsIntersect(p0, p1, p2, p3 Point) (bool, bool) {
	abc := integer.CCW(p0.X, p0.Y, p1.X, p1.Y, p2.X, p2.Y)
	abd := integer.CCW(p0.X, p0.Y, p1.X, p1.Y, p3.X, p3.Y)
	cda := integer.CCW(p2.X, p2.Y, p3.X, p3.Y, p0.X, p0.Y)
	cdb := integer.CCW(p2.X, p2.Y, p3.X, p3.Y, p1.X, p1.Y)

	if equalLine := abc == 0 && abd == 0 && cda == 0 && cdb == 0; equalLine {
		const epsilon = 1e-9
		return true, MaxDistance(p0, p1, p2, p3) <= p0.Distance(p1)+p2.Distance(p3)+epsilon
	}

	v0 := abc * abd
	v1 := cda * cdb
	return false, v0 <= 0 && v1 <= 0
}
