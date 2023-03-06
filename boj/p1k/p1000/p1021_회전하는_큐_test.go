package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1021(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1021")

	tests := []struct {
		give string
		want string
	}{
		{
			`10 3
1 2 3`,
			"0",
		},
		{
			`10 3
2 9 5`,
			"8",
		},
		{
			`32 6
27 16 30 11 6 23`,
			"59",
		},
		{
			`10 10
1 6 3 2 7 9 8 4 10 5`,
			"14",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1021(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
