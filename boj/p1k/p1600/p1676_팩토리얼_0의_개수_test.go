package p1600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1676(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1676")

	for i, tt := range []struct {
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1676(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
