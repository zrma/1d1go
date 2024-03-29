package p1700

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1781(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1781")

	tests := []struct {
		give string
		want string
	}{
		{
			`7
1 6
1 7
3 2
3 1
2 4
2 5
6 1`,
			"15",
		},
		{
			`3
3 3
3 2
3 1`,
			"6",
		},
		{
			`4
3 3
3 2
3 4
2 4`,
			"11",
		},
		{
			`5
3 1
3 2
2 3
4 4
4 5`,
			"14",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1781(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
