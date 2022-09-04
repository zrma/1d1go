package p16900_test

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p16k/p16900"
	"1d1go/utils"
)

//go:embed test_data/p16928_want.txt
var p16928want string

func TestSolve16928(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/16928")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3 7
32 62
42 68
12 98
95 13
97 25
93 37
79 27
75 19
49 47
67 17`,
			"3",
		},
		{
			`4 9
8 52
6 80
26 42
2 72
51 19
39 11
37 29
81 3
59 5
79 23
53 7
43 33
77 21`,
			"5",
		},
		{
			`2 1
2 92
91 99
93 90`,
			"3",
		},
		{
			`2 1
2 82
81 99
83 80`,
			"4",
		},
		{
			`1 0
92 99`,
			"17",
		},
		{
			`1 0
91 99`,
			"16",
		},
		{
			`1 0
91 98`,
			"16",
		},
		{
			p16928want,
			"15",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p16900.Solve16928(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
