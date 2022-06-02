package p10250

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcNearestRoom(t *testing.T) {
	for _, tt := range []struct {
		h, w, n int
		want    int
	}{
		{1, 1, 1, 101},
		{1, 2, 1, 101},
		{1, 2, 2, 102},
		{1, 10, 10, 110},
		{2, 1, 1, 101},
		{2, 1, 2, 201},
		{2, 2, 2, 201},
		{2, 2, 3, 102},
		{6, 12, 10, 402},
		{30, 50, 60, 3002},
		{30, 50, 61, 103},
		{30, 50, 72, 1203},
	} {
		t.Run(fmt.Sprintf("%d %d %d", tt.h, tt.w, tt.n), func(t *testing.T) {
			got := calcNearestRoom(tt.h, tt.w, tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
