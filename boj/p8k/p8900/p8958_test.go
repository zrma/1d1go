package p8900_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p8k/p8900"
	"1d1go/utils/mocks"
)

func TestSolve8958(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/8958")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scanner := mocks.NewMockScanner(ctrl)
	writer := mocks.NewMockWriter(ctrl)

	//goland:noinspection SpellCheckingInspection
	inputs := []string{
		"OOXXOXXOOO",
		"OOXXOOXXOO",
		"OXOXOXOXOXOXOX",
		"OOOOOOOOOO",
		"OOOOXOOOOXOOOOX",
	}
	wants := []int{10, 9, 7, 55, 30}

	n := strconv.Itoa(len(inputs))
	scanner.EXPECT().Scan().Return(true)
	scanner.EXPECT().Text().Return(n)

	for _, input := range inputs {
		scanner.EXPECT().Scan().Return(true)
		scanner.EXPECT().Text().Return(input)
	}

	for _, want := range wants {
		b := []byte(strconv.Itoa(want) + "\n")
		writer.EXPECT().Write(b).Return(len(b), nil)
	}

	p8900.Solve8958(scanner, writer)
}
