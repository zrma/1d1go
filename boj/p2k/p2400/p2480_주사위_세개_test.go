package p2400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2480(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2480")

	tests := []struct {
		give string
		want string
	}{
		{"3 3 6", "1300"},
		{"6 3 3", "1300"},
		{"2 2 2", "12000"},
		{"6 2 5", "600"},
		{"2 6 5", "600"},
		{"2 5 6", "600"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2480(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
