package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConstruction(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/string-construction/problem")

	//noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want int32
	}{
		{"abcd", 4},
		{"abab", 2},
		{"abcabc", 3},
	} {
		t.Run(tt.s, func(t *testing.T) {
			got := stringConstruction(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
