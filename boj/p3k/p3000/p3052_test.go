package p3000_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p3k/p3000"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve3052(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3052")

	for i, tt := range []struct {
		s    string
		want int
	}{
		{
			`1
2
3
4
5
6
7
8
9
10`,
			10,
		},
		{
			`42
84
252
420
840
126
42
84
420
126`,
			1,
		},
		{
			`39
40
41
42
43
44
82
83
84
85`,
			6,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctrl)

			want := []byte(strconv.Itoa(tt.want) + "\n")
			writer.EXPECT().Write(want).Return(len(want), nil)

			p3000.Solve3052(scanner, writer)
		})
	}
}
