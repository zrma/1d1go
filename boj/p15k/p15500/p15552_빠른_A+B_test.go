package p15500_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15500"
	"1d1go/utils"
)

func TestSolve15552(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15552")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		s = `5
1 1
12 34
5 500
40 60
1000 1000`
		want = `2
46
505
100
2000
`
	)
	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p15500.Solve15552(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
