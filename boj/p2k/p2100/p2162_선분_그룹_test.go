package p2100_test

import (
	"1d1go/boj/p2k/p2100"
	"1d1go/utils"
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSolve2162(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2162")

	for i, tt := range []struct {
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
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2100.Solve2162(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
