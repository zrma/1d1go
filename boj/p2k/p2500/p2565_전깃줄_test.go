package p2500_test

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2565(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2565")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`8
1 8
3 9
2 2
4 1
6 4
10 10
9 7
7 6`,
			"3",
		},
		{
			`1
1 1`,
			"0",
		},
		{
			`2
1 1
2 2`,
			"0",
		},
		{
			`2
1 2
2 1`,
			"1",
		},
		{
			`3
1 1
2 2
3 3`,
			"0",
		},
		{
			`3
1 1
2 3
3 2`,
			"1",
		},
		{
			`3
1 3
2 2
3 1`,
			"2",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)

			p2500.Solve2565(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
