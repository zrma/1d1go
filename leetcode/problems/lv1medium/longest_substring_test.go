package lv1medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log("https://leetcode.com/problems/longest-substring-without-repeating-characters/")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"abba", 2},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
