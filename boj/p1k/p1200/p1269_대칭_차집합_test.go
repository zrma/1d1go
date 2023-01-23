package p1200

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1269(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1269")

	const (
		give = `3 5
1 2 4
2 3 4 5 6`
		want = "4"
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve1269(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
