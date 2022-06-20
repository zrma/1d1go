package p8900

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve8958(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		give string
		want int
	}{
		{"OOXXOXXOOO", 10},
		{"OOXXOOXXOO", 9},
		{"OXOXOXOXOXOXOX", 7},
		{"OOOOOOOOOO", 55},
		{"OOOOXOOOOXOOOOX", 30},
	} {
		t.Run(tt.give, func(t *testing.T) {
			got := acc(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
