package p11000

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11022(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11022")

	const (
		give = `5
1 1
2 3
3 4
9 8
5 2`
		want = `Case #1: 1 + 1 = 2
Case #2: 2 + 3 = 5
Case #3: 3 + 4 = 7
Case #4: 9 + 8 = 17
Case #5: 5 + 2 = 7
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve11022(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
