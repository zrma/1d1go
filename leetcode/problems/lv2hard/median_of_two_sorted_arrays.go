package lv2hard

import "math"

func findTwoValuesNearMid(nums1, nums2 []int, totLen int) (prev, curr int) {
	target := totLen / 2
	var idx1, idx2, idxTot int
	for idx1 < len(nums1) && idx2 < len(nums2) {
		prev = curr
		if nums1[idx1] < nums2[idx2] {
			curr = nums1[idx1]
			idxTot++
			idx1++
		} else {
			curr = nums2[idx2]
			idxTot++
			idx2++
		}
		if idxTot == target+1 {
			return
		}
	}

	remain := target - idxTot
	if idx1 < len(nums1) {
		if idx1+remain > 0 {
			prev = nums1[idx1+remain-1]
		} else {
			prev = nums2[idx2-1]
		}
		curr = nums1[idx1+remain]
	} else {
		if idx2+remain > 0 {
			prev = nums2[idx2+remain-1]
		} else {
			prev = nums1[idx1-1]
		}
		curr = nums2[idx2+remain]
	}
	return
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	totLen := len(nums1) + len(nums2)
	val1, val2 := findTwoValuesNearMid(nums1, nums2, totLen)
	if totLen%2 == 0 {
		return float64(val1+val2) / 2
	}
	return float64(val2)
}

func mergeSortedArray(nums1, nums2 []int) []int {
	totLen := len(nums1) + len(nums2)
	totNums := make([]int, totLen)

	var i, j, cur int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			totNums[cur] = nums1[i]
			cur++
			i++
		} else {
			totNums[cur] = nums2[j]
			cur++
			j++
		}
	}
	for i < len(nums1) {
		totNums[cur] = nums1[i]
		cur++
		i++

	}
	for j < len(nums2) {
		totNums[cur] = nums2[j]
		cur++
		j++
	}
	return totNums
}

func findMedianSortedArraysWithMerge(nums1, nums2 []int) float64 {
	totNums := mergeSortedArray(nums1, nums2)
	if len(totNums)%2 == 0 {
		return float64(totNums[len(totNums)/2-1]+totNums[len(totNums)/2]) / 2
	} else {
		return float64(totNums[len(totNums)/2])
	}
}

type side struct {
	nums1 int
	nums2 int
}

type arr []int

func (a arr) getVal(idx int) int {
	if idx < 0 {
		return math.MinInt32
	}
	if idx >= len(a) {
		return math.MaxInt32
	}
	return a[idx]
}

func findBoundary(nums1, nums2 []int, totalCount int) (leftMax, rightMin side) {
	// range of binary search target in nums1
	lo, hi := 0, len(nums1)

	for lo <= hi {
		// nums1:  1 2 4 | 5 6 7
		// nums2:  3 5   | 8 9
		//
		// totalCount := 6 + 4
		// cut1 := (0 + 6) / 2 = 3
		// cut2 := (totalCount+1) / 2 - cut1 = 5 - 3 = 2
		//
		// 총 10개이므로 좌우측 중 한 쪽에서 5개씩 가져갈 수 있음
		// totalCount 에 더해주는 +1 은 인덱스 계산을 편하게 하기 위한 sugar
		// 좌측의 nums1에서 먼저 3개를 가져가면 nums2는 2개만 가져갈 수 있음
		//
		// it makes both side(left and right) being equal count
		cut1 := lo + (hi-lo)/2
		cut2 := (totalCount+1)/2 - cut1

		leftMax.nums1 = arr(nums1).getVal(cut1 - 1)
		leftMax.nums2 = arr(nums2).getVal(cut2 - 1)
		rightMin.nums1 = arr(nums1).getVal(cut1)
		rightMin.nums2 = arr(nums2).getVal(cut2)

		if leftMax.nums1 <= rightMin.nums2 && leftMax.nums2 <= rightMin.nums1 {
			break
		}

		if leftMax.nums1 >= rightMin.nums2 {
			hi = cut1 - 1
		} else {
			lo = cut1 + 1
		}
	}
	return
}

func findMedianSortedArraysWithBinSearch(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	totalCount := len(nums1) + len(nums2)
	leftMax, rightMin := findBoundary(nums1, nums2, totalCount)

	leftMaxVal := math.Max(float64(leftMax.nums1), float64(leftMax.nums2))
	if totalCount%2 == 0 {
		rightMinVal := math.Min(float64(rightMin.nums1), float64(rightMin.nums2))
		return leftMaxVal + (rightMinVal-leftMaxVal)/2
	} else {
		return leftMaxVal
	}
}

func findMedianSortedArraysWithBinSearchSolution(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	half := (len(nums1) + len(nums2) - 1) / 2

	l, r := 0, len(nums1)
	i := l + (r-l)/2
	j := half - i
	for l < r {
		if nums1[i] < nums2[j] {
			l = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j+1] {
			r = i
		} else {
			break
		}
		i = l + (r-l)/2
		j = half - i
	}

	var medianLeft, medianRight float64
	if j < 0 {
		medianLeft = float64(nums1[i-1])
	} else if i < 1 {
		medianLeft = float64(nums2[j])
	} else {
		medianLeft = math.Max(float64(nums1[i-1]), float64(nums2[j]))
	}

	if (len(nums1)+len(nums2))%2 == 1 {
		return medianLeft
	} else if i == len(nums1) {
		medianRight = float64(nums2[j+1])
	} else if j+1 == len(nums2) {
		medianRight = float64(nums1[i])
	} else {
		medianRight = math.Min(float64(nums1[i]), float64(nums2[j+1]))
	}
	return medianLeft + (medianRight-medianLeft)/2
}
