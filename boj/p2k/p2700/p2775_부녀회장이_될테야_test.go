package p2700_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2700"
	"1d1go/utils"
)

func TestSolve2775(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2775")

	for _, tt := range []struct {
		give string
		want string
	}{
		{
			`2
1
3
2
3`,
			`6
10
`,
		},
	} {
		t.Run(tt.give, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2700.Solve2775(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
