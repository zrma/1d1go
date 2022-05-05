package p2500

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSolve2557(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2557")

	want := []string{"Hello World!"}
	got, err := utils.GetPrinted(func() {
		Solve2557()
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
