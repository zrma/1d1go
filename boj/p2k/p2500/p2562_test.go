package p2500_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2562(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2562")

	io := newIOWithMock(t)

	in := []int{3, 29, 38, 12, 57, 74, 40, 85, 61}
	for _, n := range in {
		io.EXPECT().
			Scan(gomock.Any()).
			SetArg(0, n).
			Return(0, nil)
	}

	var (
		want0 = 85
		want1 = 8
	)
	io.EXPECT().
		Println(want0).
		Return(0, nil)
	io.EXPECT().
		Println(want1).
		Return(0, nil)

	p2500.Solve2562(io)
}
