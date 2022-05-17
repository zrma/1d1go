package p10900_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve10950(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10950")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const s = `5
1 1
2 3
3 4
9 8
5 2`
	scanner := utils.NewStringScanner(s)
	writer := mocks.NewMockWriter(ctrl)

	for _, want := range []int{2, 5, 7, 17, 7} {
		b := []byte(strconv.Itoa(want) + "\n")
		writer.EXPECT().Write(b).Return(len(b), nil)
	}

	p10900.Solve10950(scanner, writer)
}
