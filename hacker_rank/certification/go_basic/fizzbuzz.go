package go_basic

import (
	"fmt"
)

var fmtPrintln = fmt.Println

func fizzBuzz(n int32) {
	var i int32
	for i = 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			_, _ = fmtPrintln("FizzBuzz")
		case i%3 == 0:
			_, _ = fmtPrintln("Fizz")
		case i%5 == 0:
			_, _ = fmtPrintln("Buzz")
		default:
			_, _ = fmtPrintln(i)
		}
	}
}
