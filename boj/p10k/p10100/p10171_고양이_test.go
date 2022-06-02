package p10100_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10100"
	"1d1go/utils"
)

func TestSolve10171(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10171")

	const want = `\    /\
 )  ( ')
(  /  )
 \(__)|`

	writer := utils.NewStringWriter()

	p10100.Solve10171(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
