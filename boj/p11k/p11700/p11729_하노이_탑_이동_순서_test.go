package p11700

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11729(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11729")

	const (
		give = "3"
		want = `7
1 3
1 2
3 2
1 3
2 1
2 3
1 3
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve11729(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
