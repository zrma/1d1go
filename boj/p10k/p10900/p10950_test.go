package p10900_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils"
)

func TestSolve10950(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10950")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		s = `5
1 1
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
	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p10900.Solve10950(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
