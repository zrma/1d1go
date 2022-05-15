package p4300_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p4k/p4300"
	"1d1go/utils/mocks"
)

func TestSolve4344(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4344")

	ctrl := gomock.NewController(t)
	scanner := mocks.NewMockScanner(ctrl)
	writer := mocks.NewMockWriter(ctrl)

	arr2D := [][]int{
		{50, 50, 70, 80, 100},
		{100, 95, 90, 80, 70, 60, 50},
		{70, 90, 80},
		{70, 90, 81},
		{100, 99, 98, 97, 96, 95, 94, 93, 91},
	}

	n := strconv.Itoa(len(arr2D))
	scanner.EXPECT().Scan().Return(true)
	scanner.EXPECT().Text().Return(n)

	for _, arr := range arr2D {
		scanner.EXPECT().Scan().Return(true)
		scanner.EXPECT().Text().Return(strconv.Itoa(len(arr)))
		for _, v := range arr {
			scanner.EXPECT().Scan().Return(true)
			scanner.EXPECT().Text().Return(strconv.Itoa(v))
		}
	}

	wants := []string{
		"40.000%",
		"57.143%",
		"33.333%",
		"66.667%",
		"55.556%",
	}
	for _, want := range wants {
		b := []byte(want + "\n")
		writer.EXPECT().Write(b).Return(len(b), nil)
	}

	p4300.Solve4344(scanner, writer)
}
