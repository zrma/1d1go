package p10400

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSolve10430(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10430")

	want := []string{
		"1", "1", "0", "0",
	}
	got, err := utils.GetPrinted(func() {
		Solve10430(5, 8, 4)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}