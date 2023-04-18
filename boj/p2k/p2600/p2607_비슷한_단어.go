package p2600

import (
	"fmt"
	"io"
)

func Solve2607(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	words := make([]string, n)
	for i := range words {
		_, _ = fmt.Fscanln(reader, &words[i])
	}

	first := words[0]
	arr0 := [26]int{}
	for _, c := range first {
		arr0[c-'A']++
	}

	var res int
	for _, word := range words[1:] {
		arr1 := [26]int{}
		for _, c := range word {
			arr1[c-'A']++
		}
		if isSimilar(arr0, arr1) && abs(len(first)-len(word)) <= 1 {
			res++
		}
	}
	_, _ = fmt.Fprint(writer, res)
}

func isSimilar(arr0, arr1 [26]int) bool {
	var diff int
	for i := range arr0 {
		diff += abs(arr0[i] - arr1[i])
	}
	return diff <= 2
}
