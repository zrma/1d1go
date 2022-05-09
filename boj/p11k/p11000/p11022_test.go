package p11000

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSolve11022(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11022")

	want := []string{
		"Case #1: 1 + 1 = 2",
		"Case #2: 2 + 3 = 5",
		"Case #3: 3 + 4 = 7",
		"Case #4: 9 + 8 = 17",
		"Case #5: 5 + 2 = 7",
	}
	arr2D := [][2]int{
		{1, 1},
		{2, 3},
		{3, 4},
		{9, 8},
		{5, 2},
	}
	got, err := utils.GetPrinted(func() {
		Solve11022(arr2D)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
