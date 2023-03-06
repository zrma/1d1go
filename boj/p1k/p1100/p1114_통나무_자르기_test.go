package p1100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1114(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1114")

	tests := []struct {
		give string
		want string
	}{
		{
			`9 2 1
4 5`,
			"5 4",
		},
		{
			`5 1 2
3`,
			"3 3",
		},
		{
			`5 5 3
4 2 5 3 1`,
			"2 1",
		},
		{
			`5 10 10
4 3 2 1 4 3 2 1 4 3`,
			"1 1",
		},
		{
			`20 6 5
3 6 9 12 15 18`,
			"5 3",
		},
		{
			`10 4 2
1 4 5 10`,
			"5 1",
		},
		{
			`9 9 2
1 2 3 4 5 6 7 8 9`,
			"3 3",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1114(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
