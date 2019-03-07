package strings

import (
	"github.com/zrma/1d1c/hacker_rank/common/utils"
	"math"
)

type validateFunc func() bool
type increaseFunc func(byte)
type decreaseFunc func(byte)

func getCalculators(l int) (increaseFunc, decreaseFunc, validateFunc) {
	cnt := l / 4
	var arr = [4]int{-cnt, -cnt, -cnt, -cnt}

	calcArr := func(b byte, margin int) {
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

	return func(b byte) {
			calcArr(b, 1)
		}, func(b byte) {
			calcArr(b, -1)
		}, func() bool {
			for _, num := range arr {
				if num > 0 {
					return false
				}
			}
			return true
		}
}

func steadyGene(gene string) int32 {
	l := len(gene)
	increase, decrease, validate := getCalculators(l)
	for _, g := range gene {
		increase(byte(g))
	}

	var begin, end int32
	var minRange int32 = math.MaxInt32
	for end < int32(l-1) {
		decrease(gene[end])
		end++

		for validate() {
			minRange = utils.MinInt32([]int32{minRange, end - begin})
			increase(gene[begin])
			begin++
		}
	}

	return minRange
}
