package p1500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1504(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1504")

	tests := []struct {
		give string
		want string
	}{
		{
			`4 6
1 2 3
2 3 3
3 4 1
1 3 5
2 4 5
1 4 4
2 3`,
			"7",
		},
		{
			`4 2
1 2 1
1 3 1
3 4
`,
			"-1",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1504(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
