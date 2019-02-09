package warm_up

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/time-conversion/problem", func() {
	It("문제를 풀었다", func() {
		s := "07:05:45PM"
		actual := timeConversion(s)
		Expect(actual).Should(Equal("19:05:45"))
	})
})
