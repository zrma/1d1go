package p2900_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2900"
	"1d1go/utils"
)

func TestSolve2908(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2908")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"734 893", "437"},
		{"221 231", "132"},
		{"839 237", "938"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p2900.Solve2908(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
