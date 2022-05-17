package p1000_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve1008(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1008")

	for i, tt := range []struct {
		s    string
		want float64
	}{
		{"1 3", 1.0 / 3.0},
		{"1 3", 0.33333333333333333333333333333333},
		{"4 5", 4.0 / 5.0},
		{"4 5", 0.8},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctrl)

			want := []byte(fmt.Sprintln(tt.want))
			writer.EXPECT().Write(want).Return(len(want), nil)

			p1000.Solve1008(scanner, writer)
		})
	}
}
