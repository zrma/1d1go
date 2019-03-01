package strings

import "github.com/zrma/1d1c/hacker_rank/common/utils"

func equalPrefix(s1, s2 string) bool {
	eq := s1[0] == s2[0]
	if len(s1) < 2 || len(s2) < 2 {
		return eq
	}

	return eq && (s1[1] == s2[1])
}

func palindromeIndex(s string) int32 {
	l := len(s)
	if l < 2 {
		return -1
	}

	re := utils.Reverse(s)

	var forward, backward bool
	var forwardMargin, backwardMargin, target int
	for i := 0; i < l/2; i++ {
		if s[i+forwardMargin] != re[i+backwardMargin] {
			if forward || backward {
				return -1
			}

			if s[i] == re[i+1] {
				if equalPrefix(s[i:], re[i+1:]) {
					backwardMargin++
					backward = true
					target = l - i - 1

					continue
				}
			}

			if s[i+1] == re[i] {
				if equalPrefix(s[i+1:], re[i:]) {
					forwardMargin++
					forward = true
					target = i

					continue
				}
			}

			return -1
		}
	}

	if !(forward || backward) {
		return -1
	}

	return int32(target)
}
