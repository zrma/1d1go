package p1400

import (
	"fmt"
	"sort"
)

func Solve1427(scanner Scanner, writer Writer) {
	scanner.Scan()
	s := scanner.Text()

	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})

	_, _ = fmt.Fprint(writer, string(b))
}
