package p11600

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve11657(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11657")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 4
1 2 4
1 3 3
2 3 -1
3 1 -2`,
			`4
3
`,
		},
		{
			`3 4
1 2 4
1 3 3
2 3 -4
3 1 -2`,
			`-1
`,
		},
		{
			`3 2
1 2 4
1 2 3`,
			`3
-1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve11657(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
