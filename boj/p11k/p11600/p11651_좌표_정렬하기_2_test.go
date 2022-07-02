package p11600_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11600"
	"1d1go/utils"
)

func TestSolve11651(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11651")

	const (
		give = `5
0 4
1 2
1 -1
2 2
3 3`
		want = `1 -1
1 2
2 2
3 3
0 4
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p11600.Solve11651(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
