package p4900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve4948(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/4948")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1
10
13
100
1000
10000
100000
0`,
			`1
4
3
21
135
1033
8392
`,
		},
		{
			`1
10
13
invalid`,
			`1
4
3
`,
		},
	} {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve4948(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
