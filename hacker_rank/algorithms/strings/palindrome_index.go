package strings

import "github.com/zrma/1d1c/hacker_rank/common/utils"

func equalPrefix(s1, s2 string) bool {
	eq := s1[0] == s2[0]
	if len(s1) < 2 || len(s2) < 2 {
		return eq
	}

	return eq && (s1[1] == s2[1])
}

func getChecker(s1, s2 string, length int) func(int) (int, bool, bool) {
	var forward, backward bool
	var forwardMargin, backwardMargin, target int

	return func(i int) (int, bool, bool) {
		if s1[i+forwardMargin] == s2[i+backwardMargin] {
			return target, forward || backward, true
		}

		if forward || backward {
			return -1, false, false
		}

		if s1[i] == s2[i+1] && equalPrefix(s1[i:], s2[i+1:]) {
			backwardMargin++
			backward = true
			target = length - i - 1

			return target, forward || backward, true
		}

		if s1[i+1] == s2[i] && equalPrefix(s1[i+1:], s2[i:]) {
			forwardMargin++
			forward = true
			target = i

			return target, forward || backward, true
		}

		return -1, false, false
	}
}

func getDiffPos(s1, s2 string, length int) (target int, diff bool) {

	checker := getChecker(s1, s2, length)
	var ok bool
	for i := 0; i < length/2; i++ {
		if target, diff, ok = checker(i); !ok {
			return -1, false
		}
	}

	return target, diff
}

func palindromeIndex(s string) int32 {
	l := len(s)
	if l < 2 {
		return -1
	}

	re := utils.Reverse(s)
	result, ok := getDiffPos(s, re, l)
	if !ok {
		return -1
	}

	return int32(result)
}
