package lv0easy

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Log("https://leetcode.com/problems/two-sum/")

	for i, tt := range []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			nums:   []int{3, 2, 4},
			target: 5,
			want:   []int{0, 1},
		},
		{
			nums:   []int{2, 3, 4},
			target: 6,
			want:   []int{0, 2},
		},
		{
			nums:   []int{2, 3, 4},
			target: 7,
			want:   []int{1, 2},
		},
	} {
		t.Run(fmt.Sprintf("%d/o(n^2)", i), func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})

		if sort.IntsAreSorted(tt.nums) {
			t.Run(fmt.Sprintf("%d/o(n)", i), func(t *testing.T) {
				got := twoSumSorted(tt.nums, tt.target)
				assert.Equal(t, tt.want, got)
			})
		}

		t.Run(fmt.Sprintf("%d/o(n)", i), func(t *testing.T) {
			got := twoSumHash(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTwoSumHash_Invalid(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 10
	got := twoSumHash(nums, target)
	assert.Empty(t, got)
}

func TestTwoSum2InputArrayIsSorted(t *testing.T) {
	t.Log("https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/")

	var (
		nums   = []int{2, 7, 11, 15}
		target = 9
		want   = []int{1, 2}
	)
	got := twoSum2InputArrayIsSorted(nums, target)
	assert.Equal(t, want, got)
}
