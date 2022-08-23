package p24400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p24k/p24400"
	"1d1go/utils"
)

func TestSolve24479(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/24479")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 5 1
1 4
1 2
2 3
2 4
3 4`,
			`1
2
3
4
0
`,
		},
		{
			`5 2 2
1 2
2 3`,
			`2
1
3
0
0
`,
		},
		{
			`6 7 1
1 6
1 2
2 6
2 3
2 4
3 4
3 5
4 5`,
			`1
2
3
4
5
6
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p24400.Solve24479(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
