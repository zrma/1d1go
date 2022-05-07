package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedian(t *testing.T) {
	t.Log("https://www.hackerrank.com/test/eoipgdk427n/questions/a8taf02a12a")

	for _, tt := range []struct {
		arr  []int32
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
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			got := findMedian(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
