package p10900_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve10952(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10952")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const s = `1 1
2 3
3 4
9 8
5 2
0 0`
	scanner := utils.NewStringScanner(s)
	writer := mocks.NewMockWriter(ctrl)

	for _, want := range []string{"2", "5", "7", "17", "7"} {
		b := []byte(want + "\n")
		writer.EXPECT().Write(b).Return(len(b), nil)
	}
	writer.EXPECT().Flush().Return(nil)

	p10900.Solve10952(scanner, writer)
}

func TestSolve10952_StopAbnormally(t *testing.T) {
	t.Run("first scan returns false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		const s = ``
		scanner := utils.NewStringScanner(s)
		writer := mocks.NewMockWriter(ctrl)

		writer.EXPECT().Flush().Return(nil)

		p10900.Solve10952(scanner, writer)
	})

	t.Run("second scan returns false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		const s = `1`
		scanner := utils.NewStringScanner(s)
		writer := mocks.NewMockWriter(ctrl)

		writer.EXPECT().Flush().Return(nil)

		p10900.Solve10952(scanner, writer)
	})
}
