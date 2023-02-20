package p7500

import (
	"fmt"
	"io"
)

func Solve7567(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	ans := 10
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			ans += 5
		} else {
			ans += 10
		}
	}
	_, _ = fmt.Fprint(writer, ans)
}
