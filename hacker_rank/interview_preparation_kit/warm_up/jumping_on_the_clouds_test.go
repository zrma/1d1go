package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem", func() {
	It("문제를 풀었다", func() {
		c := []int32{0, 0, 1, 0, 0, 1, 0}
		actual := jumpingOnClouds(c)
		Expect(actual).Should(BeNumerically("==", 4))

		c = []int32{0, 0, 0, 0, 1, 0}
		actual = jumpingOnClouds(c)
		Expect(actual).Should(BeNumerically("==", 3))

		c = []int32{0, 0, 0, 0, 0}
		actual = jumpingOnClouds(c)
		Expect(actual).Should(BeNumerically("==", 2))

		c = []int32{0, 1, 0, 0, 1, 0}
		actual = jumpingOnClouds(c)
		Expect(actual).Should(BeNumerically("==", 3))

		c = []int32{0, 1, 1, 0, 1, 0}
		actual = jumpingOnClouds(c)
		Expect(actual).Should(BeNumerically("==", -1))

		c = []int32{0, 0, 1, 0, 1, 0}
		actual = jumpingOnClouds(c)
		Expect(actual).Should(BeNumerically("==", 3))

		c = []int32{0, 0, 0, 1, 0, 0}
		actual = jumpingOnClouds(c)
		Expect(actual).Should(BeNumerically("==", 3))
	})
})
