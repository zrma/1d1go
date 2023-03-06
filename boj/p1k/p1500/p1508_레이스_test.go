package p1500

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1508(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1508")

	tests := []struct {
		give string
		want string
	}{
		{
			`11 3 4
0 5 10 11`,
			"1110",
		},
		{
			`11 2 4
0 5 10 11`,
			"1001",
		},
		{
			`11 4 4
0 5 10 11`,
			"1111",
		},
		{
			`1000 5 10
6 9 33 59 100 341 431 444 565 857`,
			"1000010111",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1508(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
