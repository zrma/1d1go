package p3000

import (
	"fmt"
	"math"
)

func Solve3053(reader Reader, writer Writer) {
	var r int
	_, _ = fmt.Fscan(reader, &r)

	areaOfCircleEuclid := float64(r) * float64(r) * math.Pi
	areaOfCircleTaxi := float64(r) * float64(r) * 2.0

	_, _ = fmt.Fprintf(writer, "%.6f\n", areaOfCircleEuclid)
	_, _ = fmt.Fprintf(writer, "%.6f\n", areaOfCircleTaxi)
}
