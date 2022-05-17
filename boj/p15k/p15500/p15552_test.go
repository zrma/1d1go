package p15500

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve15552(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15552")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const s = `5
1 1
12 34
5 500
40 60
1000 1000`
	scanner := utils.NewStringScanner(s)
	writer := mocks.NewMockWriter(ctrl)

	for _, want := range []int{2, 46, 505, 100, 2000} {
		b := []byte(strconv.Itoa(want) + "\n")
		writer.EXPECT().Write(b).Return(len(b), nil)
	}

	Solve15552(scanner, writer)
}
