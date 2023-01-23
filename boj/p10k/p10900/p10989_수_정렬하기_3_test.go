package p10900

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10989(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10989")

	const (
		give = `10
5
2
3
1
4
2
3
5
1
7`
		want = `1
1
2
2
3
3
4
5
5
7
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve10989(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
