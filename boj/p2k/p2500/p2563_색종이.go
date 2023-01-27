package p2500

import (
	"fmt"
	"io"
)

func Solve2563(reader io.Reader, writer io.Writer) {
	const (
		maxSize = 100
		boxSize = 10
	)
	arr := make([][]bool, maxSize)
	for i := 0; i < maxSize; i++ {
		arr[i] = make([]bool, maxSize)
	}

	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var x, y int
		_, _ = fmt.Fscan(reader, &x, &y)
		for j := x; j < x+boxSize; j++ {
			for k := y; k < y+boxSize; k++ {
				arr[j][k] = true
			}
		}
	}

	var res int
	for i := 0; i < maxSize; i++ {
		for j := 0; j < maxSize; j++ {
			if arr[i][j] {
				res++
			}
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
