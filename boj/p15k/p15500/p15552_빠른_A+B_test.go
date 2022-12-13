package p15500_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15500"
)

func TestSolve15552(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15552")

	const (
		give = `5
1 1
12 34
5 500
40 60
1000 1000`
		want = `2
46
505
100
2000
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p15500.Solve15552(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
