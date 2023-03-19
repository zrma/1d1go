package p18400

import (
	"fmt"
	"io"
)

func Solve18406(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	n := len(s)
	sum := 0
	for i := 0; i < n/2; i++ {
		sum += int(s[i] - '0')
		sum -= int(s[n-1-i] - '0')
	}

	if sum == 0 {
		_, _ = fmt.Fprint(writer, "LUCKY")
	} else {
		_, _ = fmt.Fprint(writer, "READY")
	}
}
