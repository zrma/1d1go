package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-class-vs-instance/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleClassAndInstance()

			close(done)
		}()
	})
})

func ExampleClassAndInstance() {
	ClassAndInstance(-1)
	ClassAndInstance(10)
	ClassAndInstance(16)
	ClassAndInstance(18)
	// Output:
	// Age is not valid, setting age to 0.
	// You are young.
	// You are young.
	//
	// You are young.
	// You are a teenager.
	//
	// You are a teenager.
	// You are old.
	//
	// You are old.
	// You are old.
}
