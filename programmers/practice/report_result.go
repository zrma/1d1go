package practice

import (
	"strings"
)

func reportResult(idList []string, report []string, k int) []int {
	fromTo := make(map[string]map[string]bool)
	for _, r := range report {
		token := strings.Split(r, " ")
		from, to := token[0], token[1]
		if _, ok := fromTo[from]; !ok {
			fromTo[from] = make(map[string]bool)
		}
		fromTo[from][to] = true
	}

	counter := make(map[string]int)
	for _, toMap := range fromTo {
		for to := range toMap {
			counter[to]++
		}
	}

	res := make([]int, len(idList))
	for i, id := range idList {
		tot := 0
		if toMap, ok := fromTo[id]; ok {
			for to := range toMap {
				cnt := counter[to]
				if cnt >= k {
					tot++
				}
			}
		}
		res[i] = tot
	}
	return res
}
