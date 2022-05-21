package p4600

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD(t *testing.T) {
	for _, tt := range []struct {
		n    int
		want int
	}{
		{1, 2},
		{33, 39},
		{39, 51},
		{51, 57},
		{57, 69},
		{69, 84},
		{84, 96},
		{96, 111},
		{111, 114},
		{114, 120},
		{120, 123},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := d(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
