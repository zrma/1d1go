package p14800

import (
	"fmt"
	"strconv"
)

func Solve14889(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scores := make([][]int, n)
	for i := range scores {
		scores[i] = make([]int, n)
		for j := range scores[i] {
			scanner.Scan()
			scores[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	const (
		max = 10 * 100
		min = 1 * 10
	)
	helper := startLinkHelper{
		n:         n,
		scores:    scores,
		isInTeam0: make([]bool, n),
		team0:     make([]int, n/2),
		team1:     make([]int, n/2),
		minDiff:   max - min,
	}
	helper.startLink(0, 0)

	res := helper.minDiff
	_, _ = fmt.Fprint(writer, res)
}

type startLinkHelper struct {
	n         int
	scores    [][]int
	isInTeam0 []bool
	team0     []int
	team1     []int
	minDiff   int
}

func (h *startLinkHelper) startLink(offset, count int) {
	if count == h.n/2 {
		h.calcDiff()
		return
	}

	for i := offset + 1; i < h.n; i++ {
		if h.isInTeam0[i] {
			continue
		}
		// 대상을 team0 에 추가해 봄
		h.isInTeam0[i] = true
		h.startLink(i, count+1)
		h.isInTeam0[i] = false
	}
}

func (h *startLinkHelper) calcDiff() {
	team0, team1 := 0, 0
	for i, b := range h.isInTeam0 {
		if b {
			h.team0[team0] = i
			team0++
		} else {
			h.team1[team1] = i
			team1++
		}
	}
	sum0 := h.sumScores(h.team0)
	sum1 := h.sumScores(h.team1)
	diff := sum0 - sum1
	if diff < 0 {
		diff = -diff
	}
	if diff < h.minDiff {
		h.minDiff = diff
	}
}

func (h *startLinkHelper) sumScores(team []int) int {
	sum := 0
	for i := 0; i < len(team); i++ {
		for j := i + 1; j < len(team); j++ {
			sum += h.scores[team[i]][team[j]]
			sum += h.scores[team[j]][team[i]]
		}
	}
	return sum
}
