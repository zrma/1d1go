package p2500_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils/mocks"
)

func TestSolve2577(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2577")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scanner := mocks.NewMockScanner(ctrl)
	writer := mocks.NewMockWriter(ctrl)

	arr := []int{150, 266, 427}
	wants := []int{3, 1, 0, 2, 0, 0, 0, 2, 0, 0}

	for _, n := range arr {
		s := strconv.Itoa(n)
		scanner.EXPECT().Scan().Return(true)
		scanner.EXPECT().Text().Return(s)
	}

	for _, want := range wants {
		s := []byte(strconv.Itoa(want) + "\n")
		writer.EXPECT().Write(s).Return(len(s), nil)
	}

	p2500.Solve2577(scanner, writer)
}
