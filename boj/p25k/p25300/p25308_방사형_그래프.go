package p25300

import (
	"fmt"
	"io"
)

const maxLen = 8

func Solve25308(reader io.Reader, writer io.Writer) {
	var arr [maxLen]int
	for i := 0; i < maxLen; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	var visited [maxLen]bool
	comb := make([]int, 0, maxLen)
	res := combination(arr, visited, comb, 0)
	_, _ = fmt.Fprint(writer, res)
}

func combination(arr [8]int, visited [8]bool, comb []int, depth int) int {
	if len(comb) == maxLen {
		if isConcave(comb[0], comb[1], comb[2]) &&
			isConcave(comb[1], comb[2], comb[3]) &&
			isConcave(comb[2], comb[3], comb[4]) &&
			isConcave(comb[3], comb[4], comb[5]) &&
			isConcave(comb[4], comb[5], comb[6]) &&
			isConcave(comb[5], comb[6], comb[7]) &&
			isConcave(comb[6], comb[7], comb[0]) &&
			isConcave(comb[7], comb[0], comb[1]) {
			return 1
		}
		return 0
	}

	res := 0
	for i := 0; i < maxLen; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		comb = append(comb, arr[i])
		res += combination(arr, visited, comb, depth+1)
		comb = comb[:len(comb)-1]
		visited[i] = false
	}
	return res
}

func isConcave(a, b, c int) bool {
	const cos45 = 0.7071

	var (
		//ax = float64(0)
		ay = float64(a)
		bx = float64(b) * cos45
		by = float64(b) * cos45
		cx = float64(c)
		//cy = float64(0)
	)

	//a0 := (by - ay) / (bx - ax)
	//a1 := (cy - ay) / (cx - ax)
	a0 := (by - ay) / bx
	a1 := -ay / cx

	return a0 > a1
}
