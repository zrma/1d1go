package p1100

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1181(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1181")

	const (
		give = `13
but
i
wont
hesitate
no
more
no
more
it
cannot
wait
im
yours`
		want = `i
im
it
no
but
more
wait
wont
yours
cannot
hesitate
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve1181(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
