package p12800

import (
	"fmt"
	"io"

	"1d1go/boj/p1k/p1400"
)

func Solve12852(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	dp := p1400.MoveToOne(n)

	_, _ = fmt.Fprintln(writer, dp[n])

	for i := n; i >= 1; {
		_, _ = fmt.Fprint(writer, i, " ")
		if i%2 == 0 && dp[i] == dp[i/2]+1 {
			i /= 2
		} else if i%3 == 0 && dp[i] == dp[i/3]+1 {
			i /= 3
		} else {
			i--
		}
	}
}
