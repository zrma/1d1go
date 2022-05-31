package p1200_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1200"
	"1d1go/utils"
)

func TestSolve1269(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1269")

	const (
		s = `3 5
1 2 4
2 3 4 5 6`
		want = "4"
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p1200.Solve1269(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
