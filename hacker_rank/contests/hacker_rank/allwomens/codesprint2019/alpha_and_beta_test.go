package codesprint2019

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/contests/hackerrank-all-womens-codesprint-2019/challenges/alpha-and-beta", func() {
	It("문제를 풀었다", func() {
		Expect(alphaBeta([]int32{1, 2, 3, 3, 2, 1})).Should(BeNumerically("==", 2))
		Expect(alphaBeta([]int32{1, 2, 3, 4, 5})).Should(BeNumerically("==", 4))
		Expect(alphaBeta([]int32{1, 2, 3, 2, 1})).Should(BeNumerically("==", 2))
		Expect(alphaBeta([]int32{1, 2, 4, 3, 2, 1})).Should(BeNumerically("==", 3))
	})
})
