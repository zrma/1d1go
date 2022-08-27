package strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConstruction(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/string-construction/problem")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give string
		want int32
	}{
		{"abcd", 4},
		{"abab", 2},
		{"abcabc", 3},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := stringConstruction(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
