package p8900

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

	Solve8958(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}

func TestAcc(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		give string
		want int
	}{
		{"OOXXOXXOOO", 10},
		{"OOXXOOXXOO", 9},
		{"OXOXOXOXOXOXOX", 7},
		{"OOOOOOOOOO", 55},
		{"OOOOXOOOOXOOOOX", 30},
	} {
		t.Run(tt.give, func(t *testing.T) {
			got := acc(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
