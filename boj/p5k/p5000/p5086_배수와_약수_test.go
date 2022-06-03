package p5000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p5k/p5000"
	"1d1go/utils"
)

func TestSolve5086(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/5086")

	const (
		s = `8 16
32 4
17 5
0 0`
		want = `factor
multiple
neither
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p5000.Solve5086(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
