package p10250_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10250"
)

func TestSolve10250(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10250")

	const (
		give = `4
6 12 10
1 2 2
2 1 2
30 50 72`
		want = `402
102
201
1203
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	p10250.Solve10250(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
