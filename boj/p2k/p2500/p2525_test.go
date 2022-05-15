package p2500_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils/mocks"
)

func TestSolve2525(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2525")

	for i, tt := range []struct {
		hour, minute, duration int
		wantHour, wantMinute   int
	}{
		{14, 30, 20, 14, 50},
		{17, 40, 80, 19, 0},
		{23, 48, 25, 0, 13},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := mocks.NewMockScanner(ctrl)
			writer := mocks.NewMockWriter(ctrl)

			hour := strconv.Itoa(tt.hour)
			minute := strconv.Itoa(tt.minute)
			duration := strconv.Itoa(tt.duration)

			scanner.EXPECT().Scan().Return(true).Times(3)
			scanner.EXPECT().Text().Return(hour)
			scanner.EXPECT().Text().Return(minute)
			scanner.EXPECT().Text().Return(duration)

			want := fmt.Sprintln(tt.wantHour, tt.wantMinute)
			writer.EXPECT().Write([]byte(want)).Return(len(want), nil)

			p2500.Solve2525(scanner, writer)
		})
	}
}
