package solve

import (
	"math/bits"
)

func combinations(total, num int) [][]int {
	var set []int
	for i := 0; i < total; i++ {
		set = append(set, i)
	}

	length := uint(len(set))
	var subsets [][]int
	var subsetBits uint = 1
	for ; subsetBits < (1 << length); subsetBits++ {
		if bits.OnesCount(subsetBits) != num {
			continue
		}
		var subset []int
		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

type array []int

func (a array) find(n int) bool {
	for _, num := range a {
		if num == n {
			return true
		}
	}
	return false
}

func buildSteps(total, num int) [][]bool {
	var steps [][]bool
	for _, arr := range combinations(total, num) {
		var step []bool
		for i := 0; i < total; i++ {
			if array(arr).find(i) {
				step = append(step, false)
			} else {
				step = append(step, true)
			}
		}
		steps = append(steps, step)
	}
	return steps
}

func buildMat(downWay, upWay []bool, row, col int) [][]int {
	var mat [][]int
	for i := 0; i < col; i++ {
		mat = append(mat, make([]int, row))
	}

	mat[0][0] = 1

	posX, posY := 0, 0
	for _, step := range downWay {
		if step {
			posX++
		} else {
			posY++
		}
		mat[posY][posX] = 1
	}

	posX, posY = 0, 0
	for _, step := range upWay {
		if step {
			posX++
		} else {
			posY++
		}
		mat[posY][posX] = 1
	}

	return mat
}

func mulMat(mat1, mat2 [][]int) (int, bool) {
	row := len(mat1)
	col := len(mat1[0])

	total := 0
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			v1, v2 := mat1[x][y], mat2[y][x]
			if v2 == 1 && v1 == -1 {
				return -1, false
			}
			total += v1 * v2
		}
	}
	return total, true
}

func collectMax(mat [][]int) int {
	row := len(mat)
	if row < 1 {
		return 0
	}

	col := len(mat[0])
	if col < 1 {
		return 0
	}

	if mat[row-1][col-1] == -1 {
		return 0
	}

	rowStep := row - 1
	colStep := col - 1
	totalStep := rowStep + colStep

	downWays := buildSteps(totalStep, colStep)
	upWays := buildSteps(totalStep, colStep)

	highestScore := 0
	for _, downWay := range downWays {
		for _, upWay := range upWays {
			mat2 := buildMat(downWay, upWay, row, col)
			if score, ok := mulMat(mat, mat2); ok && (highestScore < score) {
				highestScore = score
			}
		}
	}

	return highestScore
}
