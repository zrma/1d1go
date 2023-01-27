package p10800

import (
	"fmt"
	"io"
)

func Solve10830(reader io.Reader, writer io.Writer) {
	var n int
	var b int64
	_, _ = fmt.Fscan(reader, &n, &b)

	mat := make([][]int64, n)
	for i := range mat {
		mat[i] = make([]int64, n)
		for j := range mat[i] {
			_, _ = fmt.Fscan(reader, &mat[i][j])
			mat[i][j] %= mod
		}
	}

	cache = make(map[int64][][]int64)

	res := matPow(mat, b)
	for i := range res {
		for j := range res[i] {
			if j < n-1 {
				_, _ = fmt.Fprint(writer, res[i][j]%mod, " ")
			} else {
				_, _ = fmt.Fprintln(writer, res[i][j]%mod)
			}
		}
	}
}

var cache map[int64][][]int64

const mod = 1_000

func matPow(mat [][]int64, b int64) [][]int64 {
	if b < 2 {
		return mat
	}
	if v, ok := cache[b]; ok {
		return v
	}

	if b%2 == 0 {
		v := matMul(matPow(mat, b/2), matPow(mat, b/2))
		cache[b] = v
		return v
	} else {
		v := matMul(matPow(mat, b-1), mat)
		cache[b] = v
		return v
	}
}

func matMul(mat0, mat1 [][]int64) [][]int64 {
	n := len(mat0)
	tmp := make([][]int64, n)
	for i := range tmp {
		tmp[i] = make([]int64, n)
	}

	for i := 0; i < len(mat0); i++ {
		for j := 0; j < len(mat1[0]); j++ {
			for k := 0; k < len(mat1); k++ {
				tmp[i][j] += mat0[i][k] * mat1[k][j]
			}
			tmp[i][j] %= mod
		}
	}
	return tmp
}
