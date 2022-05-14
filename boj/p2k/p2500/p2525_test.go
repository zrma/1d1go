package p2500_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
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
			io := newIOWithMock(t)
			io.EXPECT().
				Scan(gomock.Any()).
				Do(func(args ...any) {
					assert.Len(t, args, 2)
					pA, ok := args[0].(*int)
					assert.True(t, ok)
					pB, ok := args[1].(*int)
					assert.True(t, ok)
					*pA = tt.hour
					*pB = tt.minute
				}).
				Return(0, nil)
			io.EXPECT().
				Scan(gomock.Any()).
				SetArg(0, tt.duration).
				Return(0, nil)

			io.EXPECT().
				Println(tt.wantHour, tt.wantMinute).
				Return(0, nil)

			p2500.Solve2525(io)
		})
	}
}
