package tutorial30daysofcode

import (
	"fmt"
)

var fmtPrintln = fmt.Println

func helloWorld(s string) {
	_, _ = fmtPrintln("Hello, World.")
	_, _ = fmtPrintln(s)
}
