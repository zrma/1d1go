package utils

import (
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrintTest 함수 검증", func() {
	It("정상 동작", func() {
		err := PrintTest(func() error {
			fmt.Println("test")
			return nil
		}, []string{
			"test",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("오류 상황", func() {
		err := PrintTest(func() error {
			return errors.New("test")
		}, []string{})
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("반환값 없음", func() {
		err := PrintTest(func() error {
			return nil
		}, []string{})
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("반환값이 있어야 하는데 안 나오는 경우", func() {
		err := PrintTest(func() error {
			return nil
		}, []string{
			"failed",
		})
		Expect(err).Should(HaveOccurred())
	})
})
