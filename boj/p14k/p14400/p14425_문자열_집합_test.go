package p14400_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p14k/p14400"
	"1d1go/utils"
)

func TestSolve14425(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/14425")

	//goland:noinspection SpellCheckingInspection
	const (
		give = `5 11
baekjoononlinejudge
startlink
codeplus
sundaycoding
codingsh
baekjoon
codeplus
codeminus
startlink
starlink
sundaycoding
codingsh
codinghs
sondaycoding
startrink
icerink`
		want = "4"
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p14400.Solve14425(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
