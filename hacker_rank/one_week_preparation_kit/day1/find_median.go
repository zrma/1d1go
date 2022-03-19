package day1

import "sort"

func findMedian(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[len(arr)/2]
}
