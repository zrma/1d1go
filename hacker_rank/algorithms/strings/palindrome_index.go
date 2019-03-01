package strings

import "github.com/zrma/1d1c/hacker_rank/common/utils"

func equalPrefix(s1, s2 string) bool {
	eq := s1[0] == s2[0]
	if len(s1) < 2 || len(s2) < 2 {
		return eq
	}

	return eq && (s1[1] == s2[1])
}

type checkData struct {
	s1, s2                                string
	forward, backward                     bool
	forwardMargin, backwardMargin, target int
}

func NewData(s1, s2 string) *checkData {
	return &checkData{s1: s1, s2: s2}
}

func checkIt(i, length int, data *checkData) bool {
	if data.s1[i+data.forwardMargin] == data.s2[i+data.backwardMargin] {
		return true
	}

	if data.forward || data.backward {
		return false
	}

	if data.s1[i] == data.s2[i+1] && equalPrefix(data.s1[i:], data.s2[i+1:]) {
		data.backwardMargin++
		data.backward = true
		data.target = length - i - 1

		return true
	}

	if data.s1[i+1] == data.s2[i] && equalPrefix(data.s1[i+1:], data.s2[i:]) {
		data.forwardMargin++
		data.forward = true
		data.target = i

		return true
	}

	return false
}

func checkSimilarity(s1, s2 string, length int) (int, bool) {
	data := NewData(s1, s2)
	for i := 0; i < length/2; i++ {
		if ok := checkIt(i, length, data); !ok {
			return -1, false
		}
	}

	return data.target, data.forward || data.backward
}

func palindromeIndex(s string) int32 {
	l := len(s)
	if l < 2 {
		return -1
	}

	re := utils.Reverse(s)
	result, similar := checkSimilarity(s, re, l)
	if !similar {
		return -1
	}

	return int32(result)
}
