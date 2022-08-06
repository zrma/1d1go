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

func TestSolve11660(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/11660")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4 4
1 2 3 4
2 3 4 5
3 4 5 6
4 5 6 7
4 4 4 4
2 2 3 4
3 4 3 4
1 1 4 4`,
			`7
27
6
64
`,
		},
		{
			`2 4
1 2
3 4
1 1 1 1
1 2 1 2
2 1 2 1
2 2 2 2`,
			`1
2
3
4
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p11600.Solve11660(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
