package p10800_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
	"1d1go/utils"
)

func TestSolve10814(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10814")

	//goland:noinspection SpellCheckingInspection
	const (
		s = `3
21 Junkyu
21 Dohyun
20 Sunyoung`
		want = `20 Sunyoung
21 Junkyu
21 Dohyun
`
	)
	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p10800.Solve10814(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
