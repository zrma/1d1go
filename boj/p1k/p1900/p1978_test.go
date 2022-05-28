package p1900_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
	"1d1go/utils"
)

func TestSolve1978(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1978")

	const (
		s = `6
1 3 5 7 8 9`
		want = "3"
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p1900.Solve1978(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
