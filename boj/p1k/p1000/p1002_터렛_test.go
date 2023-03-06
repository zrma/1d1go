package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1002(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1002")

	tests := []struct {
		give string
		want string
	}{
		{
			`7
0 0 13 40 0 37
0 0 3 0 7 4
1 1 1 1 1 5
1 1 1 1 1 1
0 0 3 0 1 1
0 0 1 0 3 1
0 0 2 0 1 1`,
			`2
1
0
-1
0
0
1
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1002(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
