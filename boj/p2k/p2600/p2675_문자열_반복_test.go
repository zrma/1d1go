package p2600_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2600"
)

func TestSolve2675(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2675")

	//goland:noinspection SpellCheckingInspection
	const (
		give = `3
3 ABC
5 /HTP
2 \$%*+-./`
		want = `AAABBBCCC
/////HHHHHTTTTTPPPPP
\\$$%%**++--..//
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	p2600.Solve2675(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
