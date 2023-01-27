package p24000

import (
	"fmt"
	"io"
)

func Solve24060(reader io.Reader, writer io.Writer) {
	var a, k int
	_, _ = fmt.Fscan(reader, &a, &k)

	arr := make([]int, a)
	for i := 0; i < a; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	saveCount := 0
	tmp := make([]int, a+2)
	merge := func(p, q, r int) {
		var (
			i = p
			j = q + 1
			t = 1
		)
		for i <= q && j <= r {
			if arr[i] <= arr[j] {
				tmp[t] = arr[i]
				t++
				i++
			} else {
				tmp[t] = arr[j]
				t++
				j++
			}
		}
		for i <= q {
			tmp[t] = arr[i]
			t++
			i++
		}
		for j <= r {
			tmp[t] = arr[j]
			t++
			j++
		}

		i, j = p, 1
		for i <= r {
			arr[i] = tmp[j]
			saveCount++
			if saveCount == k {
				_, _ = fmt.Fprint(writer, tmp[j])
			}
			i++
			j++
		}
	}

	var mergeSort func(p, r int)
	mergeSort = func(p, r int) {
		if p >= r {
			return
		}
		q := p + (r-p)/2
		mergeSort(p, q)
		mergeSort(q+1, r)
		merge(p, q, r)
	}

	mergeSort(0, a-1)

	if saveCount < k {
		_, _ = fmt.Fprint(writer, -1)
	}
}
