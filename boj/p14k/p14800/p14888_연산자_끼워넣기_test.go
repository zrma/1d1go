package p14800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p14k/p14800"
	"1d1go/utils"
)

func TestSolve14888(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/14888")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
5 6
0 0 1 0`,
			`30
30
`,
		},
		{
			`3
3 4 5
1 0 1 0`,
			`35
17
`,
		},
		{
			`6
1 2 3 4 5 6
2 1 1 1`,
			`54
-24
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p14800.Solve14888(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
