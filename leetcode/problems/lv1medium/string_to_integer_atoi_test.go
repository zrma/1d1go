package lv1medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoI(t *testing.T) {
	t.Log("https://leetcode.com/problems/string-to-integer-atoi/")

	for _, tt := range []struct {
		description string
		given       string
		want        int
	}{
		{"integer", "42", 42},
		{"blank + negative", "   -42", -42},
		{"integer + character", "4193 with words", 4193},
		{"character + integer", "words and 987", 0},
		{"negative", "-91283472332", -2147483648},
		{"empty", "", 0},
		{"decimal", "3.14159", 3},
		{"blank + negative + character", "  -0012a42", -12},
		{"blank", "          ", 0},
		{"blank + sign", "          -", 0},
		{"+ explicit positive", "+1", 1},
		{"large number overflow", "9223372036854775808", 2147483647},
		{"large number overflow", "18446744073709551617", 2147483647},
		{"leading space 0 truncated", "  0000000000012345678", 12345678},
		{"zero padded negative numbers", "-000000000000001", -1},
	} {
		t.Run(tt.description, func(t *testing.T) {
			got := myAtoI(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
