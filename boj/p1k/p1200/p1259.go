package p1200

import (
	"fmt"
)

func Solve1259(scanner Scanner, writer Writer) {
	for scanner.Scan() {
		s := scanner.Text()
		if s == "0" {
			break
		}
		if isPalindrome(s) {
			_, _ = fmt.Fprintln(writer, "yes")
		} else {
			_, _ = fmt.Fprintln(writer, "no")
		}
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
