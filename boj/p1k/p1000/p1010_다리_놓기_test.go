package p1000

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1010(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1010")

	const (
		give = `3
2 2
1 5
13 29`
		want = `1
5
67863915
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve1010(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
