package integer_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils/integer"
)

func TestMin(t *testing.T) {
	for i, tt := range []struct {
		give    []int
		wantMin int
		wantMax int
	}{
		{give: []int{1, 3, 5}, wantMin: 1, wantMax: 5},
		{give: []int{3, 1, 5}, wantMin: 1, wantMax: 5},
		{give: []int{5, 3, 1}, wantMin: 1, wantMax: 5},
		{give: []int{1, 1, 1}, wantMin: 1, wantMax: 1},
		{give: []int{1, 1, 2}, wantMin: 1, wantMax: 2},
		{give: []int{9, 1, 1, 1, 1, 1, 1, 1}, wantMin: 1, wantMax: 9},
		{give: nil, wantMin: 0, wantMax: 0},
		{give: []int{}, wantMin: 0, wantMax: 0},
	} {
		t.Run(fmt.Sprintf("%d/Min", i), func(t *testing.T) {
			got := integer.Min(tt.give...)
			assert.Equal(t, tt.wantMin, got)
		})

		t.Run(fmt.Sprintf("%d/Max", i), func(t *testing.T) {
			got := integer.Max(tt.give...)
			assert.Equal(t, tt.wantMax, got)
		})
	}
}

func TestPow(t *testing.T) {
	for _, tt := range []struct {
		n, p, want int64
	}{
		{3, 5, int64(math.Pow(3, 5))},
		{0, 5, 0},
		{10, 0, 1},
	} {
		actual := integer.Pow(tt.n, tt.p)
		assert.Equal(t, actual, tt.want)
	}
}
