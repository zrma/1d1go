package p2500_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
)

func TestSolve2557(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2557")

	const want = `Hello World!`
	writer := utils.NewStringWriter()

	p2500.Solve2557(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
