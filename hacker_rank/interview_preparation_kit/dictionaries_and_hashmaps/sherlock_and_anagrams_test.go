package dictionaries_and_hashmaps

import (
	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem", func() {
	DescribeTable("문제를 풀었다",
		func(s string, i int) {
			actual := sherlockAndAnagrams(s)
			Expect(actual).Should(BeNumerically("==", i))
		},
		Entry("test_0", "abba", 4),
		Entry("test_1", "abcd", 0),
		Entry("test_2", "ifailuhkqq", 3),
		Entry("test_3", "kkkk", 10),
		Entry("test_4", "cdcd", 5),
	)

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long string", func() {
			actual := sherlockAndAnagrams("ifailuhkqq")
			Expect(actual).Should(BeNumerically("==", 3))
			actual = sherlockAndAnagrams("dbcfibibcheigfccacfegicigcefieeeeegcghggdheichgafhdigffgifidfbeaccadabecbdcgieaffbigffcecahafcafhcdg")
			Expect(actual).Should(BeNumerically("==", 1464))
			actual = sherlockAndAnagrams("dfcaabeaeeabfffcdbbfaffadcacdeeabcadabfdefcfcbbacadaeafcfceeedacbafdebbffcecdbfebdbfdbdecbfbadddbcec")
			Expect(actual).Should(BeNumerically("==", 2452))
			actual = sherlockAndAnagrams("gjjkaaakklheghidclhaaeggllagkmblhdlmihmgkjhkkfcjaekckaafgabfclmgahmdebjekaedhaiikdjmfbmfbdlcafamjbfe")
			Expect(actual).Should(BeNumerically("==", 873))
			actual = sherlockAndAnagrams("fdbdidhaiqbggqkhdmqhmemgljaphocpaacdohnokfqmlpmiioedpnjhphmjjnjlpihmpodgkmookedkehfaceklbljcjglncfal")
			Expect(actual).Should(BeNumerically("==", 585))
			actual = sherlockAndAnagrams("bcgdehhbcefeeadchgaheddegbiihehcbbdffiiiifgibhfbchffcaiabbbcceabehhiffggghbafabbaaebgediafabafdicdhg")
			Expect(actual).Should(BeNumerically("==", 1305))
			actual = sherlockAndAnagrams("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			Expect(actual).Should(BeNumerically("==", 166650))
			actual = sherlockAndAnagrams("mhmgmbbccbbaffhbncgndbffkjbhmkfncmihhdhcebmchnfacdigflhhbekhfejblegakjjiejeenibemfmkfjbkkmlichlkbnhc")
			Expect(actual).Should(BeNumerically("==", 840))
			actual = sherlockAndAnagrams("fdacbaeacbdbaaacafdfbbdcefadgfcagdfcgbgeafbfbggdedfbdefdbgbefcgdababafgffedbefdecbaabdaafggceffbacgb")
			Expect(actual).Should(BeNumerically("==", 2134))
			actual = sherlockAndAnagrams("bahdcafcdadbdgagdddcidaaicggcfdbfeeeghiibbdhabdhffddhffcdccfdddhgiceciffhgdibfdacbidgagdadhdceibbbcc")
			Expect(actual).Should(BeNumerically("==", 1571))
			actual = sherlockAndAnagrams("dichcagakdajjhhdhegiifiiggjebejejciaabbifkcbdeigajhgfcfdgekfajbcdifikafkgjjjfefkdbeicgiccgkjheeiefje")
			Expect(actual).Should(BeNumerically("==", 1042))
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 5), "시간 초과")
	}, 10)
})
