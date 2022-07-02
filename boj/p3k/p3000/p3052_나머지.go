package p3000

import (
	"fmt"
)

func Solve3052(reader Reader, write Writer) {
	m := make(map[int]bool)
	for i := 0; i < 10; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		m[n%42] = true
	}
	_, _ = fmt.Fprint(write, len(m))
}
