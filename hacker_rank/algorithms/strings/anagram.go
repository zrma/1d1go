package strings

import (
	"github.com/zrma/1d1c/hacker_rank/common/utils/str"
)

func anagram(s string) int32 {
	l := len(s)
	if l < 1 || l%2 == 1 {
		return -1
	}

	s1 := str.Sort(s[0 : l/2])
	s2 := str.Sort(s[l/2:])

	l = l / 2
	var eq int32
	var i, j int
	for i < l && j < l {
		diff := int(s1[i]) - int(s2[j])

		if diff < 0 {
			i++
		} else if diff > 0 {
			j++
		} else {
			eq++
			i++
			j++
		}
	}

	return int32(l) - eq
}
