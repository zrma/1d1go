package p2500_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2557(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2557")

	const want = `Hello World!`
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p2500.Solve2557(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
