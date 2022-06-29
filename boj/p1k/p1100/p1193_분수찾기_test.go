package p1100_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1100"
	"1d1go/utils"
)

func TestSolve1193(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1193")

	for _, tt := range []struct {
		give string
		want string
	}{
		{"1", "1/1"},
		{"2", "1/2"},
		{"3", "2/1"},
		{"4", "3/1"},
		{"5", "2/2"},
		{"6", "1/3"},
		{"7", "1/4"},
		{"8", "2/3"},
		{"9", "3/2"},
		{"10", "4/1"},
		{"11", "5/1"},
		{"12", "4/2"},
		{"13", "3/3"},
		{"14", "2/4"},
		{"15", "1/5"},
		{"16", "1/6"},
		{"17", "2/5"},
		{"18", "3/4"},
		{"19", "4/3"},
		{"20", "5/2"},
		{"21", "6/1"},
		{"22", "7/1"},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1100.Solve1193(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
