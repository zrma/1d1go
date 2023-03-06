package p2200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2212(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2212")

	tests := []struct {
		give string
		want string
	}{
		{
			`6
2
1 6 9 3 6 7`,
			"5",
		},
		{
			`10
5
20 3 14 6 7 8 18 10 12 15`,
			"7",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2212(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
