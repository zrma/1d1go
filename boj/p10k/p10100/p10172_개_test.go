package p10100_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10100"
	"1d1go/utils"
)

func TestSolve10172(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10172")

	const want = "|\\_/|\n|q p|   /}\n( 0 )\"\"\"\\\n|\"^\"`    |\n||_/=\\\\__|"

	writer := utils.NewStringWriter()

	p10100.Solve10172(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
