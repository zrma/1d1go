package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1700(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1700")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2 7
2 3 2 3 1 2 7`,
			"2",
		},
		{
			`3 13

2 3 4 2 3 4 1 5 5 5 2 3 2`,
			"2",
		},
		{
			`3 5
1 1 1 1 2`,
			"0",
		},
		{
			`3 5
1 1 1 1 2`,
			"0",
		},
		{
			`1 5
1 2 3 4 1`,
			"4",
		},
		{
			`3 14
1 4 3 2 5 4 3 2 5 3 4 2 3 4`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1700(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
