package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1052(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1052")

	tests := []struct {
		give string
		want string
	}{
		{"3 1", "1"},
		{"3 2", "0"},
		{"3 3", "0"},
		{"13 2", "3"},
		{"1000000 5", "15808"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1052(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
