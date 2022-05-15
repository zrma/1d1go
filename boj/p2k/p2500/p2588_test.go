package p2500_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p2k/p2500"
	"1d1go/utils/mocks"
)

func TestSolve2588(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2588")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scanner := mocks.NewMockScanner(ctrl)
	writer := mocks.NewMockWriter(ctrl)

	const a, b = 472, 385
	wants := []int{2360, 3776, 1416, 181720}

	scanner.EXPECT().Scan().Return(true).Times(2)
	scanner.EXPECT().Text().Return(strconv.Itoa(a))
	scanner.EXPECT().Text().Return(strconv.Itoa(b))

	for _, want := range wants {
		b := []byte(strconv.Itoa(want) + "\n")
		writer.EXPECT().Write(b).Return(len(b), nil)
	}

	p2500.Solve2588(scanner, writer)
}
