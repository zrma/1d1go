package p2700_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
	"1d1go/utils"
)

func TestSolve2741(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2741")

	for _, tt := range []struct {
		give string
		want string
	}{
		{
			"1",
			`1
`,
		},
		{
			"2",
			`1
2
`,
		},
		{
			"3",
			`1
2
3
`,
		},
		{
			"4",
			`1
2
3
4
`,
		},
		{
			"5",
			`1
2
3
4
5
`,
		},
		{
			"6",
			`1
2
3
4
5
6
`,
		},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2700.Solve2741(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
