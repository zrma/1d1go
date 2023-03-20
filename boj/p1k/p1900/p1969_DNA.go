package p1900

import (
	"fmt"
	"io"
)

func Solve1969(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscanln(reader, &n, &m)

	dna := make([]string, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscanln(reader, &dna[i])
	}

	hamming := 0
	for i := 0; i < m; i++ {
		cnt := [4]int{}
		for j := 0; j < n; j++ {
			switch dna[j][i] {
			case 'A':
				cnt[0]++
			case 'C':
				cnt[1]++
			case 'G':
				cnt[2]++
			case 'T':
				cnt[3]++
			}
		}

		max := 0
		for j := 1; j < 4; j++ {
			if cnt[max] < cnt[j] {
				max = j
			}
		}
		hamming += n - cnt[max]

		switch max {
		case 0:
			_, _ = fmt.Fprint(writer, "A")
		case 1:
			_, _ = fmt.Fprint(writer, "C")
		case 2:
			_, _ = fmt.Fprint(writer, "G")
		case 3:
			_, _ = fmt.Fprint(writer, "T")
		}
	}

	_, _ = fmt.Fprintln(writer)
	_, _ = fmt.Fprint(writer, hamming)
}
