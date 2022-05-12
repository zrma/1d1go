package p10900

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSolve10951(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10951")

	for i, tt := range []struct {
		a, b int
		want string
	}{
		{1, 1, "2"},
		{2, 3, "5"},
		{3, 4, "7"},
		{9, 8, "17"},
		{5, 2, "7"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := NewMockWriter(ctrl)

			want := []byte(fmt.Sprintf("%s\n", tt.want))
			m.EXPECT().Write(want).Return(0, nil)

			Solve10951(m, tt.a, tt.b)
		})
	}
}
