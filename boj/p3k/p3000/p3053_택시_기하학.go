package p3000

import (
	"fmt"
	"io"
	"math"
)

func Solve3053(reader io.Reader, writer io.Writer) {
	var r int
	_, _ = fmt.Fscan(reader, &r)

	areaOfCircleEuclid := float64(r) * float64(r) * math.Pi
	areaOfCircleTaxi := float64(r) * float64(r) * 2.0

	_, _ = fmt.Fprintf(writer, "%.6f\n", areaOfCircleEuclid)
	_, _ = fmt.Fprintf(writer, "%.6f\n", areaOfCircleTaxi)
}
