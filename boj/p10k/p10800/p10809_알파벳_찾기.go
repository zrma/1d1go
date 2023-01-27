package p10800

import (
	"fmt"
	"io"
)

func Solve10809(reader io.Reader, writer io.Writer) {
	const notFound = -1
	arr := [26]int{}
	for i := range arr {
		arr[i] = notFound
	}

	var s string
	_, _ = fmt.Fscan(reader, &s)

	for i, c := range s {
		if arr[c-'a'] == notFound {
			arr[c-'a'] = i
		}
	}

	for _, n := range arr {
		_, _ = fmt.Fprintf(writer, "%d ", n)
	}
}
