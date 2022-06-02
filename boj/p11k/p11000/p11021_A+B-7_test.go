package p11000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11000"
	"1d1go/utils"
)

func TestSolve11021(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11021")

	const (
		s = `5
1 1
2 3
3 4
9 8
5 2`
		want = `Case #1: 2
Case #2: 5
Case #3: 7
Case #4: 17
Case #5: 7
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p11000.Solve11021(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
