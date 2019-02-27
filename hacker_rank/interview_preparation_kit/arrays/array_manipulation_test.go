package arrays

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/crush/problem", func() {
	DescribeTable("문제를 풀었다",
		func(n int, queries [][]int32, expected int) {
			actual := arrayManipulation(int32(n), queries)
			Expect(actual).Should(BeNumerically("==", expected))
		},
		Entry("0", 10, [][]int32{
			{1, 5, 3}, {4, 8, 7}, {6, 9, 1},
		}, 10),
		Entry("1", 5, [][]int32{
			{1, 2, 100}, {2, 5, 100}, {3, 4, 100},
		}, 200),
		Entry("2", 10, [][]int32{
			{2, 6, 8}, {3, 5, 7}, {1, 8, 1}, {5, 9, 15},
		}, 31),
		Entry("3", 10, [][]int32{
			{3, 5, 7}, {2, 6, 8}, {1, 8, 1}, {5, 9, 15},
		}, 31),
	)

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long array", func() {
			file, err := os.Open("./test_data/array_manipulation.csv")
			Expect(err).Should(BeNil())

			r := csv.NewReader(bufio.NewReader(file))
			rows, err := r.ReadAll()
			Expect(err).Should(BeNil())

			var arr [][]int32
			for _, row := range rows {
				begin, err := strconv.ParseInt(row[0], 10, 64)
				Expect(err).Should(BeNil())

				end, err := strconv.ParseInt(row[1], 10, 64)
				Expect(err).Should(BeNil())

				value, err := strconv.ParseInt(row[2], 10, 64)
				Expect(err).Should(BeNil())

				arr = append(arr, []int32{int32(begin), int32(end), int32(value)})
			}

			Expect(len(arr)).Should(Equal(100000))
			actual := arrayManipulation(10000000, arr)
			Expect(actual).Should(BeNumerically("==", 2506721627))
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 5), "시간 초과")
	}, 10)
})