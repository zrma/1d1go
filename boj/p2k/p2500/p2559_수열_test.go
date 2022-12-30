package p2500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2559(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2559")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`10 2
3 -2 -4 -9 0 3 7 13 8 -3`,
			"21",
		},
		{
			`10 5
3 -2 -4 -9 0 3 7 13 8 -3`,
			"31",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2500.Solve2559(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
