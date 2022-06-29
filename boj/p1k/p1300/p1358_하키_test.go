package p1300_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1300"
	"1d1go/utils"
)

func TestSolve1358(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1358")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`20 10 5 0 3
15 5
1 5
1 1`,
			"2",
		},
		{
			`20 10 0 0 14
-5 5
-4 2
-4 8
-3 1
-3 9
0 0
0 10
20 0
20 10
23 1
23 9
24 2
24 8
25 5`,
			"14",
		},
		{
			`52 84 -44 66 10
26 118
-33 106
-49 128
40 114
-10 101
47 85
25 142
-16 140
-82 126
7 145`,
			"8",
		},
		{
			`24 100 -62 71 8
-63 109
-26 164
-9 91
-113 80
-124 75
-95 140
-89 116
-55 113`,
			"6",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1300.Solve1358(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
