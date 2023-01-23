package p25000

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

	Solve25083(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.Equal(t, want, got)
}
