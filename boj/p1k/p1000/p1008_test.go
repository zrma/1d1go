package p1000_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils/mocks"
)

func TestSolve1008(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1008")

	for i, tt := range []struct {
		a, b int
		want float64
	}{
		{1, 3, 1.0 / 3.0},
		{1, 3, 0.33333333333333333333333333333333},
		{4, 5, 4.0 / 5.0},
		{4, 5, 0.8},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := mocks.NewMockScanner(ctrl)
			writer := mocks.NewMockWriter(ctrl)

			a := strconv.Itoa(tt.a)
			b := strconv.Itoa(tt.b)

			scanner.EXPECT().Scan().Return(true).Times(2)
			scanner.EXPECT().Text().Return(a)
			scanner.EXPECT().Text().Return(b)

			want := []byte(fmt.Sprintln(tt.want))
			writer.EXPECT().Write(want).Return(len(want), nil)

			p1000.Solve1008(scanner, writer)
		})
	}
}
