package p7800

import (
	"fmt"
	"math"
)

func Solve7869(reader Reader, writer Writer) {
	var x1, y1, r1, x2, y2, r2 float64
	_, _ = fmt.Fscan(reader, &x1, &y1, &r1, &x2, &y2, &r2)

	if r1 < r2 {
		x1, y1, r1, x2, y2, r2 = x2, y2, r2, x1, y1, r1
	}

	var d = distance(x1, y1, x2, y2)
	// 겹치지 않음
	if d >= r1+r2 {
		_, _ = fmt.Fprintf(writer, "%0.3f", 0.0)
		return
	}
	// 포함
	if d+r2 <= r1 {
		_, _ = fmt.Fprintf(writer, "%0.3f", math.Pi*r2*r2)
		return
	}

	var a = math.Acos((r1*r1 + d*d - r2*r2) / (2 * r1 * d))
	var b = math.Acos((r2*r2 + d*d - r1*r1) / (2 * r2 * d))

	var result = r1*r1*a + r2*r2*b - d*r1*math.Sin(a)
	_, _ = fmt.Fprintf(writer, "%0.3f", result)
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
