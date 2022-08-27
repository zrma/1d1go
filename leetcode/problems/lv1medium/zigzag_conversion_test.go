package lv1medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Log("https://leetcode.com/problems/zigzag-conversion/")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		s    string
		n    int
		want string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"AB", 1, "AB"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			assert.NotPanics(t, func() {
				got := convert(tt.s, tt.n)
				assert.Equal(t, tt.want, got)
			})
		})
	}
}
