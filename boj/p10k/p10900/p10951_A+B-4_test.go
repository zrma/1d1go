package p10900_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils"
)

func TestSolve10951(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10951")

	const (
		give = `1 1
2 3
3 4
9 8
5 2`
		want = `2
5
7
17
7
`
	)
	reader := bufio.NewReader(strings.NewReader(give))
	writer := utils.NewStringWriter()

	p10900.Solve10951(reader, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}

func TestSolve10951_StopAbnormally(t *testing.T) {
	t.Run("first scan returns false", func(t *testing.T) {
		const give = ``
		reader := bufio.NewReader(strings.NewReader(give))
		writer := utils.NewStringWriter()

		p10900.Solve10951(reader, writer)

		err := writer.Flush()
		assert.NoError(t, err)

		got := writer.String()
		assert.Empty(t, got)
	})

	t.Run("second scan returns false", func(t *testing.T) {
		const give = `1`
		reader := bufio.NewReader(strings.NewReader(give))
		writer := utils.NewStringWriter()

		p10900.Solve10951(reader, writer)

		err := writer.Flush()
		assert.NoError(t, err)

		got := writer.String()
		assert.Empty(t, got)
	})
}
