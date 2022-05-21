package p10100_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10100"
	"1d1go/utils"
)

func TestSolve10171(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10171")

	want := []string{
		`\    /\`,
		` )  ( ')`,
		`(  /  )`,
		` \(__)|`,
	}
	got, err := utils.GetPrinted(func() {
		p10100.Solve10171()
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
