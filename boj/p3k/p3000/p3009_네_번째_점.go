package p3000

import (
	"fmt"
	"io"
)

func Solve3009(reader io.Reader, writer io.Writer) {
	resX := 0
	resY := 0
	for i := 0; i < 3; i++ {
		var x, y int
		_, _ = fmt.Fscan(reader, &x, &y)
		resX ^= x
		resY ^= y
	}
	_, _ = fmt.Fprintf(writer, "%d %d", resX, resY)
}
