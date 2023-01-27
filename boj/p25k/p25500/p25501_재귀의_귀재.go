package p25500

import (
	"fmt"
	"io"
)

func Solve25501(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		l := 0
		isPalindrome := true
		for r := len(s) - 1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				isPalindrome = false
				break
			}
		}

		l += 1
		if isPalindrome {
			_, _ = fmt.Fprintf(writer, "%d %d\n", 1, l)
		} else {
			_, _ = fmt.Fprintf(writer, "%d %d\n", 0, l)
		}
	}
}
