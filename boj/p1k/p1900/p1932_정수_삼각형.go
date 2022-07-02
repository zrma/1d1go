package p1900

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve1932(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	sumMatrix := make([][]int, n)
	for i := range sumMatrix {
		sumMatrix[i] = make([]int, n)
	}

	inputMatrix := make([][]int, n)
	for i := range inputMatrix {
		inputMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j <= i {
				_, _ = fmt.Fscan(reader, &inputMatrix[i][j])
			}
		}
	}

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if y == 0 {
				sumMatrix[y][x] = inputMatrix[y][x]
			} else {
				neighbors := getNextMax(sumMatrix, x, y)
				sumMatrix[y][x] = inputMatrix[y][x] + neighbors
			}
		}
	}

	res := integer.Max(sumMatrix[n-1]...)
	_, _ = fmt.Fprint(writer, res)
}

func getNextMax(sumMatrix [][]int, x, y int) int {
	if x == 0 {
		return sumMatrix[y-1][x]
	}
	if x >= y {
		return sumMatrix[y-1][x-1]
	}
	return integer.Max(sumMatrix[y-1][x-1], sumMatrix[y-1][x])
}
