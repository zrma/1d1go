package p5000

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve5086(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5086")

	const (
		give = `8 16
32 4
17 5
0 0`
		want = `factor
multiple
neither
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve5086(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
