package p11600

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11650(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11650")

	const (
		give = `5
3 4
1 1
1 -1
2 2
3 3`
		want = `1 -1
1 1
2 2
3 3
3 4
`
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve11650(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
