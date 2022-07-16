package p1900

import (
	"fmt"
)

func Solve1966(reader Reader, writer Writer) {
	var count int
	_, _ = fmt.Fscan(reader, &count)

	for i := 0; i < count; i++ {
		printerQueue(reader, writer)
	}
}

func printerQueue(reader Reader, writer Writer) {
	var n, target int
	_, _ = fmt.Fscan(reader, &n, &target)

	type entry struct {
		v      int
		target bool
	}

	arr := make([]entry, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i].v)
		if i == target {
			arr[i].target = true
		}
	}

	count := 1
	for len(arr) > 0 {
		cur := arr[0]
		found := false
		for _, other := range arr[1:] {
			if other.v > cur.v {
				found = true
				break
			}
		}
		if found {
			arr = append(arr[1:], arr[0])
		} else {
			if cur.target {
				_, _ = fmt.Fprintln(writer, count)
				return
			}
			arr = arr[1:]
			count++
		}
	}
}
