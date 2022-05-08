package day2

import (
	"1d1go/utils/integer"
)

func flipMatrix(matrix [][]int32) int32 {
	n := len(matrix) / 2
	lastIndex := len(matrix) - 1
	var res int32
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			res += integer.MaxInt32(
				matrix[row][col],                     // top left
				matrix[row][lastIndex-col],           // top right
				matrix[lastIndex-row][col],           // bottom left
				matrix[lastIndex-row][lastIndex-col], // bottom right
			)
		}
	}
	return res
}
