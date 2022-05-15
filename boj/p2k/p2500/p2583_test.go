package p2500_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils/mocks"
)

func TestSolve2583(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2583")

	for i, tt := range []struct {
		row, col int
		rects    []p2500.Rect
		want     []int
	}{
		{
			row: 5, col: 7,
			rects: []p2500.Rect{
				{0, 2, 4, 4},
				{1, 1, 2, 5},
				{4, 0, 6, 2},
			},
			want: []int{1, 7, 13},
		},
		{
			row: 100, col: 100,
			want: []int{10000},
		},
		{
			row: 100, col: 100,
			rects: []p2500.Rect{
				{0, 0, 1, 1},
			},
			want: []int{9999},
		},
		{
			row: 100, col: 100,
			rects: []p2500.Rect{
				{1, 0, 2, 100},
			},
			want: []int{100, 9800},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()

			scanner := mocks.NewMockScanner(ctr)
			writer := mocks.NewMockWriter(ctr)

			row := strconv.Itoa(tt.row)
			col := strconv.Itoa(tt.col)
			k := strconv.Itoa(len(tt.rects))

			scanner.EXPECT().Scan().Return(true).Times(3)
			scanner.EXPECT().Text().Return(row)
			scanner.EXPECT().Text().Return(col)
			scanner.EXPECT().Text().Return(k)

			for _, rect := range tt.rects {
				x1 := strconv.Itoa(rect[0])
				y1 := strconv.Itoa(rect[1])
				x2 := strconv.Itoa(rect[2])
				y2 := strconv.Itoa(rect[3])

				scanner.EXPECT().Scan().Return(true).Times(4)
				scanner.EXPECT().Text().Return(x1)
				scanner.EXPECT().Text().Return(y1)
				scanner.EXPECT().Text().Return(x2)
				scanner.EXPECT().Text().Return(y2)
			}

			wantLen := []byte(strconv.Itoa(len(tt.want)) + "\n")
			writer.EXPECT().Write(wantLen).Return(len(wantLen), nil)

			for _, want := range tt.want {
				wantStr := []byte(strconv.Itoa(want) + " ")
				writer.EXPECT().Write(wantStr).Return(len(wantStr), nil)
			}

			p2500.Solve2583(scanner, writer)
		})
	}
}
