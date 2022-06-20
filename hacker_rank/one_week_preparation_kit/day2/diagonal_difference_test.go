package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifference(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/diagonal-difference/problem")
	t.Log("https://www.hackerrank.com/challenges/one-week-preparation-kit-diagonal-difference/problem")

	for i, tt := range []struct {
		give [][]int32
		want int32
	}{
		{
			give: [][]int32{
				{11, 2, 4},
				{4, 5, 6},
				{10, 8, -12},
			},
			want: 15,
		},
		{
			give: [][]int32{
				{4, 2, 11},
				{4, 5, 6},
				{-12, 8, 10},
			},
			want: 15,
		},
		{},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := diagonalDifference(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
