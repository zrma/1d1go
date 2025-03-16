package benchmarks

import (
	"fmt"
	"testing"
)

func benchmarkSliceMake(size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int64, size)
		for j := 0; j < size; j++ {
			arr[j] = int64(j)
		}
		_ = arr
	}
}

func benchmarkSliceAppend(size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int64, 0, size)
		for j := 0; j < size; j++ {
			arr = append(arr, int64(j))
		}
		_ = arr
	}
}

func BenchmarkSliceConstruction(b *testing.B) {
	sizes := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size=%d/make", size), func(b *testing.B) {
			benchmarkSliceMake(size, b)
		})
		b.Run(fmt.Sprintf("size=%d/append", size), func(b *testing.B) {
			benchmarkSliceAppend(size, b)
		})
	}
}
