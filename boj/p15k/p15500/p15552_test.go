package p15500

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/utils/mocks"
)

func TestSolve15552(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15552")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scanner := mocks.NewMockScanner(ctrl)
	writer := mocks.NewMockWriter(ctrl)

	arr2D := [][2]int{
		{1, 1}, {2, 3}, {3, 4}, {9, 8}, {5, 2},
	}
	wants := []int{2, 5, 7, 17, 7}

	n := strconv.Itoa(len(arr2D))

	scanner.EXPECT().Scan().Return(true)
	scanner.EXPECT().Text().Return(n)

	for i, arr := range arr2D {
		a := strconv.Itoa(arr[0])
		b := strconv.Itoa(arr[1])

		scanner.EXPECT().Scan().Return(true).Times(2)
		scanner.EXPECT().Text().Return(a)
		scanner.EXPECT().Text().Return(b)

		want := []byte(strconv.Itoa(wants[i]) + "\n")
		writer.EXPECT().Write(want).Return(len(want), nil)
	}

	Solve15552(scanner, writer)
}
