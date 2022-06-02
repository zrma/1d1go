package p3000

import (
	"fmt"
	"math"
	"strconv"
)

func Solve3053(scanner Scanner, writer Writer) {
	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())

	areaOfCircleEuclid := float64(r) * float64(r) * math.Pi
	areaOfCircleTaxi := float64(r) * float64(r) * 2.0

	_, _ = fmt.Fprintf(writer, "%.6f\n", areaOfCircleEuclid)
	_, _ = fmt.Fprintf(writer, "%.6f\n", areaOfCircleTaxi)
}
