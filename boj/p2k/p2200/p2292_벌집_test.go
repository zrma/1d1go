package p2200_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2200"
	"1d1go/utils"
)

func TestSolve2292(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2292")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"1", "1"},
		{"2", "2"},
		{"7", "2"},
		{"8", "3"},
		{"19", "3"},
		{"20", "4"},
		{"37", "4"},
		{"38", "5"},
		{"61", "5"},
		{"62", "6"},
		{"91", "6"},
		{"92", "7"},
		{"1000000000", "18258"},
		{"1234567890", "20287"},
		{"9876543210", "57379"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2200.Solve2292(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
