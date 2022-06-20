package p2700_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
	"1d1go/utils"
)

func TestSolve2742(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2742")

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
			`2
1
`,
		},
		{
			"3",
			`3
2
1
`,
		},
		{
			"4",
			`4
3
2
1
`,
		},
		{
			"5",
			`5
4
3
2
1
`,
		},
		{
			"6",
			`6
5
4
3
2
1
`,
		},
	} {
		t.Run(tt.give, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p2700.Solve2742(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
