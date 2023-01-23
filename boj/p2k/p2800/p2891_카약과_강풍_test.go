package p2800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2891(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2891")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 2 1
2 4
3`,
			"1",
		},
		{
			`5 2 3
2 4
1 3 5`,
			"0",
		},
		{
			`10 1 1
1 3`,
			"1",
		},
		{
			`10 5 2
1 2 3 6 7
7 8`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2891(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}

}
