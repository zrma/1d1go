package p15800

import (
	"fmt"
)

func Solve15829(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	const (
		r = int64(31)
		m = int64(1_234_567_891)
	)

	var base int64 = 1
	var res int64

	var s string
	_, _ = fmt.Fscan(reader, &s)

	for i := 0; i < n; i++ {
		v := int64(s[i] - 'a' + 1)
		res = (res + v*base) % m
		base = (base * r) % m
	}
	_, _ = fmt.Fprint(writer, res)
}
