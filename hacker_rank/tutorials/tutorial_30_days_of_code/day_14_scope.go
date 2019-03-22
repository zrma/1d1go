package tutorial_30_days_of_code

import (
	"github.com/zrma/1d1c/hacker_rank/common/utils/integer_util"
	"math"
)

type Difference struct {
	elements []int
}

func NewDifference(elements []int) *Difference {
	return &Difference{elements}
}

func (d Difference) computeDifference() int {
	var min int32 = math.MaxInt32
	var max int32 = math.MinInt32

	for _, num := range d.elements {
		min = integer_util.MinInt32([]int32{min, int32(num)})
		max = integer_util.MaxInt32([]int32{max, int32(num)})
	}

	return int(max - min)
}
