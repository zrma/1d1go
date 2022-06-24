package integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncSubSeqLen(t *testing.T) {
	for i, tt := range []struct {
		give []int
		want int
	}{
		{},
		{
			[]int{},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{10, 20, 10, 30, 20, 50},
			4,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
		},
		{
			[]int{1, 5, 3, 4, 5, 6, 7, 4, 9, 10},
			8,
		},
		{
			[]int{1, 2, 3, 1, 2, 3, 4, 1, 2, 3},
			4,
		},
	} {
		t.Run(fmt.Sprintf("%d/O(n^2)", i), func(t *testing.T) {
			got := LongestIncSubSeqLenSquare(tt.give)
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/O(log n)", i), func(t *testing.T) {
			got := LongestIncSubSeqLenLog(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
