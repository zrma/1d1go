package p10800

import (
	"fmt"
)

func Solve10809(scanner Scanner, writer Writer) {
	const notFound = -1
	arr := [26]int{}
	for i := range arr {
		arr[i] = notFound
	}

	scanner.Scan()
	s := scanner.Text()

	for i, c := range s {
		if arr[c-'a'] == notFound {
			arr[c-'a'] = i
		}
	}

	for _, n := range arr {
		_, _ = fmt.Fprintf(writer, "%d ", n)
	}
}
