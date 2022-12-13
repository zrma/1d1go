package p10800_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
)

func TestSolve10815(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10815")

	const (
		give = `5
6 3 2 10 -10
8
10 9 -5 2 3 4 5 -10`
		want = "1 0 0 1 1 0 0 1 "
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p10800.Solve10815(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
