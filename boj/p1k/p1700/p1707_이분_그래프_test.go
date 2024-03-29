package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1707(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1707")

	tests := []struct {
		give string
		want string
	}{
		{
			`2
3 2
1 3
2 3
4 4
1 2
2 3
3 4
4 2`,
			`YES
NO
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1707(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
