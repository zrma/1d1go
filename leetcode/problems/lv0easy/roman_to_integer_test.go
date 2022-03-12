package lv0easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInteger(t *testing.T) {
	t.Log("https://leetcode.com/problems/roman-to-integer/")

	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		given string
		want  int
	}{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
		{"IV", 4},
		{"IX", 9},
		{"XL", 40},
		{"XC", 90},
		{"CD", 400},
		{"CM", 900},
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MMMCMXCIX", 3999},
	} {
		t.Run(tt.given, func(t *testing.T) {
			got := romanToInt(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
