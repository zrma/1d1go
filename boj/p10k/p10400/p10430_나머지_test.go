package p10400_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10400"
	"1d1go/utils"
)

func TestSolve10430(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10430")

	const (
		give = "5 8 4"
		want = `1
1
0
0
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p10400.Solve10430(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
