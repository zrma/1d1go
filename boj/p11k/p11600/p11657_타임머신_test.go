package p11600_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p11k/p11600"
	"1d1go/utils"
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
			writer := utils.NewStringWriter()

			p11600.Solve11657(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
