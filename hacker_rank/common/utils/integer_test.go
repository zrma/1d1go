package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("정수형 관련 유틸 함수 검증", func() {
	Context("Min 함수는", func() {
		It("두 정수 중 작은 수를 반환한다.", func() {
			actual := Min(3, 5)
			Expect(actual).Should(BeNumerically("==", 3))

			actual = Min(5, 3)
			Expect(actual).Should(BeNumerically("==", 3))
		})
	})

	Context("Max 함수는", func() {
		It("두 정수 중 큰 수를 반환한다.", func() {
			actual := Max(3, 5)
			Expect(actual).Should(BeNumerically("==", 5))

			actual = Max(5, 3)
			Expect(actual).Should(BeNumerically("==", 5))
		})
	})

	Context("Pow 함수는", func() {
		It("거듭제곱 연산을 잘 수행한다.", func() {
			actual := Pow(3, 5)
			expected := math.Pow(3, 5)
			Expect(actual).Should(BeNumerically("==", expected))

			actual = Pow(0, 5)
			Expect(actual).Should(BeNumerically("==", 0))

			actual = Pow(10, 0)
			Expect(actual).Should(BeNumerically("==", 1))
		})
	})
})
