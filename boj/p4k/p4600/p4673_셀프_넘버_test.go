package p4600_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p4k/p4600"
	"1d1go/utils"
)

func TestSolve4673(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4673")

	want, err := utils.ReadStringFromFile("./test_data/p4673_want.txt")
	assert.NoError(t, err)

	writer := utils.NewStringWriter()
	p4600.Solve4673(writer)

	err = writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
