package p2500_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2588(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2588")

	const a, b = 472, 385
	want := []int{2360, 3776, 1416, 181720}

	io := newIOWithMock(t)
	io.EXPECT().
		Scan(gomock.Any()).
		Do(func(args ...any) {
			assert.Len(t, args, 2)
			pA, ok := args[0].(*int)
			assert.True(t, ok)
			pB, ok := args[1].(*int)
			assert.True(t, ok)
			*pA = a
			*pB = b
		}).
		Return(0, nil)
	for _, n := range want {
		io.EXPECT().
			Println(n).
			Return(0, nil)
	}

	p2500.Solve2588(io)
}
