package p10800_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
	"1d1go/utils"
)

func TestSolve10869(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10869")

	const (
		s    = "7 3"
		want = `10
4
21
2
1
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p10800.Solve10869(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
