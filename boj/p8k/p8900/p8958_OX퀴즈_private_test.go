package p8900

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve8958(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want int
	}{
		{"OOXXOXXOOO", 10},
		{"OOXXOOXXOO", 9},
		{"OXOXOXOXOXOXOX", 7},
		{"OOOOOOOOOO", 55},
		{"OOOOXOOOOXOOOOX", 30},
	} {
		t.Run(tt.s, func(t *testing.T) {
			got := acc(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
