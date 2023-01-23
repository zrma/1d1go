package p1200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1260(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1260")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4 5 1
1 2
1 3
1 4
2 4
3 4`,
			`1 2 4 3
1 2 3 4`,
		},
		{
			`5 5 3
5 4
5 2
1 2
3 4
3 1`,
			`3 1 2 5 4
3 1 4 2 5`,
		},
		{
			`1000 1 1000
999 1000`,
			`1000 999
1000 999`,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1260(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
