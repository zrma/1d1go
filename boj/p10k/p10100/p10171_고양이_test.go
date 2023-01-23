package p10100

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10171(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10171")

	const want = `\    /\
 )  ( ')
(  /  )
 \(__)|`

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve10171(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
