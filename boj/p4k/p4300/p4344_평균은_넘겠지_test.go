package p4300

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve4344(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4344")

	const (
		give = `5
5 50 50 70 80 100
7 100 95 90 80 70 60 50
3 70 90 80
3 70 90 81
9 100 99 98 97 96 95 94 93 91`
		want = `40.000%
57.143%
33.333%
66.667%
55.556%
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	Solve4344(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
