package p2200

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2213(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2213")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`7
10 30 40 10 20 20 70
1 2
2 3
4 3
4 5
6 2
6 7`,
			`140
1 3 5 7 `,
		},
		{
			`6
10 1 1 1 10 10
1 2
2 3
3 4
4 5
3 6`,
			`30
1 5 6 `,
		},
		{
			`6
10 1 1 1 10 10
6 3
5 4
4 3
3 2
2 1`,
			`30
1 5 6 `,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2213(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
