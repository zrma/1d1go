package p1500

import (
	"fmt"
	"io"
)

func Solve1508(reader io.Reader, writer io.Writer) {
	var n, m, k int
	_, _ = fmt.Fscan(reader, &n, &m, &k)

	arr := make([]int, k)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	res := ""

	lo := 0
	hi := n
	for lo <= hi {
		term := lo + (hi-lo)/2

		layout := "1"
		pick := 1
		prev := 0

		// term 간격으로, k 개의 위치 중에 차례대로 m개 배치
		for curr := 1; curr < k; curr++ {
			distance := arr[curr] - arr[prev]
			if distance >= term {
				prev = curr
				layout += "1"
				pick++

				// m명 배치하면 배치 종료
				if pick == m {
					break
				}
			} else {
				layout += "0"
			}
		}

		for k > len(layout) {
			layout += "0"
		}

		if pick == m {
			// m명 배치할 수 있었음
			// 더 큰 term을 찾아보자
			lo = term + 1
			res = layout
		} else {
			// m명 배치할 수 없었음
			// 더 작은 term을 찾아보자
			hi = term - 1
		}
	}

	_, _ = fmt.Fprint(writer, res)
}
