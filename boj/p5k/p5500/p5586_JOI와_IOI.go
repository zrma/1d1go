package p5500

import (
	"fmt"
	"io"
)

func Solve5586(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	joi := 0
	ioi := 0
	for i := 0; i < len(s)-2; i++ {
		if s[i:i+3] == "JOI" {
			joi++
		} else if s[i:i+3] == "IOI" {
			ioi++
		}
	}

	_, _ = fmt.Fprintln(writer, joi)
	_, _ = fmt.Fprint(writer, ioi)
}
