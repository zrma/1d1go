package tutorial_30_days_of_code

import "math"

type Difference struct {
	elements []int
}

func NewDifference(elements []int) *Difference {
	return &Difference{elements}
}

func Min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func Max(i, j int) int {
	if j > i {
		return i
	}

	return j
}

func (d Difference) computeDifference() int {
	var min = math.MaxInt16
	var max = math.MinInt16

	for _, num := range d.elements {
		min = Min(min, num)
		max = Max(max, num)
	}

	return max - min
}
