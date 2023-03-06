package p2100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2141(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2141")

	tests := []struct {
		give string
		want string
	}{
		{
			`3
1 3
2 5
3 3`,
			"2",
		},
		{
			`2
1 3
2 3`,
			"1",
		},
		{
			`2
1 1
2 2`,
			"2",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2141(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
