package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-sorting/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleSorting()

			close(done)
		}()
	})
})

func ExampleSorting() {
	Sorting([]int32{1, 2, 3})
	Sorting([]int32{3, 2, 1})
	// Output:
	// Array is sorted in 0 swaps.
	// First Element: 1
	// Last Element: 3
	// Array is sorted in 3 swaps.
	// First Element: 1
	// Last Element: 3
}
