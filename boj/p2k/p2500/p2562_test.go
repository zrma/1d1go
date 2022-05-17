package p2500_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve2562(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2562")

	ctr := gomock.NewController(t)
	defer ctr.Finish()

	const s = `3
29
38
12
57
74
40
85
61`
	scanner := utils.NewStringScanner(s)
	writer := mocks.NewMockWriter(ctr)

	var (
		want0 = 85
		want1 = 8
	)
	b0 := []byte(strconv.Itoa(want0) + "\n")
	writer.EXPECT().Write(b0).Return(len(b0), nil)
	b1 := []byte(strconv.Itoa(want1) + "\n")
	writer.EXPECT().Write(b1).Return(len(b1), nil)

	p2500.Solve2562(scanner, writer)
}
