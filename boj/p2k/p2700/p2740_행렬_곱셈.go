package p2700

import (
	"fmt"
)

func Solve2740(reader Reader, writer Writer) {
	var n, m, k int
	_, _ = fmt.Fscan(reader, &n, &m)

	mat0 := make([][]int, n)
	for i := range mat0 {
		mat0[i] = make([]int, m)
		for j := range mat0[i] {
			_, _ = fmt.Fscan(reader, &mat0[i][j])
		}
	}

	_, _ = fmt.Fscan(reader, &m, &k)
	mat1 := make([][]int, m)
	for i := range mat1 {
		mat1[i] = make([]int, k)
		for j := range mat1[i] {
			_, _ = fmt.Fscan(reader, &mat1[i][j])
		}
	}

	mat2 := make([][]int, n)
	for i := range mat2 {
		mat2[i] = make([]int, k)
	}

	for _, row := range matMul(mat0, mat1, mat2) {
		for j, col := range row {
			if j < len(row)-1 {
				_, _ = fmt.Fprint(writer, col, " ")
			} else {
				_, _ = fmt.Fprintln(writer, col)
			}
		}
	}
}

func matMul(mat0, mat1, mat2 [][]int) [][]int {
	for i := 0; i < len(mat0); i++ {
		for j := 0; j < len(mat1[0]); j++ {
			for k := 0; k < len(mat1); k++ {
				mat2[i][j] += mat0[i][k] * mat1[k][j]
			}
		}
	}
	return mat2
}
