package university_codesprint_5

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/contests/university-codesprint-5/challenges/exceeding-speed-limit", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleExceedingTheSpeedLimit()

			close(done)
		}()
	})
})

func ExampleExceedingTheSpeedLimit() {
	ExceedingTheSpeedLimit(100)
	ExceedingTheSpeedLimit(140)
	ExceedingTheSpeedLimit(85)
	// Output:
	// 3000 Warning
	// 25000 License removed
	// 0 No punishment
}
