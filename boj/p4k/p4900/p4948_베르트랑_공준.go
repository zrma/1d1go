package p4900

import (
	"fmt"
	"io"

	"1d1go/utils/integer"
)

func Solve4948(reader io.Reader, writer io.Writer) {
	const maxLen = 123456*2 + 1
	sieveOfEratosthenes := integer.BuildSieveOfEratosthenes(maxLen)
	for {
		var n int
		if _, err := fmt.Fscan(reader, &n); err != nil {
			break
		}
		if n == 0 {
			return
		}
		cnt := 0
		for i := n + 1; i <= 2*n; i++ {
			if !sieveOfEratosthenes[i] {
				cnt++
			}
		}
		_, _ = fmt.Fprintln(writer, cnt)
	}
}
