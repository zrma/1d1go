package p1300

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGroupWord(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isGroupWord(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
