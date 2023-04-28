package p2100

import (
	"fmt"
	"io"
)

func Solve2153(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	var sum int
	for _, c := range s {
		if c >= 'a' {
			sum += int(c - 'a' + 1)
		} else {
			sum += int(c - 'A' + 27)
		}
	}

	if isPrime(sum) {
		_, _ = fmt.Fprint(writer, "It is a prime word.")
	} else {
		_, _ = fmt.Fprint(writer, "It is not a prime word.")
	}
}

func isPrime(n int) bool {
	if n < 4 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
