package dictionariesandhashmaps

import (
	"fmt"
	"sort"
	"strings"

	"1d1go/utils/str"
)

var fmtPrint = fmt.Print

func checkMagazine(magazine, note []string) {
	sort.Sort(str.SortAdapter(magazine))
	sort.Sort(str.SortAdapter(note))

	i := len(magazine) - 1
	j := len(note) - 1

	for i >= 0 && j >= 0 {
		diff := strings.Compare(magazine[i], note[j])
		if diff > 0 {
			i--
		} else if diff == 0 {
			i--
			j--
		} else {
			_, _ = fmtPrint("No")
			return
		}
	}

	_, _ = fmtPrint("Yes")
}
