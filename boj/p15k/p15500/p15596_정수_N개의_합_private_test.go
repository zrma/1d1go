package p15500

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	for i, tt := range []struct {
		give []int
		want int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			15,
		},
		{
			[]int{5, 4, 3, 2},
			14,
		},
		{
			[]int{3, 6, 9},
			18,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := sum(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
