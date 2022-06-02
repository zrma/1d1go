package p2500_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
)

func TestSolve2588(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2588")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const s = `472
385`
	const want = `2360
3776
1416
181720
`
	scanner := utils.NewStringScanner(s)
	writer := utils.NewStringWriter()

	p2500.Solve2588(scanner, writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := writer.String()
	assert.Equal(t, want, got)
}
