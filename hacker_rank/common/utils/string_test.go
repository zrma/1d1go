package utils

import (
	"github.com/go-test/deep"
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("문자열 관련 유틸 함수 검증", func() {
	Context("SortString 함수는", func() {
		It("하나의 문자열을 잘 정렬한다.", func() {
			actual := SortString("dcba")
			Expect(actual).Should(Equal("abcd"))

			actual = SortString("ffbbaa")
			Expect(actual).Should(Equal("aabbff"))
		})
	})

	Context("SortAdapter 구조체는", func() {
		actual := []string{"def", "abc", "bcd"}

		It("정상적으로 잘 정렬한다.", func() {
			expect := []string{"abc", "bcd", "def"}

			sort.Sort(SortAdapter(actual))
			Expect(deep.Equal(actual, expect)).Should(BeNil())
		})

	})
})
