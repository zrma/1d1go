package p20000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve20040(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/20040")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6 5
0 1
1 2
2 3
5 4
0 4`,
			"0",
		},
		{
			`6 5
0 1
1 2
1 3
0 3
4 5`,
			"4",
		},
		{
			`6 5
0 1
1 2
1 3
0 3
2 3`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve20040(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
