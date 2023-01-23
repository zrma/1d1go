package p10900

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10950(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10950")

	const (
		give = `5
1 1
2 3
3 4
9 8
5 2`
		want = `2
5
7
17
7
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve10950(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
