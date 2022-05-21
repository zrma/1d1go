package p1400_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1400"
)

func TestSolve1463(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1463")

	for _, test := range []struct {
		n    int
		want int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{6, 2},
		{7, 3},
		{8, 3},
		{9, 2},
		{10, 3},
		{11, 4},
		{12, 3},
	} {
		t.Run(fmt.Sprintf("%d", test.n), func(t *testing.T) {
			got := p1400.Solve1463(test.n)
			assert.Equal(t, test.want, got)
		})
	}
}
