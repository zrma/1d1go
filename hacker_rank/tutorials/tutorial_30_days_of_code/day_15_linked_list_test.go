package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-linked-list/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleLinkedList()

			close(done)
		}()
	})
})

func ExampleLinkedList() {
	LinkedList([]int{2, 3, 4, 1})
	// Output:
	// 2
	// 3
	// 4
	// 1
}
