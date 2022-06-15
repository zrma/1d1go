package p9400_test

import (
	"testing"

	"1d1go/boj/p9k/p9400"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve9461(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/9461")

	const (
		s = `11
1
2
3
4
5
6
7
8
9
10
12`
		want = `1
1
1
2
2
3
4
5
7
9
16
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p9400.Solve9461(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
