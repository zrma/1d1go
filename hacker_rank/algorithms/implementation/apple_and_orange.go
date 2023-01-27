package implementation

import (
	"fmt"
)

var fmtPrintln = fmt.Println

func countApplesAndOranges(s, t, a, b int32, apples, oranges []int32) {
	containFrom := func(origin int32) func(int32) bool {
		return func(distance int32) bool {
			return (s <= origin+distance) && (origin+distance <= t)
		}
	}

	_, _ = fmtPrintln(count(apples, containFrom(a)))
	_, _ = fmtPrintln(count(oranges, containFrom(b)))
}

func count(distances []int32, contain func(int32) bool) (sum int32) {
	for _, distance := range distances {
		if contain(distance) {
			sum++
		}
	}
	return
}
