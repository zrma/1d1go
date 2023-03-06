package p1800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1806(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1806")

	tests := []struct {
		give string
		want string
	}{
		{
			`10 15
5 1 3 5 10 7 4 9 2 8`,
			"2",
		},
		{
			`10 11
10 1 1 1 1 1 1 1 1 1`,
			"2",
		},
		{
			`10 10
10 1 1 1 1 1 1 1 1 1`,
			"1",
		},
		{
			`10 100
1 1 1 1 1 1 1 1 1 1`,
			"0",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1806(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
