package p9400_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p9k/p9400"
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
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p9400.Solve9461(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
