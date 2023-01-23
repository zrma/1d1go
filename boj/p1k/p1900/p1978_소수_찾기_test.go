package p1900

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1978(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1978")

	const (
		give = `6
1 3 5 7 8 9`
		want = "3"
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve1978(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
