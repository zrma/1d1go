package p2400_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2400"
	"1d1go/utils"
)

func TestSolve2439(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2439")

	const (
		s    = "5"
		want = `    *
   **
  ***
 ****
*****
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p2400.Solve2439(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
