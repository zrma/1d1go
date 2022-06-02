package p10700_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10700"
	"1d1go/utils"
)

func TestSolve10718(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10718")

	//goland:noinspection SpellCheckingInspection
	const want = `강한친구 대한육군
강한친구 대한육군
`
	writer := utils.NewStringWriter()
	p10700.Solve10718(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
