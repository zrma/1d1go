package lv2hard

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
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

// goos: windows
// goarch: amd64
// pkg: 1d1go/leetcode/problems/lv2hard
// cpu: AMD Ryzen 7 3700X 8-Core Processor
// BenchmarkFindMedian1/By_merge-16                          93229951                14.35 ns/op
// BenchmarkFindMedian1/By_index-16                        1000000000               0.6978 ns/op
// BenchmarkFindMedian1/By_bin_search-16                    867582799                1.346 ns/op
// BenchmarkFindMedian1/By_bin_search_solution-16          1000000000               0.9351 ns/op
// BenchmarkFindMedian2/By_merge-16                          19999832                57.21 ns/op
// BenchmarkFindMedian2/By_index-16                         612280555                1.731 ns/op
// BenchmarkFindMedian2/By_bin_search-16                   1000000000                1.090 ns/op
// BenchmarkFindMedian2/By_bin_search_solution-16          1000000000               0.9409 ns/op
// BenchmarkFindMedian4/By_merge-16                          25208694                48.17 ns/op
// BenchmarkFindMedian4/By_index-16                         676985428                1.746 ns/op
// BenchmarkFindMedian4/By_bin_search-16                    482554848                2.585 ns/op
// BenchmarkFindMedian4/By_bin_search_solution-16           794007878                1.414 ns/op
// BenchmarkFindMedian8/By_merge-16                          16084474                70.82 ns/op
// BenchmarkFindMedian8/By_index-16                         646136136                1.593 ns/op
// BenchmarkFindMedian8/By_bin_search-16                   1000000000                1.055 ns/op
// BenchmarkFindMedian8/By_bin_search_solution-16          1000000000               0.8970 ns/op
// BenchmarkFindMedian16/By_merge-16                          8572186                140.2 ns/op
// BenchmarkFindMedian16/By_index-16                        335440082                3.452 ns/op
// BenchmarkFindMedian16/By_bin_search-16                   633400027                1.899 ns/op
// BenchmarkFindMedian16/By_bin_search_solution-16          681747295                1.792 ns/op
// BenchmarkFindMedian32/By_merge-16                          7271400                165.8 ns/op
// BenchmarkFindMedian32/By_index-16                        298076437                3.985 ns/op
// BenchmarkFindMedian32/By_bin_search-16                   673297515                1.839 ns/op
// BenchmarkFindMedian32/By_bin_search_solution-16          854694588                1.333 ns/op
// BenchmarkFindMedian64/By_merge-16                          4225348                278.0 ns/op
// BenchmarkFindMedian64/By_index-16                        161114265                6.852 ns/op
// BenchmarkFindMedian64/By_bin_search-16                   397667810                3.016 ns/op
// BenchmarkFindMedian64/By_bin_search_solution-16          693510532                1.804 ns/op
// BenchmarkFindMedian128/By_merge-16                         3281581                374.8 ns/op
// BenchmarkFindMedian128/By_index-16                       560939364                2.078 ns/op
// BenchmarkFindMedian128/By_bin_search-16                  683299036                1.685 ns/op
// BenchmarkFindMedian128/By_bin_search_solution-16        1000000000                1.029 ns/op
// BenchmarkFindMedian256/By_merge-16                         1952606                579.3 ns/op
// BenchmarkFindMedian256/By_index-16                       128957961                7.988 ns/op
// BenchmarkFindMedian256/By_bin_search-16                  917845316                1.333 ns/op
// BenchmarkFindMedian256/By_bin_search_solution-16         979635819                1.197 ns/op
// BenchmarkFindMedian512/By_merge-16                          857203                 1252 ns/op
// BenchmarkFindMedian512/By_index-16                       539345961                1.983 ns/op
// BenchmarkFindMedian512/By_bin_search-16                  411230214                2.848 ns/op
// BenchmarkFindMedian512/By_bin_search_solution-16         725652381                1.692 ns/op
// BenchmarkFindMedian1024/By_merge-16                         499945                 2264 ns/op
// BenchmarkFindMedian1024/By_index-16                      500163594                2.145 ns/op
// BenchmarkFindMedian1024/By_bin_search-16                 592210534                2.024 ns/op
// BenchmarkFindMedian1024/By_bin_search_solution-16       1000000000                1.057 ns/op
// BenchmarkFindMedian2048/By_merge-16                         272257                 4373 ns/op
// BenchmarkFindMedian2048/By_index-16                      958722748                1.206 ns/op
// BenchmarkFindMedian2048/By_bin_search-16                 519023497                2.345 ns/op
// BenchmarkFindMedian2048/By_bin_search_solution-16        780124360                1.655 ns/op
// BenchmarkFindMedian4096/By_merge-16                         108961                10165 ns/op
// BenchmarkFindMedian4096/By_index-16                      228664723                5.168 ns/op
// BenchmarkFindMedian4096/By_bin_search-16                 461453269                2.657 ns/op
// BenchmarkFindMedian4096/By_bin_search_solution-16       1000000000                1.177 ns/op
// BenchmarkFindMedian8192/By_merge-16                          69519                17916 ns/op
// BenchmarkFindMedian8192/By_index-16                     1000000000                1.093 ns/op
// BenchmarkFindMedian8192/By_bin_search-16                1000000000                1.080 ns/op
// BenchmarkFindMedian8192/By_bin_search_solution-16       1000000000               0.9189 ns/op
// PASS
func BenchmarkFindMedian1(b *testing.B)    { benchmarkFindMedianSortedArrays(b, 1) }
func BenchmarkFindMedian2(b *testing.B)    { benchmarkFindMedianSortedArrays(b, 2) }
func BenchmarkFindMedian4(b *testing.B)    { benchmarkFindMedianSortedArrays(b, 4) }
func BenchmarkFindMedian8(b *testing.B)    { benchmarkFindMedianSortedArrays(b, 8) }
func BenchmarkFindMedian16(b *testing.B)   { benchmarkFindMedianSortedArrays(b, 16) }
func BenchmarkFindMedian32(b *testing.B)   { benchmarkFindMedianSortedArrays(b, 32) }
func BenchmarkFindMedian64(b *testing.B)   { benchmarkFindMedianSortedArrays(b, 64) }
func BenchmarkFindMedian128(b *testing.B)  { benchmarkFindMedianSortedArrays(b, 128) }
func BenchmarkFindMedian256(b *testing.B)  { benchmarkFindMedianSortedArrays(b, 256) }
func BenchmarkFindMedian512(b *testing.B)  { benchmarkFindMedianSortedArrays(b, 512) }
func BenchmarkFindMedian1024(b *testing.B) { benchmarkFindMedianSortedArrays(b, 1024) }
func BenchmarkFindMedian2048(b *testing.B) { benchmarkFindMedianSortedArrays(b, 2048) }
func BenchmarkFindMedian4096(b *testing.B) { benchmarkFindMedianSortedArrays(b, 4096) }
func BenchmarkFindMedian8192(b *testing.B) { benchmarkFindMedianSortedArrays(b, 8192) }

func benchmarkFindMedianSortedArrays(b *testing.B, cnt1 int) {
	b.StopTimer()
	nums1 := make([]int, cnt1)
	for i := 0; i < cnt1; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(40))
		assert.NoError(b, err)
		nums1[i] = int(n.Int64())
	}

	n, err := rand.Int(rand.Reader, big.NewInt(40))
	assert.NoError(b, err)

	cnt2 := int(n.Int64()) + 1
	nums2 := make([]int, cnt2)
	for i := 0; i < cnt2; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(40))
		assert.NoError(b, err)
		nums2[i] = int(n.Int64())
	}
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
}
