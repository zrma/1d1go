package p1700_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1700"
)

func TestSolve1715(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1715")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
10
20
40`,
			"100",
		},
		{
			`10
123
456
234
998
12
7
876
887
455
234`,
			"12108",
		},
		{
			`1
100`,
			"0",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1700.Solve1715(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
