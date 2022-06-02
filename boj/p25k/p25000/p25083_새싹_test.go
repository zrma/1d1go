package p25000_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p25k/p25000"
	"1d1go/utils"
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

	writer := utils.NewStringWriter()

	p25000.Solve25083(writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
