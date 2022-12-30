package p15800_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15800"
)

func TestSolve15828(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15828")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
1
2
0
3
4
0
5
6
0
0
-1`,
			"5 6 ",
		},
		{
			`1
1
2
3
4
5
6
7
-1`,
			"1 ",
		},
		{
			`1
1
2
0
3
4
0
5
6
0
7
0
-1`,
			"empty",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p15800.Solve15828(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
