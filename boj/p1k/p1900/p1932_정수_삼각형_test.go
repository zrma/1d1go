package p1900

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1932(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1932")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
7
3 8
8 1 0
2 7 4 4
4 5 2 6 5`,
			"30",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1932(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
