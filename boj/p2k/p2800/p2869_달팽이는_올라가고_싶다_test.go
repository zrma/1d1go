package p2800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2800"
)

func TestSolve2869(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2869")

	for i, tt := range []struct {
		give string
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2800.Solve2869(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
