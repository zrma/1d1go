package lv1medium

import "math"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	curLength := 1
	maxLength := 1
	prevIndex := 0

	var visited [math.MaxUint8]int
	for i := range visited {
		visited[i] = -1
	}
	visited[s[0]] = 0

	// 한 문자가 다시 출현할 때 까지의 최대 거리를 구한다.
	// a b c a b c
	// |-----|
	// 0 1 2 3
	//
	// prevIndex = 0
	// curLength -> 2까지 증가
	//
	// i < 3 일 경우
	//
	// b, c의 경우 모두 prevIndex == -1
	//
	// i == 1 : 1 - 1 == 0 (prevIndex == 0)
	// i == 2 : 2 - 1 == 0 -> curLength++ -> 2
	// i == 3 : curLength(2) > maxLength(1) -> maxLength(2), curLength = 3 - 0 = 3
	//
	// ... 중략
	//
	// curLength(3) > maxLength(2) -> maxLength(3)
	for i := 1; i < len(s); i++ {
		prevIndex = visited[s[i]]
		if prevIndex == -1 || (i-curLength > prevIndex) {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			curLength = i - prevIndex
		}
		visited[s[i]] = i
	}

	if curLength > maxLength {
		maxLength = curLength
	}
	return maxLength
}
