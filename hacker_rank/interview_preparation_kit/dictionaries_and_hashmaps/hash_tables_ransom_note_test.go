package dictionaries_and_hashmaps

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/ctci-ransom-note/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleCheckMagazine()

			close(done)
		}()
	})
})

func ExampleCheckMagazine() {
	CheckMagazine(
		[]string{"give", "me", "one", "grand", "today", "night"},
		[]string{"give", "one", "grand", "today"},
	)
	CheckMagazine(
		[]string{"two", "times", "three", "is", "not", "four"},
		[]string{"two", "times", "two", "is", "four"},
	)
	CheckMagazine(
		[]string{"ive", "got", "a", "lovely", "brunch", "of", "coconuts"},
		[]string{"ive", "got", "some", "coconuts"},
	)
	// Output:
	// Yes
	// No
	// No
}
