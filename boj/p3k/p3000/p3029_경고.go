package p3000

import (
	"fmt"
	"io"
)

func Solve3029(reader io.Reader, writer io.Writer) {
	var h1, m1, s1, h2, m2, s2 int
	_, _ = fmt.Fscanf(reader, "%d:%d:%d\n", &h1, &m1, &s1)
	_, _ = fmt.Fscanf(reader, "%d:%d:%d\n", &h2, &m2, &s2)

	t1 := h1*3600 + m1*60 + s1
	t2 := h2*3600 + m2*60 + s2

	if t1 == t2 {
		_, _ = fmt.Fprint(writer, "24:00:00")
		return
	}

	if t1 > t2 {
		t2 += 24 * 3600
	}

	t := t2 - t1

	h := t / 3600
	m := (t % 3600) / 60
	s := t % 60

	_, _ = fmt.Fprintf(writer, "%02d:%02d:%02d", h, m, s)
}
