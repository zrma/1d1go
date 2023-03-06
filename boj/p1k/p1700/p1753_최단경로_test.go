package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1753(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1753")

	tests := []struct {
		give string
		want string
	}{
		{
			`5 6
1
5 1 1
1 2 2
1 3 3
2 3 4
2 4 5
3 4 6`,
			`0
2
3
7
INF
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1753(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
