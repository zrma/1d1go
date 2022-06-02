package p2200_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2200"
	"1d1go/utils"
)

func TestSolve2231(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2231")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"1", "0"},
		{"2", "1"},
		{"10", "5"},
		{"11", "10"},
		{"12", "6"},
		{"13", "11"},
		{"110", "0"},
		{"111", "96"},
		{"216", "198"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p2200.Solve2231(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
