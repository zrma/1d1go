package p2200_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2200"
	"1d1go/utils"
)

func TestSolve2206(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2206")

	for i, tt := range []struct {
		give string
		want string
	}{
		{
			`6 4
0100
1110
1000
0000
0111
0000`,
			"15",
		},
		{
			`4 4
0111
1111
1111
1110`,
			"-1",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.give))
			writer := utils.NewStringWriter()

			p2200.Solve2206(reader, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
