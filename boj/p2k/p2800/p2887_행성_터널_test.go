package p2800

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2887(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2887")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1
1 2 3`,
			"0",
		},
		{
			`3
0 0 0
1 0 0
1 1 0`,
			"0",
		},
		{
			`5
11 -15 -15
14 -5 -15
-1 -1 -5
10 -4 -1
19 -4 19`,
			"4",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve2887(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
