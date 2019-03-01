package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-linked-list/problem", func() {
	//noinspection SpellCheckingInspection
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleLinkedList()

			close(done)
		}()
	})
})

//noinspection SpellCheckingInspection
func ExampleLinkedList() {
	LinkedList([]int{2, 3, 4, 1})
	// Output:
	// 2
	// 3
	// 4
	// 1
}
