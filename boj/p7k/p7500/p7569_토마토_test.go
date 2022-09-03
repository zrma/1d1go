package p7500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p7k/p7500"
	"1d1go/utils"
)

func TestSolve7569(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/7569")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 3 1
0 -1 0 0 0
-1 -1 0 1 1
0 0 0 1 1`,
			"-1",
		},
		{
			`5 3 2
0 0 0 0 0
0 0 0 0 0
0 0 0 0 0
0 0 0 0 0
0 0 1 0 0
0 0 0 0 0`,
			"4",
		},
		{
			`4 3 2
1 1 1 1
1 1 1 1
1 1 1 1
1 1 1 1
-1 -1 -1 -1
1 1 1 -1`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p7500.Solve7569(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
