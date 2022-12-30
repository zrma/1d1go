package p15500_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15500"
)

func TestSolve15596(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15596")

	const (
		give = `10
1 10 4 9 2 3 8 5 7 6`
		want = "55"
	)

	reader := bufio.NewReader(strings.NewReader(give))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	p15500.Solve15596(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
