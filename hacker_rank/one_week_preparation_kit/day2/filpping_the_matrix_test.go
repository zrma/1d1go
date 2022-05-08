package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipMatrix(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/flipping-the-matrix/problem")

	for i, tt := range []struct {
		matrix [][]int32
		want   int32
	}{
		{
			[][]int32{
				{1, 2},
				{3, 4},
			},
			4,
		},
		{
			[][]int32{
				{112, 42, 83, 119},
				{56, 125, 56, 49},
				{15, 78, 101, 43},
				{62, 98, 114, 108},
			},
			414,
		},
		{
			[][]int32{
				{107, 54, 128, 15},
				{12, 75, 110, 138},
				{100, 96, 34, 85},
				{75, 15, 28, 112},
			},
			488,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := flipMatrix(tt.matrix)
			assert.Equal(t, tt.want, got)
		})
	}
}
