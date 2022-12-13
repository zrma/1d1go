package p1200_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1200"
)

func TestSolve1259(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1259")

	const (
		give = `121
1231
12421
5
0`
		want = `yes
no
yes
yes
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p1200.Solve1259(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
