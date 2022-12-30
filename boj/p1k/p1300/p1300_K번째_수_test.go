package p1300_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1300"
)

func TestSolve1300(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1300")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`3
7`,
			"6",
		},
		{
			`100000
1000000000`,
			"204535821",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p1300.Solve1300(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
