package p10900_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve10998(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10998")

	for i, tt := range []struct {
		s    string
		want int
	}{
		{"1 1", 1 * 1},
		{"1 1", 1},
		{"1 2", 1 * 2},
		{"1 2", 2},
		{"1 9", 1 * 9},
		{"1 9", 9},
		{"9 1", 9 * 1},
		{"9 1", 9},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctrl)

			want := []byte(strconv.Itoa(tt.want) + "\n")
			writer.EXPECT().Write(want).Return(len(want), nil)

			p10900.Solve10998(scanner, writer)
		})
	}
}
