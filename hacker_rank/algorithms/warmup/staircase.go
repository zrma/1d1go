package warmup

import (
	"fmt"
)

var fmtPrintln = fmt.Println

func staircase(n int32) {
	var i, j int32
	for ; i < n; i++ {
		var s string
		for j = n - 1; j >= 0; j-- {
			if i < j {
				s += " "
			} else {
				s += "#"
			}
		}
		_, _ = fmtPrintln(s)
	}
}
