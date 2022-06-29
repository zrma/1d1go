package p1100_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1100"
	"1d1go/utils"
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
	writer := utils.NewStringWriter()

	p1100.Solve1181(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
