package p10800

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve10815(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
