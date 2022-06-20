package p2800_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2800"
	"1d1go/utils"
)

func TestSolve2884(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2884")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"10 10", "9 25"},
		{"0 30", "23 45"},
		{"23 40", "22 55"},
		{"0 45", "0 0"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p2800.Solve2884(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
