package p10800_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
	"1d1go/utils"
)

func TestSolve10816(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10816")

	const (
		s = `10
6 3 2 10 10 10 -10 -10 7 3
8
10 9 -5 2 3 4 5 -10`
		want = "3 0 0 1 2 0 0 2 "
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p10800.Solve10816(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
