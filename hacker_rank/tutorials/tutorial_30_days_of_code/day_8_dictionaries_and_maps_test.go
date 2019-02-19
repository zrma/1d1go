package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-dictionaries-and-maps/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleDictionariesAndMaps()

			close(done)
		}()
	})
})

func ExampleDictionariesAndMaps() {
	DictionariesAndMaps(3, []string{
		"sam 99912222",
		"tom 11122222",
		"harry 12299933",
		"sam",
		"edward",
		"harry",
	})
	// Output:
	// sam=99912222
	// Not found
	// harry=12299933
}
