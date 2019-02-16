package dictionaries_and_hashmaps

import "github.com/zrma/1d1c/hacker_rank/common/utils"

func twoStrings(s1 string, s2 string) string {
	s1 = utils.SortString(s1)
	s2 = utils.SortString(s2)

	var i, j int
	for i < len(s1)-1 && j < len(s2)-1 {
		diff := int(s1[i]) - int(s2[j])

		if diff < 0 {
			i++
		} else if diff > 0 {
			j++
		} else {
			return "YES"
		}
	}

	return "NO"
}
