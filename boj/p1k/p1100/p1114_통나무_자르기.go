package p1100

import (
	"fmt"
	"io"
	"sort"
)

func Solve1114(reader io.Reader, writer io.Writer) {
	var l, k, c int
	_, _ = fmt.Fscan(reader, &l, &k, &c)

	arr := make([]int, k)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	// 처음과 끝 추가
	arr = append(arr, 0, l)

	sort.Ints(arr)

	dists := make([]int, len(arr)-1)
	for i := 1; i < len(arr); i++ {
		dists[i-1] = arr[i] - arr[i-1]
	}
	sort.Ints(dists)

	lo := dists[0]
	hi := l

	for lo < hi {
		dist := lo + (hi-lo)/2
		if cut1114(arr, dist, c) {
			// 더 작은 간격으로 잘라봄
			hi = dist
		} else {
			// 자를 수 있는 횟수가 모자라다.
			// 더 큰 간격으로 잘라야 함
			lo = dist + 1
		}
	}
	longest := hi

	curr := len(arr) - 1
	firstCut := arr[curr]
	cutCnt := 0
	for i := curr - 2; i >= 0; i-- {
		if dist := arr[curr] - arr[i]; dist > longest {
			cutCnt++
			i++
			firstCut = arr[i]
			curr = i
		}
	}
	if cutCnt < c {
		firstCut = arr[1]
	}
	_, _ = fmt.Fprint(writer, longest, firstCut)
}

func cut1114(points []int, dist, cntLimit int) bool {
	cnt := 0
	prev := 0
	for i := 1; i < len(points); i++ {
		if cnt > cntLimit {
			return false
		}

		currDist := points[i] - points[prev]
		if currDist > dist {
			cnt++
			prev = i - 1
			i--
		}
	}
	return true
}
