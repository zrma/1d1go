package p1000_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
)

func TestSolve1002(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1002")

	const (
		give = `7
0 0 13 40 0 37
0 0 3 0 7 4
1 1 1 1 1 5
1 1 1 1 1 1
0 0 3 0 1 1
0 0 1 0 3 1
0 0 2 0 1 1`
		want = `2
1
0
-1
0
0
1
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p1000.Solve1002(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
