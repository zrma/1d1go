package p1000_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve1065(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1065")

	for _, tt := range []struct {
		s    string
		want int
	}{
		{"110", 99},
		{"1", 1},
		{"210", 105},
		{"1000", 144},
		{"500", 119},
	} {
		t.Run(tt.s, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctrl)

			want := []byte(strconv.Itoa(tt.want) + "\n")
			writer.EXPECT().Write(want).Return(len(want), nil)

			p1000.Solve1065(scanner, writer)
		})
	}
}
