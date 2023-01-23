package p9400

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve9461(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9461")

	const (
		give = `11
1
2
3
4
5
6
7
8
9
10
12`
		want = `1
1
1
2
2
3
4
5
7
9
16
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve9461(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
