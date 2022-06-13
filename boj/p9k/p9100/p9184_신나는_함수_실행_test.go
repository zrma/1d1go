package p9100_test

import (
	"bufio"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9100"
	"1d1go/utils"
)

func TestSolve9184(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9184")

	const (
		s = `21 0 0
1 1 1
2 2 2
10 4 6
50 50 50
-1 7 18
19 20 20
19 20 21
-1 -1 -1`
		want = `w(21, 0, 0) = 1
w(1, 1, 1) = 2
w(2, 2, 2) = 4
w(10, 4, 6) = 523
w(50, 50, 50) = 1048576
w(-1, 7, 18) = 1
w(19, 20, 20) = 524288
w(19, 20, 21) = 1048576
`
	)

	reader := bufio.NewReader(strings.NewReader(s))
	writer := utils.NewStringWriter()

	assert.Eventually(t, func() bool {
		p9100.Solve9184(reader, writer)

		err := writer.Flush()
		assert.NoError(t, err)

		got := writer.String()
		return assert.Equal(t, want, got)
	}, time.Second, time.Millisecond*100, "시간 초과")
}
