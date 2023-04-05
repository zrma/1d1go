package p1200

import (
	"fmt"
	"io"
)

func Solve1254(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	n := len(s)
	for i := 0; i < n; i++ {
		if isPalindrome(s[i:]) {
			_, _ = fmt.Fprint(writer, n+i)
			return
		}
	}
}
