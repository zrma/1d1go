package p11000_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11000"
)

func TestSolve11021(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11021")

	const (
		give = `5
1 1
2 3
3 4
9 8
5 2`
		want = `Case #1: 2
Case #2: 5
Case #3: 7
Case #4: 17
Case #5: 7
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	p11000.Solve11021(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
