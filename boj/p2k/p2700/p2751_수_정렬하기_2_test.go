package p2700_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
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
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p2700.Solve2751(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
