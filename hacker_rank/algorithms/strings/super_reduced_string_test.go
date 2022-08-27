package strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperReducedString(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/reduced-string/problem")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{"aaabccddd", "abd"},
		{"aa", "Empty String"},
		{"baab", "Empty String"},
		{"abab", "abab"},
		{"abccba", "Empty String"},
		{"", "Empty String"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := superReducedString(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
