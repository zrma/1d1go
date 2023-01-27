package p12000

import (
	"fmt"
	"io"
	"sort"
)

func Solve12018(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	credits := make([]int, n)
	for i := range credits {
		var p, l int
		_, _ = fmt.Fscan(reader, &p, &l)

		person := make([]int, p)
		for j := range person {
			_, _ = fmt.Fscan(reader, &person[j])
		}
		sort.Ints(person)

		if p < l {
			credits[i] = 1
		} else {
			credits[i] = person[p-l]
		}
	}
	sort.Ints(credits)

	res := 0
	for i := range credits {
		if credits[i] <= m {
			res++
			m -= credits[i]
		} else {
			break
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
