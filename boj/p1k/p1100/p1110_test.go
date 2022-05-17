package p1100_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p1k/p1100"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve1110(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1110")

	for _, tt := range []struct {
		s    string
		want int
	}{
		{"26", 4},
		{"55", 3},
		{"1", 60},
		{"0", 1},
		{"71", 12},
	} {
		t.Run(tt.s, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctrl)

			want := []byte(strconv.Itoa(tt.want) + "\n")
			writer.EXPECT().Write(want).Return(len(want), nil)

			p1100.Solve1110(scanner, writer)
		})
	}
}
