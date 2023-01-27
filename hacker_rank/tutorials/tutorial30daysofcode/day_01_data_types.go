package tutorial30daysofcode

import (
	"fmt"
)

var fmtPrintf = fmt.Printf

func dataType(i1, i2 uint64, d1, d2 float64, s1, s2 string) {
	_, _ = fmtPrintln(i1 + i2)
	_, _ = fmtPrintf("%.1f\n", d1+d2)
	_, _ = fmtPrintln(s1 + s2)
}
