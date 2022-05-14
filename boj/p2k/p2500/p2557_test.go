package p2500_test

import (
	"testing"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2557(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2557")

	const want = "Hello World!"
	io := newIOWithMock(t)
	io.EXPECT().
		Println(want).
		Return(0, nil)

	p2500.Solve2557(io)
}
