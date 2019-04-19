package tutorial30daysofcode

import (
	"github.com/zrma/1d1c/hacker_rank/common/utils/integer"
	"math"
)

type difference struct {
	elements []int
}

func newDifference(elements []int) *difference {
	return &difference{elements}
}

func (d difference) computeDifference() int {
	var min int32 = math.MaxInt32
	var max int32 = math.MinInt32

	for _, num := range d.elements {
		min = integer.MinInt32([]int32{min, int32(num)})
		max = integer.MaxInt32([]int32{max, int32(num)})
	}

	return int(max - min)
}
