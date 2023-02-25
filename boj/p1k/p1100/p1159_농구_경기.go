package p1100

import (
	"fmt"
	"io"
	"sort"
)

func Solve1159(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	players := make(map[string]int)
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)
		players[string(s[0])]++
	}

	var results []string
	for k, v := range players {
		if v >= 5 {
			results = append(results, k)
		}
	}

	if len(results) == 0 {
		//goland:noinspection SpellCheckingInspection
		_, _ = fmt.Fprint(writer, "PREDAJA")
	} else {
		sort.Strings(results)

		for _, r := range results {
			_, _ = fmt.Fprint(writer, r)
		}
	}
}
