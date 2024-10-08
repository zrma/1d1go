package p1400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1450(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1450")

	tests := []struct {
		give string
		want string
	}{
		{
			`2 1
1 1`,
			"3",
		},
		{
			`1 1
1`,
			"2",
		},
		{
			`1 2
1`,
			"2",
		},
		{
			`2 1
2 2`,
			"1",
		},
		{
			`2 2
1 1`,
			"4",
		},
		{
			`30 30
1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1`,
			"1073741824",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1450(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
