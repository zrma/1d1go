package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedian(t *testing.T) {
	t.Log("https://www.hackerrank.com/test/eoipgdk427n/questions/a8taf02a12a")

	for i, tt := range []struct {
		give []int32
		want int32
	}{
		{
			[]int32{0, 1, 2, 4, 6, 5, 3},
			3,
		},
		{
			[]int32{1, 2, 3},
			2,
		},
		{
			[]int32{3},
			3,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := findMedian(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}
