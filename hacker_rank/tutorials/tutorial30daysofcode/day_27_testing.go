package tutorial30daysofcode

import "errors"

const (
	errMsg       = "cannot get the minimum value index from an empty sequence"
	invalidIndex = -1
)

func minimumIndex(arr []int) (int, error) {
	if len(arr) == 0 {
		return invalidIndex, errors.New(errMsg)
	}
	var minIndex = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex, nil
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
