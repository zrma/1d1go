package p2500

import (
	"fmt"
)

func Solve2530(reader Reader, writer Writer) {
	var h, m, s, t int
	_, _ = fmt.Fscan(reader, &h, &m, &s, &t)

	s += t
	m += s / 60
	s %= 60
	h += m / 60
	m %= 60
	h %= 24

	_, _ = fmt.Fprint(writer, h, m, s)
}
