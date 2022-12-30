package p24000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p24k/p24000"
)

func TestSolve24060(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/24060")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`5 1
4 5 1 3 2`,
			"4",
		},
		{
			`5 2
4 5 1 3 2`,
			"5",
		},
		{
			`5 3
4 5 1 3 2`,
			"1",
		},
		{
			`5 4
4 5 1 3 2`,
			"4",
		},
		{
			`5 5
4 5 1 3 2`,
			"5",
		},
		{
			`5 6
4 5 1 3 2`,
			"2",
		},
		{
			`5 7
4 5 1 3 2`,
			"3",
		},
		{
			`5 8
4 5 1 3 2`,
			"1",
		},
		{
			`5 9
4 5 1 3 2`,
			"2",
		},
		{
			`5 10
4 5 1 3 2`,
			"3",
		},
		{
			`5 11
4 5 1 3 2`,
			"4",
		},
		{
			`5 12
4 5 1 3 2`,
			"5",
		},
		{
			`5 13
4 5 1 3 2`,
			"-1",
		},
		{
			`5 13
4 5 1 3 2`,
			"-1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			p24000.Solve24060(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
