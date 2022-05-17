package p1000_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve1000(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1000")

	for i, tt := range []struct {
		s    string
		want int
	}{
		{"1 2", 1 + 2},
		{"1 2", 3},
		{"1 9", 1 + 9},
		{"1 9", 10},
		{"9 1", 9 + 1},
		{"9 1", 10},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctrl)

			want := []byte(strconv.Itoa(tt.want) + "\n")
			writer.EXPECT().Write(want).Return(len(want), nil)

			p1000.Solve1000(scanner, writer)
		})
	}
}
