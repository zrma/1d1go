package p2900_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2900"
	"1d1go/utils"
)

func TestSolve2920(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2920")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"1 2 3 4 5 6 7 8", "ascending"},
		{"8 7 6 5 4 3 2 1", "descending"},
		{"8 1 7 2 6 3 5 4", "mixed"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p2900.Solve2920(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
