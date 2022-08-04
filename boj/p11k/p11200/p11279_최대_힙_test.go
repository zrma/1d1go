package p11200_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11200"
	"1d1go/utils"
)

func TestSolve11279(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11279")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`13
0
1
2
0
0
3
2
1
0
0
0
0
0`,
			`0
2
1
3
2
1
0
0
`,
		},
		{
			`12
1
2
3
4
5
6
0
0
0
0
0
0
`,
			`6
5
4
3
2
1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11200.Solve11279(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
