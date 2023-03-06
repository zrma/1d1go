package p2200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2231(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2231")

	tests := []struct {
		give string
		want string
	}{
		{"1", "0"},
		{"2", "1"},
		{"10", "5"},
		{"11", "10"},
		{"12", "6"},
		{"13", "11"},
		{"110", "0"},
		{"111", "96"},
		{"216", "198"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2231(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
