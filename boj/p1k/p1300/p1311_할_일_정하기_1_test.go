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

func TestSolve1311(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1311")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
2 3 3
3 2 3
3 3 2`,
			`6`,
		},
		{
			`4
1 2 3 4
2 3 4 5
6 7 8 9
10 11 12 13`,
			`25`,
		},
		{
			`10
1 2 3 4 5 6 7 8 9 10
11 12 13 14 15 16 17 18 19 20
21 22 23 24 25 26 27 28 29 30
31 32 33 34 35 36 37 38 39 40
41 42 43 44 45 46 47 48 49 50
51 52 53 54 55 56 57 58 59 60
61 62 63 64 65 66 67 68 69 70
71 72 73 74 75 76 77 78 79 80
81 82 83 84 85 86 87 88 89 90
91 92 93 94 95 96 97 98 99 100`,
			`505`,
		},
	} {
		t.Run(fmt.Sprintf("%d/DFS", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1300.Solve1311(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})

		t.Run(fmt.Sprintf("%d/HungarianAlgorithm", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1300.Solve1311WithHungarianAlgorithm(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
