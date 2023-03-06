package p1000

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1037(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1037")

	tests := []struct {
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
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)

			Solve1037(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
