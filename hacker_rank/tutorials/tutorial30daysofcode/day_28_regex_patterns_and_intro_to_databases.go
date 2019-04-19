package tutorial30daysofcode

import (
	"fmt"
	"github.com/zrma/1d1c/hacker_rank/common/utils/str"
	"sort"
	"strings"
)

func filter(arr [][]string) {
	var output []string
	for _, ss := range arr {
		firstName, emailID := ss[0], ss[1]
		//noinspection SpellCheckingInspection
		if strings.HasSuffix(emailID, "@gmail.com") {
			output = append(output, firstName)
		}
	}

	sort.Sort(str.SortAdapter(output))
	for _, s := range output {
		fmt.Println(s)
	}
}
