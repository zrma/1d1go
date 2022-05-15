package p2500_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils/mocks"
)

func TestSolve2557(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2557")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	writer := mocks.NewMockWriter(ctrl)

	want := []byte("Hello World!" + "\n")
	writer.EXPECT().Write(want).Return(len(want), nil)

	p2500.Solve2557(writer)
}
