package p2500

import (
	"bufio"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve2562(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)

}
