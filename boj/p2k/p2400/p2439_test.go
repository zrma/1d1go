package p2400_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2400"
	"1d1go/utils"
)

func TestSolve2439(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2439")

	want := []string{
		"      *",
		"     **",
		"    ***",
		"   ****",
		"  *****",
		" ******",
		"*******",
	}
	got, err := utils.GetPrinted(func() {
		p2400.Solve2439(7)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
