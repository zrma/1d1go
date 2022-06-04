package p4600_test

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p4k/p4600"
	"1d1go/utils"
)

//go:embed test_data/p4673_want.txt
var p4673want string

func TestSolve4673(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4673")

	writer := utils.NewStringWriter()
	p4600.Solve4673(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, p4673want, got)
}
