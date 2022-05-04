package lv0easy

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}
	}
	return result
}

func twoSumSorted(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum == target {
			break
		}
		if sum < target {
			lo++
		}
		if sum > target {
			hi--
		}
	}
	return []int{lo, hi}
}

func twoSumHash(nums []int, target int) []int {
	m := make(map[int]int)
	for curIdx, num := range nums {
		if prevIdx, ok := m[target-num]; ok {
			return []int{prevIdx, curIdx}
		}
		m[num] = curIdx
	}
	return []int{}
}

func twoSum2InputArrayIsSorted(nums []int, target int) []int {
	result := twoSumSorted(nums, target)
	for i := range result {
		result[i]++
	}
	return result
}
