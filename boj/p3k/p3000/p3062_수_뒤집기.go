package p3000

import (
	"fmt"
	"io"

	"1d1go/utils/str"
)

func Solve3062(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)

		s2 := str.Reverse(s)

		var a, b int
		_, _ = fmt.Sscanf(s, "%d", &a)
		_, _ = fmt.Sscanf(s2, "%d", &b)

		sum := fmt.Sprint(a + b)
		isPalindrome := true
		for j := 0; j < len(sum)/2; j++ {
			if sum[j] != sum[len(sum)-1-j] {
				isPalindrome = false
				break
			}
		}

		if isPalindrome {
			_, _ = fmt.Fprintln(writer, "YES")
		} else {
			_, _ = fmt.Fprintln(writer, "NO")
		}
	}
}
