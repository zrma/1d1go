package p1400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1463(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1463")

	tests := []struct {
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
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1463(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
