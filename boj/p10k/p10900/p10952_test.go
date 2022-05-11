package p10900

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSolve10952(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10952")

	want := []string{
		"2",
		"5",
		"7",
		"17",
		"7",
	}
	got, err := utils.GetPrinted(func() {
		Solve10952([][2]int{
			{1, 1},
			{2, 3},
			{3, 4},
			{9, 8},
			{5, 2},
		})
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
