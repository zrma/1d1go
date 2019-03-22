package tutorial_30_days_of_code

import "errors"

func minimumIndex(seq []int) (int, error) {
	if len(seq) == 0 {
		return -1, errors.New("cannot get the minimum value index from an empty sequence")
	}
	var minIdx = 0
	for i := 1; i < len(seq); i++ {
		if seq[i] < seq[minIdx] {
			minIdx = i
		}
	}
	return minIdx, nil
}

func emptyArray() []int {
	return nil
}

func uniqueValues() []int {
	return []int{1, 2}
}

func exactlyTwoDifferentMinimums() []int {
	return []int{2, 1, 1}
}
