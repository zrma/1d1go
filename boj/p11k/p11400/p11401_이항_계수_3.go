package p11400

import (
	"fmt"
	"io"
)

func Solve11401(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	const mod = 1_000_000_007

	factorial := make([]int64, n+1)
	factorial[0] = 1
	factorial[1] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * int64(i) % mod
	}

	res := factorial[n] * pow(factorial[k]*factorial[n-k]%mod, mod-2, mod) % mod
	_, _ = fmt.Fprint(writer, res)
}

func pow(a, b, c int64) int64 {
	if b == 1 {
		return a % c
	}
	if b%2 == 0 {
		return pow(a*a%c, b/2, c)
	}
	return a * pow(a*a%c, b/2, c) % c
}
