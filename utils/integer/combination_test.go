package integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationDP(t *testing.T) {
	for i, tt := range []struct {
		n, r int
		want int
	}{
		{1, 0, 1},
		{1, 1, 1},
		{1, 2, 0},
		{5, 2, 10},
		{6, 1, 6},
		{6, 2, 15},
		{6, 3, 20},
		{1000, 0, 1},
		{1000, 1, 1000},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			cache := BuildCache2DArr(tt.n+1, -1)
			got := CombinationDP(tt.n, tt.r, cache, func(v int) int { return v })
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCombinationDP_Modifier(t *testing.T) {
	const (
		n    = 1000
		r    = 500
		want = 5418
	)
	modifier := func(v int) int { return v % 10_007 }
	cache := BuildCache2DArr(n+1, -1)
	got := CombinationDP(n, r, cache, modifier)
	assert.Equal(t, want, got)
}

func TestBuildCache2DArr(t *testing.T) {
	for i, tt := range []struct {
		maxLen, initVal int
		want            [][]int
	}{
		{
			1, 2,
			[][]int{
				{2},
			},
		},
		{
			5, -1,
			[][]int{
				{-1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1},
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := BuildCache2DArr(tt.maxLen, tt.initVal)
			assert.Equal(t, tt.want, got)
		})
	}
}
