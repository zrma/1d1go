package p2200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2206(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2206")

	tests := []struct {
		give string
		want string
	}{
		{
			`6 4
0100
1110
1000
0000
0111
0000`,
			"15",
		},
		{
			`4 4
0111
1111
1111
1110`,
			"-1",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2206(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
