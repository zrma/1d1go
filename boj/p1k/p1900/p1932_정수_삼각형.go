package p1900

import (
	"fmt"
	"strconv"

	"1d1go/utils/integer"
)

func Solve1932(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	sumMatrix := make([][]int, n)
	for i := range sumMatrix {
		sumMatrix[i] = make([]int, n)
	}

	inputMatrix := make([][]int, n)
	for i := range inputMatrix {
		inputMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j <= i {
				scanner.Scan()
				inputMatrix[i][j], _ = strconv.Atoi(scanner.Text())
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
