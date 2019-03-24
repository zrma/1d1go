package tutorial_30_days_of_code

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zrma/1d1c/hacker_rank/common/utils/string_util"
)

func filter(arr [][]string) {
	var output []string
	for _, ss := range arr {
		firstName, emailId := ss[0], ss[1]
		//noinspection SpellCheckingInspection
		if strings.HasSuffix(emailId, "@gmail.com") {
			output = append(output, firstName)
		}
	}

	sort.Sort(string_util.SortAdapter(output))
	for _, s := range output {
		fmt.Println(s)
	}
}
