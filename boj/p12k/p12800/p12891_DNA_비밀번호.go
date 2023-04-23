package p12800

import (
	"fmt"
	"io"
)

func Solve12891(reader io.Reader, writer io.Writer) {
	var s, p int
	_, _ = fmt.Fscanln(reader, &s, &p)

	var dna string
	_, _ = fmt.Fscanln(reader, &dna)

	const (
		indexA = 0  // 'A' - 'A'
		indexC = 2  // 'C' - 'A'
		indexG = 6  // 'G' - 'A'
		indexT = 19 // 'T' - 'A'
	)

	limits := [26]int{}
	_, _ = fmt.Fscan(reader, &limits[indexA])
	_, _ = fmt.Fscan(reader, &limits[indexC])
	_, _ = fmt.Fscan(reader, &limits[indexG])
	_, _ = fmt.Fscan(reader, &limits[indexT])

	counts := [26]int{}
	isSatisfied := func() bool {
		return counts[indexA] >= limits[indexA] &&
			counts[indexC] >= limits[indexC] &&
			counts[indexG] >= limits[indexG] &&
			counts[indexT] >= limits[indexT]
	}

	for i := 0; i < p; i++ {
		counts[dna[i]-'A']++
	}

	res := 0
	if isSatisfied() {
		res++
	}

	for i := p; i < s; i++ {
		counts[dna[i-p]-'A']--
		counts[dna[i]-'A']++
		if isSatisfied() {
			res++
		}
	}

	_, _ = fmt.Fprint(writer, res)
}
