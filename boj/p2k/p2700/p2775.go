package p2700

import (
	"fmt"
	"strconv"
)

func Solve2775(scanner Scanner, writer Writer) {
	scanner.Scan()
	cnt, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < cnt; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		res := calcPeopleCounts(k, n)
		_, _ = fmt.Fprintln(writer, res)
	}
}

func calcPeopleCounts(row, col int) int {
	const max = 15
	arr := [max][max]int{}
	for x := 1; x <= col; x++ {
		arr[0][x] = x
	}

	for y := 1; y <= row; y++ {
		arr[y][1] = 1
		for x := 2; x <= col; x++ {
			arr[y][x] = arr[y-1][x] + arr[y][x-1]
		}
	}
	return arr[row][col]
}
