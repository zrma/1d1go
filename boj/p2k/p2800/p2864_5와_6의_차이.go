package p2800

import (
	"fmt"
	"io"
)

func Solve2864(reader io.Reader, writer io.Writer) {
	var a, b string
	_, _ = fmt.Fscanln(reader, &a, &b)

	minA := make([]rune, len(a))
	maxA := make([]rune, len(a))
	minB := make([]rune, len(b))
	maxB := make([]rune, len(b))

	for i, c := range a {
		minA[i] = c
		maxA[i] = c
		switch c {
		case '5':
			maxA[i] = '6'
		case '6':
			minA[i] = '5'
		}
	}

	for i, c := range b {
		minB[i] = c
		maxB[i] = c
		switch c {
		case '5':
			maxB[i] = '6'
		case '6':
			minB[i] = '5'
		}
	}

	var num0, num1 int
	_, _ = fmt.Sscanf(string(minA), "%d", &num0)
	_, _ = fmt.Sscanf(string(minB), "%d", &num1)
	min := num0 + num1

	_, _ = fmt.Sscanf(string(maxA), "%d", &num0)
	_, _ = fmt.Sscanf(string(maxB), "%d", &num1)
	max := num0 + num1

	_, _ = fmt.Fprint(writer, min, max)
}
