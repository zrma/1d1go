package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1032(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1032")

	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		give string
		want string
	}{
		{
			`3
config.sys
config.inf
configures`,
			`config????`,
		},
		{
			`2
contest.txt
context.txt`,
			`conte?t.txt`,
		},
		{
			`3
c.user.mike.programs
c.user.nike.programs
c.user.rice.programs`,
			`c.user.?i?e.programs`,
		},
		{
			`4
a
a
b
b`,
			`?`,
		},
		{
			`1
onlyonefile`,
			`onlyonefile`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1032(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
