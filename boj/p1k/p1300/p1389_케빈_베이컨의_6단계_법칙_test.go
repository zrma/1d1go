package p1300_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1300"
	"1d1go/utils"
)

func TestSolve1389(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1389")

	for i, tt := range []struct {
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
	} {
		t.Run(fmt.Sprintf("%d/floydWarshall", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1300.Solve1389(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/bfs", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1300.Solve1389WithBFS(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
