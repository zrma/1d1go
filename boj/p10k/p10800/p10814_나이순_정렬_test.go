package p10800_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
)

func TestSolve10814(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10814")

	//goland:noinspection SpellCheckingInspection
	const (
		give = `3
21 Junkyu
21 Dohyun
20 Sunyoung`
		want = `20 Sunyoung
21 Junkyu
21 Dohyun
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p10800.Solve10814(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
