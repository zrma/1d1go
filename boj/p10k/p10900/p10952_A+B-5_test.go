package p10900_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils"
)

func TestSolve10952(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10952")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		s = `1 1
2 3
3 4
9 8
5 2
0 0`
		want = `2
5
7
17
7
`
	)
	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p10900.Solve10952(scanner, writer)

	got := writer.String()
	assert.Equal(t, want, got)
}

func TestSolve10952_StopAbnormally(t *testing.T) {
	t.Run("first scan returns false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		const s = ``
		scanner := utils.NewStringScanner(s)
		writer := utils.NewStringWriter()

		p10900.Solve10952(scanner, writer)

		got := writer.String()
		assert.Empty(t, got)
	})

	t.Run("second scan returns false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		const s = `1`
		scanner := utils.NewStringScanner(s)
		writer := utils.NewStringWriter()

		p10900.Solve10952(scanner, writer)

		got := writer.String()
		assert.Empty(t, got)
	})
}
