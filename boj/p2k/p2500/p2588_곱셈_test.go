package p2500

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2588(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2588")

	const give = `472
385`
	const want = `2360
3776
1416
181720
`
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve2588(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
