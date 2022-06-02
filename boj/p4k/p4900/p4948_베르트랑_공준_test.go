package p4900_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p4k/p4900"
	"1d1go/utils"
)

func TestSolve4948(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4948")

	const (
		s = `1
10
13
100
1000
10000
100000
0`
		want = `1
4
3
21
135
1033
8392
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p4900.Solve4948(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
