package p1000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
)

func TestSolve1010(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1010")

	const (
		s = `3
2 2
1 5
13 29`
		want = `1
5
67863915
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p1000.Solve1010(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
