package p10800

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSolve10869(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10869")

	want := []string{
		"10",
		"4",
		"21",
		"2",
		"1",
	}
	got, err := utils.GetPrinted(func() {
		Solve10869(7, 3)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
