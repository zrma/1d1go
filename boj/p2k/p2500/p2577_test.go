package p2500_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve2577(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2577")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const s = `150
266
427`
	scanner := utils.NewStringScanner(s)
	writer := mocks.NewMockWriter(ctrl)

	wants := []int{3, 1, 0, 2, 0, 0, 0, 2, 0, 0}

	for _, want := range wants {
		b := []byte(strconv.Itoa(want) + "\n")
		writer.EXPECT().Write(b).Return(len(b), nil)
	}

	p2500.Solve2577(scanner, writer)
}
