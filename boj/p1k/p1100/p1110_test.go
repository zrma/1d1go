package p1100_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p1k/p1100"
)

func TestSolve1110(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1110")

	for _, tt := range []struct {
		n    int
		want int
	}{
		{26, 4},
		{55, 3},
		{1, 60},
		{0, 1},
		{71, 12},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			io := newIOWithMock(t)
			io.EXPECT().
				Scan(gomock.Any()).
				SetArg(0, tt.n).
				Return(0, nil)
			io.EXPECT().
				Println(tt.want).
				Return(0, nil)

			p1100.Solve1110(io)
		})
	}
}
