package p2500_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2530(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2530")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`14 30 0
200`,
			"14 33 20",
		},
		{
			`17 40 45
6015`,
			"19 21 0",
		},
		{
			`23 48 59
2515`,
			"0 30 54",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p2500.Solve2530(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
