package p1900_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1900"
	"1d1go/utils"
)

func TestSolve1929(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1929")

	const (
		give = "3 16"
		want = `3
5
7
11
13
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p1900.Solve1929(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
