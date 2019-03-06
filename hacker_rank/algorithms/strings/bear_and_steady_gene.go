package strings

import (
	"github.com/zrma/1d1c/hacker_rank/common/utils"
	"math"
)

func isSteady(arr [4]int, validCnt int) bool {
	for _, num := range arr {
		if num > validCnt {
			return false
		}
	}
	return true
}

func calcArr(b byte, arr *[4]int, increase bool) {
	margin := 1
	if !increase {
		margin = -1
	}
	switch b {
	case 'A':
		arr[0] += margin
	case 'C':
		arr[1] += margin
	case 'T':
		arr[2] += margin
	case 'G':
		arr[3] += margin
	}
}

func steadyGene(gene string) int32 {
	validCnt := len(gene) / 4
	var arr [4]int
	for i := range gene {
		calcArr(gene[i], &arr, true)
	}

	var begin, end int32
	var minRange int32 = math.MaxInt32
	for end < int32(len(gene)-1) {
		calcArr(gene[end], &arr, false)
		end++

		for isSteady(arr, validCnt) {
			minRange = utils.MinInt32([]int32{minRange, end - begin})
			calcArr(gene[begin], &arr, true)
			begin++
		}
	}

	return minRange
}
