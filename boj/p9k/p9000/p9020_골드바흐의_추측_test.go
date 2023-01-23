package p9000

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve9020(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9020")

	const (
		give = `5
1
4
8
10
16`
		want = `0 1
2 2
3 5
5 5
5 11
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve9020(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
