package p1900_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
	"1d1go/utils"
)

func TestSolve1934(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1934")

	const (
		s = `4
1 45000
6 10
13 17
44999 45000`
		want = `45000
30
221
2024955000
`
	)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p1900.Solve1934(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}

func TestSolve1934_Performance(t *testing.T) {
	s, err := utils.ReadStringFromFile("./test_data/p1934_give.txt")
	assert.NoError(t, err)

	want, err := utils.ReadStringFromFile("./test_data/p1934_want.txt")
	assert.NoError(t, err)

	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	assert.Eventually(t, func() bool {
		p1900.Solve1934(scanner, writer)
		return true
	}, time.Second, time.Millisecond*100, "시간 초과")

	err = writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
