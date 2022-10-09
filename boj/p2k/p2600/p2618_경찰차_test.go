package p2600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2600"
	"1d1go/utils"
)

func TestSolve2618(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2618")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6
3
3 5
5 5
2 3`,
			`9
2
2
1
`,
		},
		{
			`3
3
1 1
1 1
1 1`,
			`0
1
1
1
`,
		},
		{
			`3
3
3 3
3 3
3 3`,
			`0
2
2
2
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2600.Solve2618(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
