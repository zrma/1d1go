package p2100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2156(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2156")

	tests := []struct {
		give string
		want string
	}{
		{
			`6
6
10
13
9
8
1`,
			"33",
		},
		{
			`7
6
10
13
9
8
1
2`,
			"35",
		},
		{
			`8
6
10
13
9
8
1
2
101`,
			"136",
		},
		{
			`1
123`,
			"123",
		},
		{
			`2
12
34`,
			"46",
		},
		{
			`3
1
2
3`,
			"5",
		},
		{
			`4
1
2
3
4`,
			"8",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2156(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
