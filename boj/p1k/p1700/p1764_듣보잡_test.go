package p1700_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1700"
	"1d1go/utils"
)

func TestSolve1764(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1764")

	//goland:noinspection SpellCheckingInspection
	const (
		give = `3 4
ohhenrie
charlie
baesangwook
obama
baesangwook
ohhenrie
clinton`
		want = `2
baesangwook
ohhenrie
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p1700.Solve1764(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
