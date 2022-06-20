package p2700_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
	"1d1go/utils"
)

func TestSolve2753(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2753")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"2000", "1"},
		{"1999", "0"},
		{"2100", "0"},
		{"2104", "1"},
		{"2200", "0"},
		{"2400", "1"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p2700.Solve2753(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
