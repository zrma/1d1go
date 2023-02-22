package p1300

import (
	"fmt"
	"io"
	"sort"
)

func Solve1302(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	books := make(map[string]int)
	for i := 0; i < n; i++ {
		var book string
		_, _ = fmt.Fscanln(reader, &book)
		books[book]++
	}

	max := 0
	for _, v := range books {
		if v > max {
			max = v
		}
	}

	var best []string
	for k, v := range books {
		if v == max {
			best = append(best, k)
		}
	}

	sort.Strings(best)
	_, _ = fmt.Fprint(writer, best[0])
}
