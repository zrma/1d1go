package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1712(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1712")

	tests := []struct {
		give string
		want string
	}{
		{"1000 70 170", "11"},
		{"3 2 1", "-1"},
		{"2100000000 9 10", "2100000001"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1712(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
