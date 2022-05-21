package p25000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p25k/p25000"
	"1d1go/utils"
)

func TestSolve25083(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/25083")

	want := []string{
		"         ,r'\"7",
		"r`-_   ,'  ,/",
		" \\. \". L_r'",
		"   `~\\/",
		"      |",
		"      |",
	}
	got, err := utils.GetPrinted(func() {
		p25000.Solve25083()
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
