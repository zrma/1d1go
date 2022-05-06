package p25000

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
		Solve25083()
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
