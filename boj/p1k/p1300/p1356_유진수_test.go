package p1300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1356(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1356")

	tests := []struct {
		give string
		want string
	}{
		{"1236", "YES"},
		{"1", "NO"},
		{"1221", "YES"},
		{"123", "NO"},
		{"1234", "NO"},
		{"142", "NO"},
		{"4729382", "NO"},
		{"42393338", "YES"},
		{"1001", "YES"},
		{"1000", "YES"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1356(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
