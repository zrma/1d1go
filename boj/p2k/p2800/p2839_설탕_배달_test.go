package p2800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2839(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2839")

	for i, tt := range []struct {
		give string
		want string
	}{
		{"0", "-1"},
		{"1", "-1"},
		{"2", "-1"},
		{"3", "1"},
		{"4", "-1"},
		{"5", "1"},
		{"6", "2"},
		{"7", "-1"},
		{"8", "2"},
		{"9", "3"},
		{"10", "2"},
		{"11", "3"},
		{"12", "4"},
		{"13", "3"},
		{"14", "4"},
		{"15", "3"},
		{"16", "4"},
		{"17", "5"},
		{"18", "4"},
		{"100", "20"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2839(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
