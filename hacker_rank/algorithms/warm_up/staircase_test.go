package warm_up

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("https://www.hackerrank.com/challenges/staircase/problem", func() {
	It("문제를 풀었다", func(done Done) {
		go func() {
			defer GinkgoRecover()

			ExampleStaircase()

			close(done)
		}()
	})
})

func ExampleStaircase() {
	var n int32 = 6

	Staircase(n)
	// Output:
	//      #
	//     ##
	//    ###
	//   ####
	//  #####
	// ######
}
