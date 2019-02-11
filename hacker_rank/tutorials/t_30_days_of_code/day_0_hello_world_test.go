package t_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-hello-world/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleHelloWorld()

			close(done)
		}()
	})
})

func ExampleHelloWorld() {
	HelloWorld("Welcome to 30 Days of Code!")
	// Output:
	// Hello, World.
	// Welcome to 30 Days of Code!
}
