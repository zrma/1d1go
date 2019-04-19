package codesprint2019

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/contests/hackerrank-all-womens-codesprint-2019/challenges/name-the-product", func() {
	It("문제를 풀었다", func() {
		Expect(productName([]string{
			"bubby",
			"bunny",
			"berry",
		})).Should(Equal("zzzzz"))
		Expect(productName([]string{
			"ready",
			"stedy",
			"zebra",
		})).Should(Equal("yzzzz"))
		Expect(productName([]string{
			"aaaaa", "bbbbb", "ccccc", "ddddd", "eeeee", "fffff", "ggggg", "hhhhh", "iiiii", "jjjjj",
			"kkkkk", "lllll", "mmmmm", "nnnnn", "ooooo", "ppppp", "qqqqq", "rrrrr", "sssss", "ttttt",
			"uuuuu", "vvvvv", "wwwww", "xxxxx", "yyyyy", "zzzzz", "zzzzz",
		})).Should(Equal("yyyyy"))
		Expect(productName([]string{
			"uuuuu", "vvvvv", "wwwww", "yyyyy", "zzzzz",
		})).Should(Equal("xxxxx"))
	})
})
