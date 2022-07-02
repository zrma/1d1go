package p1600_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1600"
	"1d1go/utils"
)

func TestSolve1676(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1676")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"1", "0"},
		{"2", "0"},
		{"3", "0"},
		{"4", "0"},
		{"5", "1"},
		{"6", "1"},
		{"7", "1"},
		{"8", "1"},
		{"9", "1"},
		{"10", "2"},
		{"14", "2"},
		{"15", "3"},
		{"19", "3"},
		{"20", "4"},
		{"24", "4"},
		{"25", "6"},
		{"29", "6"},
		{"30", "7"},
		{"75", "18"},
		{"79", "18"},
		{"80", "19"},
		{"84", "19"},
		{"85", "20"},
		{"124", "28"},
		{"125", "31"},
		{"499", "121"},
		{"500", "124"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1600.Solve1676(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
