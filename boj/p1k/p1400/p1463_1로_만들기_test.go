package p1400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1400"
)

func TestSolve1463(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1463")

	for i, tt := range []struct {
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1400.Solve1463(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
