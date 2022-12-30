package p25000_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p25k/p25000"
)

func TestSolve25083(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/25083")

	const (
		want = "         ,r'\"7\n" +
			"r`-_   ,'  ,/\n" +
			" \\. \". L_r'\n" +
			"   `~\\/\n" +
			"      |\n" +
			"      |"
	)

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)

	p25000.Solve25083(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
