package p2100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2162(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2162")

	tests := []struct {
		give string
		want string
	}{
		{
			`3
1 1 2 3
2 1 0 0
1 0 1 1`,
			`1
3
`,
		},
		{
			`3
-1 -1 1 1
-2 -2 2 2
0 1 -1 0`,
			`2
2
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2162(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
