package p11700_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11700"
	"1d1go/utils"
)

func TestSolve11729(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11729")

	const (
		give = "3"
		want = `7
1 3
1 2
3 2
1 3
2 1
2 3
1 3
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p11700.Solve11729(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
