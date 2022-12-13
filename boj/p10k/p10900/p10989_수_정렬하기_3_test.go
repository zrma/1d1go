package p10900_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10900"
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
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p10900.Solve10989(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
