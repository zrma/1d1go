package p2100_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2100"
	"1d1go/utils"
)

func TestSolve2178(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2178")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`4 6
101111
101010
101011
111011`,
			"15",
		},
		{
			`4 6
110110
110110
111111
111101`,
			"9",
		},
		{
			`2 25
1011101110111011101110111
1110111011101110111011101`,
			"38",
		},
		{
			`7 7
1011111
1110001
1000001
1000001
1000001
1000001
1111111`,
			"13",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2100.Solve2178(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
