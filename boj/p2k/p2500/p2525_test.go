package p2500_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve2525(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2525")

	for i, tt := range []struct {
		s                    string
		wantHour, wantMinute int
	}{
		{
			`14 30
20`,
			14, 50,
		},
		{
			`17 40
80`,
			19, 0,
		},
		{
			`23 48
25`,
			0, 13,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctrl)

			want := fmt.Sprintln(tt.wantHour, tt.wantMinute)
			writer.EXPECT().Write([]byte(want)).Return(len(want), nil)

			p2500.Solve2525(scanner, writer)
		})
	}
}
