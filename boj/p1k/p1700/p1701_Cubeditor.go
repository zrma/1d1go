package p1700

import (
	"fmt"
	"io"
)

func Solve1701(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	max := 0

	for i := 0; i < len(s); i++ {
		sub := s[i:]
		m := getMax(sub)
		if max < m {
			max = m
		}
	}

	_, _ = fmt.Fprint(writer, max)
}

func getMax(s string) int {
	max := 0
	kmp := make([]int, len(s))
	j := 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = kmp[j-1]
		}

		if s[i] == s[j] {
			j++
			kmp[i] = j
			if max < j {
				max = j
			}
		}
	}

	return max
}
