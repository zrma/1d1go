package p1500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1500"
	"1d1go/utils"
)

func TestSolve1504(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1504")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4 6
1 2 3
2 3 3
3 4 1
1 3 5
2 4 5
1 4 4
2 3`,
			"7",
		},
		{
			`4 2
1 2 1
1 3 1
3 4
`,
			"-1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1500.Solve1504(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
