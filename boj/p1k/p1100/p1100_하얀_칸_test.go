package p1100

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1100(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1100")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`.F.F...F
F...F.F.
...F.F.F
F.F...F.
.F...F..
F...F.F.
.F.F.F.F
..FF..F.`,
			"1",
		},
		{
			`........
........
........
........
........
........
........
........`,
			"0",
		},
		{
			`FFFFFFFF
FFFFFFFF
FFFFFFFF
FFFFFFFF
FFFFFFFF
FFFFFFFF
FFFFFFFF
FFFFFFFF`,
			"32",
		},
		{
			`........
..F.....
.....F..
.....F..
........
........
.......F
.F......`,
			"2",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1100(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
