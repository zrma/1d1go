package p2500_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve2583(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2583")

	for i, tt := range []struct {
		s    string
		want []int
	}{
		{
			`5 7 3
0 2 4 4
1 1 2 5
4 0 6 2`,
			[]int{1, 7, 13},
		},
		{
			`100 100 0`,
			[]int{10000},
		},
		{
			`100 100 1
0 0 1 1`,
			[]int{9999},
		},
		{
			`100 100 1
1 0 2 100`,
			[]int{100, 9800},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctr := gomock.NewController(t)
			defer ctr.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctr)

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
