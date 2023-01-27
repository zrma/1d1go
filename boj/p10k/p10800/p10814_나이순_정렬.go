package p10800

import (
	"fmt"
	"io"
	"sort"
)

func Solve10814(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	type pair struct {
		age  int
		name string
	}

	pairs := make([]pair, n)
	for i := range pairs {
		_, _ = fmt.Fscan(reader, &pairs[i].age, &pairs[i].name)
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].age < pairs[j].age
	})

	for _, p := range pairs {
		_, _ = fmt.Fprintln(writer, p.age, p.name)
	}
}
