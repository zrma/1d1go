package p7500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve7562(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/7562")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
8
0 0
7 0
100
0 0
30 50
10
1 1
1 1`,
			`5
28
0
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve7562(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
