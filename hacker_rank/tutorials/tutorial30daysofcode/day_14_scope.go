package tutorial30daysofcode

import (
	"math"

	"1d1go/utils/integer"
)

type difference struct {
	elements []int
}

func newDifference(elements []int) *difference {
	return &difference{elements}
}

func (d difference) computeDifference() int {
	var min = math.MaxInt
	var max = math.MinInt

	for _, num := range d.elements {
		min = integer.Min(min, num)
		max = integer.Max(max, num)
	}

	return max - min
}
