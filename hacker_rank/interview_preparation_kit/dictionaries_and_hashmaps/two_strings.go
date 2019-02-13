package dictionaries_and_hashmaps

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

func twoStrings(s1 string, s2 string) string {
	s1 = sortString(s1)
	s2 = sortString(s2)

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
