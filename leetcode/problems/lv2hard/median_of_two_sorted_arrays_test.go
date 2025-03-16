package lv2hard

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log("https://leetcode.com/problems/median-of-two-sorted-arrays/")

	for i, tt := range []struct {
		arr1, arr2 []int
		want       float64
	}{
		{
			arr1: []int{1, 3},
			arr2: []int{2},
			want: 2.0,
		},
		{
			arr1: []int{1, 2},
			arr2: []int{3, 4},
			want: 2.5,
		},
		{
			arr1: []int{1, 2},
			arr2: []int{100},
			want: 2.0,
		},
		{
			arr1: []int{1, 200},
			arr2: []int{10, 90},
			want: 50.0,
		},
		{
			arr1: []int{1, 10, 200},
			arr2: []int{90},
			want: 50.0,
		},
		{
			arr1: []int{1, 2, 2, 3},
			arr2: []int{100},
			want: 2.0,
		},
		{
			arr1: []int{1, 2, 2, 3, 3, 3},
			arr2: []int{100},
			want: 3.0,
		},
		{
			arr1: []int{2, 2, 3, 3, 3, 100},
			arr2: []int{1},
			want: 3.0,
		},
		{
			arr1: []int{2, 2, 2, 2, 2, 2},
			arr2: []int{2},
			want: 2.0,
		},
		{
			arr1: []int{2},
			arr2: []int{2, 2, 2, 2, 2, 2},
			want: 2.0,
		},
		{
			arr1: []int{2, 3, 4, 4, 4, 4, 4, 4, 4},
			arr2: []int{1, 5, 5, 5},
			want: 4,
		},
		{
			arr1: []int{3, 4},
			arr2: []int{1, 2},
			want: 2.5,
		},
		{
			arr1: []int{1},
			arr2: []int{1},
			want: 1,
		},
		{
			arr1: []int{1},
			arr2: []int{2, 3, 4},
			want: 2.5,
		},
	} {
		t.Run(fmt.Sprintf("%d by merge", i), func(t *testing.T) {
			got := findMedianSortedArraysWithMerge(tt.arr1, tt.arr2)
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d by index", i), func(t *testing.T) {
			got := findMedianSortedArrays(tt.arr1, tt.arr2)
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d by bin search", i), func(t *testing.T) {
			got := findMedianSortedArraysWithBinSearch(tt.arr1, tt.arr2)
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d by bin search solution\"", i), func(t *testing.T) {
			got := findMedianSortedArraysWithBinSearchSolution(tt.arr1, tt.arr2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetVal(t *testing.T) {
	for _, tt := range []struct {
		give int
		want int
	}{
		{-1, math.MinInt32},
		{0, 0},
		{4, 4},
		{5, math.MaxInt32},
	} {
		t.Run(fmt.Sprintf("%d", tt.give), func(t *testing.T) {
			nums := arr([]int{0, 1, 2, 3, 4})
			got := nums.getVal(tt.give)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkFindMedian(b *testing.B) {
	sizes := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			b.StopTimer()
			nums1, nums2 := prepareSortedNumsPair(size)
			b.StartTimer()

			b.Run("By merge", func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						findMedianSortedArraysWithMerge(nums1, nums2)
					}
				})
			})

			b.Run("By index", func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						findMedianSortedArrays(nums1, nums2)
					}
				})
			})

			b.Run("By bin search", func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						findMedianSortedArraysWithBinSearch(nums1, nums2)
					}
				})
			})

			b.Run("By bin search solution", func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						findMedianSortedArraysWithBinSearchSolution(nums1, nums2)
					}
				})
			})
		})
	}
}

func prepareSortedNums(cnt int) []int {
	nums := make([]int, cnt)
	for i := range nums {
		n, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err != nil {
			panic(err)
		}
		nums[i] = int(n.Int64())
	}
	sort.Ints(nums)
	return nums
}

func prepareSortedNumsPair(size int) ([]int, []int) {
	return prepareSortedNums(size), prepareSortedNums(size)
}
