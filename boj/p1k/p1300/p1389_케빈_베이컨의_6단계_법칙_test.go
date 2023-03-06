package p1300

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1389(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1389")

	tests := []struct {
		give string
		want string
	}{
		{
			`5 5
1 3
1 4
4 5
4 3
3 2`,
			"3",
		},
		{
			`5 0`,
			"1",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d/FloydWarshall", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1389(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/bfs", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1389WithBFS(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
