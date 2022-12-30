package p3000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p3k/p3000"
)

func TestSolve3052(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/3052")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`1
2
3
4
5
6
7
8
9
10`,
			"10",
		},
		{
			`42
84
252
420
840
126
42
84
420
126`,
			"1",
		},
		{
			`39
40
41
42
43
44
82
83
84
85`,
			"6",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p3000.Solve3052(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
