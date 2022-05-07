package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLonelyInteger(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/one-week-preparation-kit-lonely-integer/problem")

	for _, tt := range []struct {
		arr  []int32
		want int32
	}{
		{[]int32{1, 1, 2}, 2},
		{[]int32{0, 0, 1, 2, 1}, 2},
		{[]int32{1, 1, 2, 2, 3, 3, 4, 4, 5}, 5},
		{[]int32{1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 5}, 4},
		{[]int32{1, 2, 3, 4, 3, 2, 1}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tt.arr), func(t *testing.T) {
			got := LonelyInteger(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
