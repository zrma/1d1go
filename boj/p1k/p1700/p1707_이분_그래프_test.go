package p1700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1700"
	"1d1go/utils"
)

func TestSolve1707(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1707")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
3 2
1 3
2 3
4 4
1 2
2 3
3 4
4 2`,
			`YES
NO
`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1700.Solve1707(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
