package p17400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p17k/p17400"
)

func TestSolve17404(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/17404")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
26 40 83
49 60 57
13 89 99`,
			"110",
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
			"201",
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p17400.Solve17404(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
