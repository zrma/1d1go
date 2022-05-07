package p2700

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSolve2739(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2739")

	want := []string{
		"2 * 1 = 2",
		"2 * 2 = 4",
		"2 * 3 = 6",
		"2 * 4 = 8",
		"2 * 5 = 10",
		"2 * 6 = 12",
		"2 * 7 = 14",
		"2 * 8 = 16",
		"2 * 9 = 18",
	}
	got, err := utils.GetPrinted(func() {
		Solve2739(2)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
