package p10800_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
	"1d1go/utils"
)

func TestSolve10809(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10809")

	//goland:noinspection SpellCheckingInspection
	const (
		s    = `baekjoon`
		want = `1 0 -1 -1 2 -1 -1 -1 -1 4 3 -1 -1 7 5 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 `
	)
	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p10800.Solve10809(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
