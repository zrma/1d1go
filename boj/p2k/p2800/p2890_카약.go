package p2800

import (
	"fmt"
	"io"
	"sort"
)

func Solve2890(reader io.Reader, writer io.Writer) {
	var r, c int
	_, _ = fmt.Fscan(reader, &r, &c)

	const boatCnt = 9
	boats := make([]string, 0, boatCnt)
	for i := 0; i < r; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		found := false
		s = s[1 : c-1]
		for _, c := range s {
			if c != '.' {
				found = true
				break
			}
		}
		if found {
			boats = append(boats, s)
		}
	}
	sort.Strings(boats)

	currRank := 0
	prevDist := -1
	ranks := make([]int, boatCnt)
	for _, boat := range boats {
		currDist := 0
		boatIdx := 0

		for j := len(boat) - 1; j >= 0; j-- {
			c := boat[j]

			if c == '.' {
				currDist++
				continue
			}

			boatIdx = int(c - '1')
			break
		}

		if prevDist != currDist {
			currRank++
		}

		ranks[boatIdx] = currRank
		prevDist = currDist
	}

	for _, rank := range ranks {
		_, _ = fmt.Fprintln(writer, rank)
	}
}
