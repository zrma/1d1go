package p11600_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11600"
)

func TestSolve11659(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11659")

	const (
		give = `5 3
5 4 3 2 1
1 3
2 4
5 5`
		want = `12
9
1
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	p11600.Solve11659(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
