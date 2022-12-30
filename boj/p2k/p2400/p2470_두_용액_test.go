package p2400_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2400"
)

func TestSolve2470(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2470")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5
-2 4 -99 -1 98`,
			"-99 98",
		},
		{
			`5
-2 3 -99 1 97`,
			"-2 3",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2400.Solve2470(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
