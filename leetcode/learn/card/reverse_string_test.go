package card

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	t.Log("https://leetcode.com/explore/learn/card/recursion-i/250/principle-of-recursion/1440/")

	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		give []byte
		want []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		{[]byte{'A', 'B'}, []byte{'B', 'A'}},
	} {
		t.Run(fmt.Sprintf("%d/for", i), func(t *testing.T) {
			give := make([]byte, len(tt.give))
			copy(give, tt.give)

			reverseString(give)
			assert.Equal(t, tt.want, give)
		})

		t.Run(fmt.Sprintf("%d/recur", i), func(t *testing.T) {
			give := make([]byte, len(tt.give))
			copy(give, tt.give)

			reverseStringRecur(give)
			assert.Equal(t, tt.want, give)
		})
	}
}
