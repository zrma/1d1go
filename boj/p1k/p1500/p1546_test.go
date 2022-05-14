package p1500_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1500"
	"1d1go/utils/mocks"
)

func TestSolve1546(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1546")

	for i, tt := range []struct {
		n      int
		scores []int
		want   float64
	}{
		{
			3,
			[]int{40, 80, 60},
			75.0,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := mocks.NewMockScanner(ctrl)
			writer := mocks.NewMockWriter(ctrl)

			n := strconv.Itoa(tt.n)

			scanner.EXPECT().Scan().Return(true)
			scanner.EXPECT().Text().Return(n)

			for _, score := range tt.scores {
				s := strconv.Itoa(score)
				scanner.EXPECT().Scan().Return(true)
				scanner.EXPECT().Text().Return(s)
			}

			writer.EXPECT().Write(gomock.Any()).Do(func(b []byte) {
				s := string(b)
				s = strings.TrimSuffix(s, "\n")
				got, err := strconv.ParseFloat(s, 64)
				assert.NoError(t, err)

				const delta = 1e-9
				assert.InDelta(t, tt.want, got, delta)
			}).Return(0, nil)

			p1500.Solve1546(scanner, writer)
		})
	}
}
