package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1003(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1003")

	tests := []struct {
		give string
		want string
	}{
		{
			`3
0
1
3`,
			`1 0
0 1
1 2
`,
		},
		{
			`2
6
22`,
			`5 8
10946 17711
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1003(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
