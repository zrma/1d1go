package p10100

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
		Solve10171()
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
