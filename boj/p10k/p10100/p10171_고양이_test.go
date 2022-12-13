package p10100_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10100"
)

func TestSolve10171(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10171")

	const want = `\    /\
 )  ( ')
(  /  )
 \(__)|`

	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p10100.Solve10171(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
