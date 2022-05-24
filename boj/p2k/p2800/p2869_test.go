package p2800_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2800"
	"1d1go/utils"
)

func TestSolve2869(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2869")

	for _, tt := range []struct {
		s    string
		want string
	}{
		{"1 0 1", "1"},
		{"2 1 1", "1"},
		{"2 1 2", "1"},
		{"3 1 4", "2"},
		{"3 1 5", "2"},
		{"2 1 5", "4"},
		{"5 1 6", "2"},
		// 10 -> 7 ->
		// 17 -> 14 ->
		// 24 -> 21 ->
		// 31 -> 28 ->
		// 38 -> 35 ->
		// 45 -> 42 ->
		// 52 -> 49 ->
		// 59 -> 56 ->
		// 66 -> 63 ->
		// 73 -> 70 ->
		// 80 -> 77 ->
		{"10 3 66", "9"},
		{"10 3 67", "10"},
		{"10 3 73", "10"},
		{"10 3 74", "11"},
		{"100 99 1000000000", "999999901"},
	} {
		t.Run(tt.s, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.s)
			writer := utils.NewStringWriter()

			p2800.Solve2869(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
