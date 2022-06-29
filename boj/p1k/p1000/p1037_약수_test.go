package p1000_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
	"1d1go/utils"
)

func TestSolve1037(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1037")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`2
4 2`,
			"8",
		},
		{
			`4
2 3 4 6`,
			"12",
		},
		{
			`7
2 3 4 6 9 12 18`,
			"36",
		},
		{
			`1
2`,
			"4",
		},
		{
			`6
3 4 2 12 6 8`,
			"24",
		},
		{
			`14
14 26456 2 28 13228 3307 7 23149 8 6614 46298 56 4 92596`,
			"185192",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p1000.Solve1037(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
