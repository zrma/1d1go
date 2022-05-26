package p10250_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10250"
	"1d1go/utils"
)

func TestSolve10250(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10250")

	const (
		s = `4
6 12 10
1 2 2
2 1 2
30 50 72`
		want = `402
102
201
1203
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p10250.Solve10250(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
