package p2700_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
	"1d1go/utils"
)

func TestSolve2739(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2739")

	const (
		s    = "2"
		want = `2 * 1 = 2
2 * 2 = 4
2 * 3 = 6
2 * 4 = 8
2 * 5 = 10
2 * 6 = 12
2 * 7 = 14
2 * 8 = 16
2 * 9 = 18
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p2700.Solve2739(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
