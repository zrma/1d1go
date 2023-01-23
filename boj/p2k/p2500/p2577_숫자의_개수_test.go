package p2500

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2577(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2577")
	const (
		give = `150
266
427`
		want = `3
1
0
2
0
0
0
2
0
0
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve2577(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
