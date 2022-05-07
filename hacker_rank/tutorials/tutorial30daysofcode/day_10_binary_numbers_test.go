package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestBinaryNumbers(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-binary-numbers/problem")

	for _, tt := range []struct {
		n    int32
		want string
	}{
		{5, "1"},
		{13, "2"},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			want := []string{tt.want}
			got, err := utils.GetPrinted(func() {
				binaryNumbers(tt.n)
			})
			assert.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}
