package p10100

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcSeatNum(t *testing.T) {
	for i, tt := range []struct {
		col, row, k  int
		wantOk       bool
		wantX, wantY int
	}{
		{3, 3, 8, true, 2, 1},
		{7, 6, 1, true, 1, 1},
		{7, 6, 6, true, 1, 6},
		{7, 6, 7, true, 2, 6},
		{7, 6, 11, true, 6, 6},
		{7, 6, 12, true, 7, 6},
		{7, 6, 13, true, 7, 5},
		{7, 6, 17, true, 7, 1},
		{7, 6, 87, false, 0, 0},
		{100, 100, 3000, true, 9, 64},
		{10000, 10000, 99999999, true, 5001, 5001},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			gotX, gotY, gotOk := calcSeatNumber(tt.col, tt.row, tt.k)
			assert.Equal(t, tt.wantOk, gotOk)
			if gotOk {
				assert.Equal(t, tt.wantX, gotX)
				assert.Equal(t, tt.wantY, gotY)
			}
		})
	}
}
