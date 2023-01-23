package p1100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1193(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1193")

	for i, tt := range []struct {
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1193(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
