package strings

import "github.com/zrma/1d1c/hacker_rank/common/utils"

func equalPrefix(s1, s2 string) bool {
	eq := s1[0] == s2[0]
	if len(s1) < 2 || len(s2) < 2 {
		return eq
	}

	return eq && (s1[1] == s2[1])
}

func getDiffChecker(s1, s2 string, length int) func(int) (int, bool, bool) {
	var forward, backward bool
	var forwardMargin, backwardMargin, target int

	isDiff := func() bool {
		return forward || backward
	}

	return func(curPos int) (int, bool, bool) {
		// It is a palindrome
		if s1[curPos+forwardMargin] == s2[curPos+backwardMargin] {
			return target, isDiff(), true
		}

		// The chance to remove the rune was given only once, but it was consumed before.
		if isDiff() {
			return -1, isDiff(), false
		}

		// The following two quadrants should be concurrent, not sequential,
		// so you should check up to the second candidate(use function equalPrefix)

		// If the difference can be avoided by a single rune difference in backward,
		// mark it in backward and increase backwardMargin.
		if s1[curPos] == s2[curPos+1] && equalPrefix(s1[curPos:], s2[curPos+1:]) {
			backwardMargin++
			backward = true
			target = length - curPos - 1

			return target, isDiff(), true
		}

		// Else if the difference can be avoided by a single rune difference in forward,
		// mark it in forward and increase forwardMargin.
		if s1[curPos+1] == s2[curPos] && equalPrefix(s1[curPos+1:], s2[curPos:]) {
			forwardMargin++
			forward = true
			target = curPos

			return target, isDiff(), true
		}

		return -1, isDiff(), false
	}
}

func getDiffPos(s1, s2 string, length int) (target int, diff bool) {

	diffChecker := getDiffChecker(s1, s2, length)
	var valid bool
	for i := 0; i < length/2; i++ {
		if target, diff, valid = diffChecker(i); !valid {
			// The input string can not be a palindrome, even if one rune is removed.
			return -1, diff
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
	pos, diff := getDiffPos(s, re, l)
	if !diff {
		// The input string is already a palindrome. You can not remove any rune.
		return -1
	}

	return int32(pos)
}
