package p1000_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils/mocks"
)

func TestSolve1001(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1001")

	for i, tt := range []struct {
		a, b int
		want int
	}{
		{3, 2, 3 - 2},
		{3, 2, 1},
		{1, 1, 1 - 1},
		{1, 1, 0},
		{9, 1, 9 - 1},
		{9, 1, 8},
		{1, 9, 1 - 9},
		{1, 9, -8},
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

			want := []byte(strconv.Itoa(tt.want) + "\n")
			writer.EXPECT().Write(want).Return(len(want), nil)

			p1000.Solve1001(scanner, writer)
		})
	}
}
