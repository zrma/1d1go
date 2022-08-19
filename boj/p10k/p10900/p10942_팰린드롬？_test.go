package p10900_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10900"
	"1d1go/utils"
)

func TestSolve10942(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10942")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
1 2 1 3 1 2 1
4
1 3
2 5
3 3
5 7`,
			`1
0
1
1
`,
		},
		{
			`2
1 1
1
1 2`,
			`1
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p10900.Solve10942(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
