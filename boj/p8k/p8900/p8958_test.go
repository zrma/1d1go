package p8900_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p8k/p8900"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve8958(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/8958")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//goland:noinspection SpellCheckingInspection
	const s = `5
OOXXOXXOOO
OOXXOOXXOO
OXOXOXOXOXOXOX
OOOOOOOOOO
OOOOXOOOOXOOOOX`
	scanner := utils.NewStringScanner(s)
	writer := mocks.NewMockWriter(ctrl)

	for _, want := range []int{10, 9, 7, 55, 30} {
		b := []byte(strconv.Itoa(want) + "\n")
		writer.EXPECT().Write(b).Return(len(b), nil)
	}

	p8900.Solve8958(scanner, writer)
}
