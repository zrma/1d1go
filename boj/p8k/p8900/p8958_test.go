package p8900_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p8k/p8900"
	"1d1go/utils"
)

func TestSolve8958(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/8958")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//goland:noinspection SpellCheckingInspection
	const (
		s = `5
OOXXOXXOOO
OOXXOOXXOO
OXOXOXOXOXOXOX
OOOOOOOOOO
OOOOXOOOOXOOOOX`
		want = `10
9
7
55
30
`
	)
	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p8900.Solve8958(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}