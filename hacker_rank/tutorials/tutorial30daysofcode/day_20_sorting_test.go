package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestSorting(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-sorting/problem")

	for i, tt := range []struct {
		given []int32
		want  []string
	}{
		{
			given: []int32{1, 2, 3},
			want: []string{
				"Array is sorted in 0 swaps.",
				"First Element: 1",
				"Last Element: 3",
			},
		},
		{
			given: []int32{3, 2, 1},
			want: []string{
				"Array is sorted in 3 swaps.",
				"First Element: 1",
				"Last Element: 3",
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, err := utils.GetPrinted(func() {
				sorting(tt.given)
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
