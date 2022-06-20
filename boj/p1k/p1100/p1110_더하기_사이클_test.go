package p1100_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1100"
	"1d1go/utils"
)

func TestSolve1110(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1110")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"26", "4"},
		{"55", "3"},
		{"1", "60"},
		{"0", "1"},
		{"71", "12"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p1100.Solve1110(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
