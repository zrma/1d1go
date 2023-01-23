package p10100

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10172(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10172")

	const want = "|\\_/|\n|q p|   /}\n( 0 )\"\"\"\\\n|\"^\"`    |\n||_/=\\\\__|"

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve10172(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
