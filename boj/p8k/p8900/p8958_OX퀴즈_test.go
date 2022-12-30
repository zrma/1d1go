package p8900_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p8k/p8900"
)

func TestSolve8958(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/8958")

	//goland:noinspection SpellCheckingInspection
	const (
		give = `5
OOXXOXXOOO
OOXXOOXXOO
OXOXOXOXOXOXOX
OOOOOOOOOO
OOOOXOOOOXOOOOX`
		want = `10
9
7
55
30
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	p8900.Solve8958(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
