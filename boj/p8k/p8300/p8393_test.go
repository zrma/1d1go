package p8300_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p8k/p8300"
)

func TestSolve8393(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/8393")

	for _, tt := range []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 3},
		{3, 6},
		{10, 55},
		{100, 5050},
		{1000, 500500},
		{10000, 50005000},
	} {
		t.Run(fmt.Sprintf("loop/%d", tt.n), func(t *testing.T) {
			got := p8300.Solve8393(tt.n)
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("arithmetic progression formula/%d", tt.n), func(t *testing.T) {
			got := p8300.Solve8393AP(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
