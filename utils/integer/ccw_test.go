package integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCCW(t *testing.T) {
	for i, tt := range []struct {
		give [3][2]int
		want int
	}{
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{1, 5},
			},
			1,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{5, 1},
			},
			-1,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{3, 3},
			},
			0,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{5, 5},
			},
			0,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{6, 6},
			},
			0,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{9, 9},
			},
			0,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{1, 3},
			},
			1,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{3, 1},
			},
			-1,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{1, 1},
			},
			0,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{5, 5},
			},
			0,
		},
		{
			[3][2]int{
				{1, 1},
				{5, 5},
				{6, 6},
			},
			0,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := CCW(tt.give[0][0], tt.give[0][1], tt.give[1][0], tt.give[1][1], tt.give[2][0], tt.give[2][1])
			assert.Equal(t, tt.want, got)
		})
	}
}
