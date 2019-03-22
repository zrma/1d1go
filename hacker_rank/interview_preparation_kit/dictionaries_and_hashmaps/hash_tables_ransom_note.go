package dictionaries_and_hashmaps

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zrma/1d1c/hacker_rank/common/utils/string_util"
)

func checkMagazine(magazine []string, note []string) {
	sort.Sort(string_util.SortAdapter(magazine))
	sort.Sort(string_util.SortAdapter(note))

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
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
