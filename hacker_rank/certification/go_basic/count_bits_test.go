package go_basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	for _, tt := range []struct {
		n    uint32
		want int32
	}{
		{127, 7},
		{10, 2},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := countBits(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
