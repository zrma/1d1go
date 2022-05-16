package p1000

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmeticProgression(t *testing.T) {
	t.Run("under 100", func(t *testing.T) {
		for n := 1; n < 100; n++ {
			got := arithmeticProgression(n)
			assert.True(t, got)
		}
	})

	t.Run("100~110 false", func(t *testing.T) {
		for n := 100; n <= 110; n++ {
			got := arithmeticProgression(n)
			assert.False(t, got)
		}
	})

	for _, tt := range []struct {
		n    int
		want bool
	}{
		{1, true},
		{111, true},
		{123, true},
		{246, true},
		{963, true},
		{999, true},
		{1000, false},
		{1111, false},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := arithmeticProgression(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
