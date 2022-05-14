package p10900_test

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils/mocks"
)

func TestSolve10952(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10952")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scanner := mocks.NewMockScanner(ctrl)
	writer := mocks.NewMockWriter(ctrl)

	for _, tt := range []struct {
		a, b int
		want string
	}{
		{1, 1, "2"},
		{2, 3, "5"},
		{3, 4, "7"},
		{9, 8, "17"},
		{5, 2, "7"},
	} {
		a := strconv.Itoa(tt.a)
		b := strconv.Itoa(tt.b)

		scanner.EXPECT().Scan().Return(true).Times(2)
		scanner.EXPECT().Text().Return(a)
		scanner.EXPECT().Text().Return(b)

		want := []byte(tt.want + "\n")
		writer.EXPECT().Write(want).Return(len(want), nil)
	}
	writer.EXPECT().Flush().Return(nil)

	const exitCode = "0"
	scanner.EXPECT().Scan().Return(true).Times(2)
	scanner.EXPECT().Text().Return(exitCode)
	scanner.EXPECT().Text().Return(exitCode)

	p10900.Solve10952(scanner, writer)
}

func TestSolve10952_StopAbnormally(t *testing.T) {
	t.Run("first scan returns false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		scanner := mocks.NewMockScanner(ctrl)
		writer := mocks.NewMockWriter(ctrl)

		scanner.EXPECT().Scan().Return(false)
		writer.EXPECT().Flush().Return(nil)

		p10900.Solve10952(scanner, writer)
	})

	t.Run("second scan returns false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		scanner := mocks.NewMockScanner(ctrl)
		writer := mocks.NewMockWriter(ctrl)

		scanner.EXPECT().Scan().Return(true)
		scanner.EXPECT().Text().Return("awesome")
		scanner.EXPECT().Scan().Return(false)
		writer.EXPECT().Flush().Return(nil)

		p10900.Solve10952(scanner, writer)
	})
}
