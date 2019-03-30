package strings

import "github.com/zrma/1d1c/hacker_rank/common/utils/string_util"

func anagram(s string) int32 {
	l := len(s)
	if l < 1 || l%2 == 1 {
		return -1
	}

	s1 := string_util.Sort(s[0 : l/2])
	s2 := string_util.Sort(s[l/2:])

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
