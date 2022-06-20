package p1400_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1400"
	"1d1go/utils"
)

func TestSolve1463(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1463")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"1", "0"},
		{"2", "1"},
		{"3", "1"},
		{"4", "2"},
		{"5", "3"},
		{"6", "2"},
		{"7", "3"},
		{"8", "3"},
		{"9", "2"},
		{"10", "3"},
		{"11", "4"},
		{"12", "3"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p1400.Solve1463(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
