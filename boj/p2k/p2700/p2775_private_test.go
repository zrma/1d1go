package p2700

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcPeopleCounts(t *testing.T) {
	//       1호      2호      3호      4호      5호
	// 5 |    1        7       28       84      210
	// 4 |    1        6       21       56      126
	// 3 |    1        5       15       35       70
	// 2 |    1        4       10       20       35
	// 1 |    1        3        6       10       15
	// 0 |    1        2        3        4        5

	for i, tt := range []struct {
		row, col int
		want     int
	}{
		{1, 1, 1},
		{1, 2, 3},
		{1, 3, 6},
		{1, 4, 10},
		{1, 5, 15},
		{2, 1, 1},
		{2, 2, 4},
		{2, 3, 10},
		{2, 4, 20},
		{2, 5, 35},
		{3, 1, 1},
		{3, 2, 5},
		{3, 3, 15},
		{3, 4, 35},
		{3, 5, 70},
		{4, 1, 1},
		{4, 2, 6},
		{4, 3, 21},
		{4, 4, 56},
		{4, 5, 126},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := calcPeopleCounts(tt.row, tt.col)
			assert.Equal(t, tt.want, got)
		})
	}
}
