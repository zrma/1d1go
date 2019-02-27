package tutorial_30_days_of_code

import (
	"github.com/zrma/1d1c/hacker_rank/common/utils"
	"math"
)

type Difference struct {
	elements []int
}

func NewDifference(elements []int) *Difference {
	return &Difference{elements}
}

func (d Difference) computeDifference() int {
	var min = math.MaxInt32
	var max = math.MinInt32

	for _, num := range d.elements {
		min = utils.Min(min, num)
		max = utils.Max(max, num)
	}

	return max - min
}
