package p2700

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2751(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2751")

	const (
		give = `5
5
2
3
4
1`
		want = `1
2
3
4
5
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve2751(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
