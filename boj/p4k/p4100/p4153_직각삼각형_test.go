package p4100

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve4153(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4153")

	const (
		give = `6 8 10
25 52 60
5 12 13
0 0 0`
		want = `right
wrong
right
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve4153(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
