package p1300

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGroupWord(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want bool
	}{
		{"happy", true},
		{"new", true},
		{"year", true},
		{"aba", false},
		{"abab", false},
		{"abcabc", false},
		{"a", true},
		{"aa", true},
		{"aca", false},
		{"ba", true},
		{"bb", true},
	} {
		t.Run(tt.s, func(t *testing.T) {
			got := isGroupWord(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
