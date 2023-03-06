package p2100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2178(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2178")

	tests := []struct {
		give string
		want string
	}{
		{
			`4 6
101111
101010
101011
111011`,
			"15",
		},
		{
			`4 6
110110
110110
111111
111101`,
			"9",
		},
		{
			`2 25
1011101110111011101110111
1110111011101110111011101`,
			"38",
		},
		{
			`7 7
1011111
1110001
1000001
1000001
1000001
1000001
1111111`,
			"13",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2178(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
