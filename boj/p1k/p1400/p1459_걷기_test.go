package p1400

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1459(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1459")

	tests := []struct {
		give string
		want string
	}{
		{
			`4 2 3 10`,
			"18",
		},
		{
			`4 2 3 5`,
			"16",
		},
		{
			`2 0 12 10`,
			"20",
		},
		{
			`25 18 7 11`,
			"247",
		},
		{
			`24 16 12 10`,
			"240",
		},
		{
			`10000000 50000000 800 901`,
			"41010000000",
		},
		{
			`135 122 43 29`,
			"3929",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1459(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
