package p10900_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve10926(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10926")

	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want string
	}{
		{"joonas", "joonas??!"},
		{"baekjoon", "baekjoon??!"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctrl)

			want := []byte(tt.want + "\n")
			writer.EXPECT().Write(want).Return(len(want), nil)

			p10900.Solve10926(scanner, writer)
		})
	}
}
