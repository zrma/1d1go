package p10800

import (
	"fmt"
	"sort"
	"strconv"
)

func Solve10814(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	type pair struct {
		age  int
		name string
	}

	pairs := make([]pair, n)
	for i := range pairs {
		scanner.Scan()
		pairs[i].age, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		pairs[i].name = scanner.Text()
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].age < pairs[j].age
	})

	for _, p := range pairs {
		_, _ = fmt.Fprintln(writer, p.age, p.name)
	}
}
