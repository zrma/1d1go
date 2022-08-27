package strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{"aabbcd", "NO"},
		{"aabbccddeefghi", "NO"},
		{"abbccc", "NO"},
		{"aaaabbcc", "NO"},
		{"aaaaabc", "NO"},
		{"aaabbbc", "YES"},
		{"aaabb", "YES"},
		{"abcdefghhgfedecba", "YES"},
		{"aabbc", "YES"},
		{"aabbcc", "YES"},
		{"aabbccc", "YES"},
		{"aabbccd", "YES"},
		{"abbc", "YES"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isValid(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
