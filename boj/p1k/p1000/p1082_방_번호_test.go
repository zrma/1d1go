package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1082(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1082")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
6 7 8
21`,
			"210",
		},
		{
			`3
5 23 24
30`,
			"20",
		},
		{
			`4
1 5 3 2
1`,
			"0",
		},
		{
			`10
1 1 1 1 1 1 1 1 1 1
50`,
			"99999999999999999999999999999999999999999999999999",
		},
		{
			`1
10
10`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1082(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
