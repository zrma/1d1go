package p2500_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2562(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2562")

	ctr := gomock.NewController(t)
	defer ctr.Finish()

	const (
		give = `3
29
38
12
57
74
40
85
61`
		want = `85
8
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)

	p2500.Solve2562(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)

}
