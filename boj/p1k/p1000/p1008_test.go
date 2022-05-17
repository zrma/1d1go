package p1000_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
)

func TestSolve1008(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1008")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"1 3", "0.33333333333333333333333333333333"},
		{"4 5", "0.8"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p1000.Solve1008(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			want, err := strconv.ParseFloat(tt.want, 64)
			assert.NoError(t, err)

			writen := writer.String()
			got, err := strconv.ParseFloat(writen, 64)
			assert.NoError(t, err)

			const epsilon = 1e-6
			assert.InDelta(t, want, got, epsilon)
		})
	}
}
