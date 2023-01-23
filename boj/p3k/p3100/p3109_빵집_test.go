package p3100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve3109(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3109")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 5
.xx..
..x..
.....
...x.
...x.`,
			"2",
		},
		{
			`6 10
..x.......
.....x....
.x....x...
...x...xx.
..........
....x.....`,
			"5",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve3109(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
