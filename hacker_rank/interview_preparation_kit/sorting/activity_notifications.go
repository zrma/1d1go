package sorting

import "sort"

func activityNotifications(expenditure []int32, d int32) int32 {
	// 0 <= expenditure[i] <= 200
	// NOTE - https://en.wikipedia.org/wiki/Counting_sort
	counts := make([]int32, 201)

	for _, n := range expenditure[:d] {
		counts[n]++
	}

	getMedian := medianOdd
	if d%2 == 0 {
		getMedian = medianEven
	}

	var res int32
	for i, target := range expenditure[d:] {
		median := getMedian(counts, d)
		if float64(target) >= median*2 {
			res++
		}

		nthPrev := expenditure[i]
		counts[nthPrev]--
		counts[target]++
	}
	return res
}

func medianOdd(counts []int32, d int32) float64 {
	halfCount := d / 2

	var median float64
	var countSum int32
	for num, count := range counts {
		countSum += count

		if countSum > halfCount {
			median = float64(num)
			break
		}
	}
	return median
}

func medianEven(counts []int32, d int32) float64 {
	halfCountLeft := d/2 - 1
	halfCountRight := d / 2

	var median0, median1 int
	var countSum int32
	for num, count := range counts {
		countSum += count

		if median0 == 0 && countSum > halfCountLeft {
			median0 = num
		}

		if median1 == 0 && countSum > halfCountRight {
			median1 = num
			break
		}
	}
	return float64(median0+median1) / 2
}

func activityNotificationsWithSort(expenditure []int32, d int32) int32 {
	chunk := make([]int32, d)
	var res int32
	if d%2 == 0 {
		idx0 := int(d/2 - 1)
		idx1 := int(d / 2)

		for pos := int(d); pos < len(expenditure); pos++ {
			start := pos - int(d)
			copy(chunk, expenditure[start:pos])

			mid := medianEvenWithSort(chunk, idx0, idx1)
			target := expenditure[pos]
			if float64(target) >= mid*2 {
				res++
			}
		}
	} else {
		idx := int(d / 2)

		for pos := int(d); pos < len(expenditure); pos++ {
			start := pos - int(d)
			copy(chunk, expenditure[start:pos])

			mid := medianOddWithSort(chunk, idx)
			target := expenditure[pos]
			if float64(target) >= mid*2 {
				res++
			}
		}
	}
	return res
}

func medianOddWithSort(arr []int32, idx int) float64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return float64(arr[idx])
}

func medianEvenWithSort(arr []int32, idx0, idx1 int) float64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return float64(arr[idx0]+arr[idx1]) / 2
}
