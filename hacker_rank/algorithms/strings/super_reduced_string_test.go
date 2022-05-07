package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperReducedString(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/reduced-string/problem")

	//noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want string
	}{
		{"aaabccddd", "abd"},
		{"aa", "Empty String"},
		{"baab", "Empty String"},
		{"abab", "abab"},
		{"abccba", "Empty String"},
		{"", "Empty String"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			got := superReducedString(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
