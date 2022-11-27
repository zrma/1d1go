package p1100

import (
	"fmt"
	"strconv"

	"1d1go/utils/integer"
)

func Solve1107(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	brokenBtn := make([]bool, numBtnCnt)
	for i := 0; i < m; i++ {
		var b int
		_, _ = fmt.Fscan(reader, &b)
		brokenBtn[b] = true
	}

	if n == curChannel {
		_, _ = fmt.Fprint(writer, 0)
		return
	}

	cnt0, _ := moveWithNumBtn(n, brokenBtn)
	cnt1 := moveWithPlusMinus(n)
	cnt2 := moveNearestSmaller(n, brokenBtn)

	minVal := integer.Min(cnt0, cnt1, cnt2)

	cnt3 := moveNearestGreater(n, brokenBtn, minVal)

	res := integer.Min(minVal, cnt3)
	_, _ = fmt.Fprint(writer, res)
}

const (
	curChannel = 100
	numBtnCnt  = 10
	maxVal     = 1_234_567_891
)

func moveWithNumBtn(target int, brokenBtn []bool) (int, bool) {
	if target == 0 {
		if brokenBtn[0] {
			return maxVal, false
		}
		return 1, true
	}

	cnt := 0
	for _, c := range strconv.Itoa(target) {
		btn := int(c - '0')
		if brokenBtn[btn] {
			return maxVal, false
		}
		cnt++
	}
	return cnt, true
}

func moveWithPlusMinus(target int) int {
	diff := target - curChannel
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func moveNearestSmaller(target int, brokenBtn []bool) int {
	cnt := 0
	for {
		if target <= 0 {
			return maxVal
		}
		cnt++
		target--
		if subCnt, ok := moveWithNumBtn(target, brokenBtn); ok {
			return cnt + subCnt
		}
	}
}

func moveNearestGreater(target int, brokenBtn []bool, maxCnt int) int {
	cnt := 0
	for cnt < maxCnt {
		target++
		cnt++
		if subCnt, ok := moveWithNumBtn(target, brokenBtn); ok {
			return cnt + subCnt
		}
	}
	return maxCnt
}
