package p11000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11000"
	"1d1go/utils"
)

func TestSolve11021(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11021")

	want := []string{
		"Case #1: 2",
		"Case #2: 5",
		"Case #3: 7",
		"Case #4: 17",
		"Case #5: 7",
	}
	arr2D := [][2]int{
		{1, 1},
		{2, 3},
		{3, 4},
		{9, 8},
		{5, 2},
	}
	got, err := utils.GetPrinted(func() {
		p11000.Solve11021(arr2D)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
