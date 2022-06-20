package p1100_test

import (
	"testing"

	"1d1go/boj/p1k/p1100"
	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolve1149(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1149")

	for _, tt := range []struct {
		give string
		want string
	}{
		{
			`3
26 40 83
49 60 57
13 89 99`,
			"96",
		},
		{
			`3
1 100 100
100 1 100
100 100 1`,
			"3",
		},
		{
			`3
1 100 100
100 100 100
1 100 100`,
			"102",
		},
		{
			`6
30 19 5
64 77 64
15 19 97
4 71 57
90 86 84
93 32 91`,
			"208",
		},
		{
			`8
71 39 44
32 83 55
51 37 63
89 29 100
83 58 11
65 13 15
47 25 29
60 66 19`,
			"253",
		},
	} {
		t.Run(tt.give, func(t *testing.T) {
			scanner := utils.NewStringScanner(tt.give)
			writer := utils.NewStringWriter()

			p1100.Solve1149(scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
