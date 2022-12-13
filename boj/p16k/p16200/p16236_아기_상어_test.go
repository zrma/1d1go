package p16200_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p16k/p16200"
)

func TestSolve16236(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/16236")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
0 0 0
0 0 0
0 9 0`,
			"0",
		},
		{
			`3
0 0 1
0 0 0
0 9 0`,
			"3",
		},
		{
			`4
4 3 2 1
0 0 0 0
0 0 9 0
1 2 3 4`,
			"14",
		},
		{
			`6
5 4 3 2 3 4
4 3 2 3 4 5
3 2 9 5 6 6
2 1 2 3 4 5
3 2 1 6 5 4
6 6 6 6 6 6`,
			"60",
		},
		{
			`6
6 0 6 0 6 1
0 0 0 0 0 2
2 3 4 5 6 6
0 0 0 0 0 2
0 2 0 0 0 0
3 9 3 0 0 1`,
			"48",
		},
		{
			`6
1 1 1 1 1 1
2 2 6 2 2 3
2 2 5 2 2 3
2 2 2 4 6 3
0 0 0 0 0 6
0 0 0 0 0 9`,
			"39",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p16200.Solve16236(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
