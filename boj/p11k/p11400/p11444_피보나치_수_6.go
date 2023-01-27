package p11400

import (
	"fmt"
	"io"
)

func Solve11444(reader io.Reader, writer io.Writer) {
	var n int64
	_, _ = fmt.Fscan(reader, &n)

	// F0 = 0
	// F1 = 1
	// F2 = 1 = F0 + F1 = 0 + 1
	// F3 = 2 = F1 + F2 = 1 + 1 = 2
	// F4 = 3 = F2 + F3 = 1 + 2 = 3
	//
	// Fn   = Fn-1 + Fn-2
	// Fn-1 = Fn-2 + Fn-3
	//
	// Fn   = | 1, 1 | * | Fn-1 |
	//                   | Fn-2 |
	// Fn-1 = | 1, 0 | * | Fn-1 |
	//                   | Fn-2 |
	//
	// | Fn   | = | 1, 1 | * | Fn-1 |
	// | Fn-1 |   | 1, 0 |   | Fn-2 |
	//
	// | Fn-1 | = | 1, 1 | * | Fn-2 |
	// | Fn-2 |   | 1, 0 |   | Fn-3 |
	//
	// ...
	//
	// | Fn   | = | 1, 1 | * | Fn-1 | = | 1, 1 | * | 1, 1 | * | Fn-2 |
	// | Fn-1 |   | 1, 0 |   | Fn-2 |   | 1, 0 |   | 1, 0 |   | Fn-3 |
	//
	//          = | 1, 1 |^2 * | Fn-2 |
	//            | 1, 0 |     | Fn-3 |
	//
	//          = | 1, 1 |^n-1 * | Fn-(n-1) |
	//            | 1, 0 |       | Fn-n     |
	//
	//          = | 1, 1 |^n-1 * | F1 |
	//            | 1, 0 |       | F0 |
	//
	//          = | 1, 1 |^n-1 * | 1 |
	//            | 1, 0 |       | 0 |
	//
	// | Fn+1 | = | 1, 1 |^n-1 * | F2 | = | 1, 1 |^n-1 * | 1 |
	// | Fn   |   | 1, 0 |       | F1 |   | 1, 0 |       | 1 |
	//
	// | Fn+1, Fn   | = | 1, 1 |^n-1 * | 1 1 | = | 1, 1 |^n
	// | Fn  , Fn-1 |   | 1, 0 |       | 1 0 |   | 1, 0 |
	mat := [][]int64{
		{1, 1},
		{1, 0},
	}
	cache = make(map[int64][][]int64)

	res := matPow(mat, n)
	_, _ = fmt.Fprint(writer, res[0][1])
}

var cache map[int64][][]int64

const mod = 1_000_000_007

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
