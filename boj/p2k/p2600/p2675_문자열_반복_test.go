package p2600_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2600"
	"1d1go/utils"
)

func TestSolve2675(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2675")

	//goland:noinspection SpellCheckingInspection
	const (
		s = `3
3 ABC
5 /HTP
2 \$%*+-./`
		want = `AAABBBCCC
/////HHHHHTTTTTPPPPP
\\$$%%**++--..//
`
	)
	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p2600.Solve2675(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
