package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem")

	//noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want string
	}{
		{"aabbcd", "NO"},
		{"aabbccddeefghi", "NO"},
		{"abbccc", "NO"},
		{"aaaabbcc", "NO"},
		{"aaaaabc", "NO"},
		{"aabbc", "YES"},
		{"aabbcc", "YES"},
		{"abcdefghhgfedecba", "YES"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			got := isValid(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
