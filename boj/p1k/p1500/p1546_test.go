package p1500_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1500"
	"1d1go/utils"
	"1d1go/utils/mocks"
)

func TestSolve1546(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1546")

	for i, tt := range []struct {
		s    string
		want float64
	}{
		{
			`3
40 80 60`,
			75.0,
		},
		{
			`3
10 20 30`,
			66.666667,
		},
		{
			`4
1 100 100 100`,
			75.25,
		},
		{
			`5
1 2 4 8 16`,
			38.75,
		},
		{
			`2
3 10`,
			65.0,
		},
		{
			`4
10 20 0 100`,
			32.5,
		},
		{
			`1
50`,
			100.0,
		},
		{
			`9
10 20 30 40 50 60 70 80 90`,
			55.55555555555556,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			scanner := utils.NewStringScanner(tt.s)
			writer := mocks.NewMockWriter(ctrl)

			writer.EXPECT().Write(gomock.Any()).Do(func(b []byte) {
				s := string(b)
				s = strings.TrimSuffix(s, "\n")
				got, err := strconv.ParseFloat(s, 64)
				assert.NoError(t, err)

				const delta = 1e-5
				assert.InDelta(t, tt.want, got, delta)
			}).Return(0, nil)

			p1500.Solve1546(scanner, writer)
		})
	}
}
