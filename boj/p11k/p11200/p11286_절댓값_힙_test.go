package p11200_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11200"
)

func TestSolve11286(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11286")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`18
1
-1
0
0
0
1
1
-1
-1
2
-2
0
0
0
0
0
0
0`,
			`-1
1
0
-1
-1
1
1
-2
2
0
`,
		},
		{
			`12
6
5
4
3
2
1
0
0
0
0
0
0
`,
			`1
2
3
4
5
6
`,
		},
		{
			`12
-6
-5
-4
-3
-2
-1
0
0
0
0
0
0
`,
			`-1
-2
-3
-4
-5
-6
`,
		},
		{
			`12
-1
-1
-1
1
1
-1
0
0
0
0
0
0`,
			`-1
-1
-1
-1
1
1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p11200.Solve11286(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
