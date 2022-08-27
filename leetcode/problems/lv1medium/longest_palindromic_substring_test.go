package lv1medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	t.Log("https://leetcode.com/problems/longest-palindromic-substring/")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"cbbc", "cbbc"},
		{"abcd", "a"},
		{"", ""},
		{"a", "a"},
		{"abcdefedcba", "abcdefedcba"},
		{"aabcdefedcba", "abcdefedcba"},
		{"abcdefedcbaa", "abcdefedcba"},
		{"aaabcdefedcba", "abcdefedcba"},
		{"abcdefedcbaaa", "abcdefedcba"},
		{"bb", "bb"},
		{"aba", "aba"},
		{"abb", "bb"},
		{"aab", "aa"},
		{"abbb", "bbb"},
		{"aabbb", "bbb"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := longestPalindrome(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
